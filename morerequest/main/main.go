package main

import (
	"errors"
	"fmt"
	"net/http"

	"io/ioutil"

	"sync"

	"github.com/json-iterator/go"
)

//var tvmid = "sjh5a7ab7ac84143d74c92b29be"
//var wx_token = "33580c57d3c86f07"

var tvmid = "sjh594b31593e800741dde24eae"
var wx_token = "46497107fa23"

var reqQ = []req{
	{
		url: "http://pmall.yaotv.tvm.cn/open/account/balance?openid=" + tvmid,
		//url: "http://qa.pmall.yaotv.tvm.cn/open/account/balance?openid=" + tvmid,
		parser: &Cash{},
	},
	{
		url: "http://life-app.yaotv.tvm.cn/fastcall/dktimeext/getholdnumber?tvmid=" + tvmid + "&code=HJSJ",
		//url:    "http://qa-wsq.mtq.tvm.cn/fastcall/dktimeext/getholdnumber?tvmid=" + tvmid + "&code=HJSJ",
		parser: &HoldNumber{},
	},
	{
		url: "http://seed.yaotv.tvm.cn/open/user/seed?openId=" + tvmid + "&yyyappid=" + wx_token,
		//url: "http://qa-seed.yaotv.tvm.cn/open/user/seed?openId=" + tvmid + "&yyyappid=" + wx_token,
		parser: &Goldseed{},
	},
}

type workers struct {
	in   chan interface{}
	done func()
}

//channel做为返回值
func createWorker(i int, group *sync.WaitGroup) workers {
	work := workers{
		in: make(chan interface{}),
		done: func() {
			group.Done()
		},
	}
	go func() {
		_, err := getHttp(reqQ[i].url, reqQ[i].parser, work)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
	}()
	return work
}

func doWork(j int) {
	fmt.Println(j, "并发开始...")
	works := [3]workers{}
	var group sync.WaitGroup

	for i, _ := range works {
		works[i] = createWorker(i, &group)
	}
	group.Add(3)

	all := AllResult{}
	for _, w := range works {
		result := <-w.in
		switch result.(type) {
		case *Cash:
			cash := result.(*Cash).Data
			all.Cash = cash
		case *HoldNumber:
			holdNumber := result.(*HoldNumber).Data.Hold_number
			all.HoldNumber = holdNumber
		case *Goldseed:
			seed := result.(*Goldseed).Data.Seed
			all.Seed = seed
		}
	}
	fmt.Printf("%+v == %v \n", all, j)
}

func main() {

	c := make(chan int)
	for i := 0; i < 100; i++ {
		go doWork(i)
	}
	<-c

}

type req struct {
	url    string
	parser Parser
}

type AllResult struct {
	Cash       int64
	HoldNumber int64
	Seed       int64
}

type Cash struct {
	Status string
	Code   int
	Data   int64
}

type Goldseed struct {
	Status string
	Code   int
	Data   GoldsData
}
type GoldsData struct {
	Seed   int64
	AppMsg string
}

type HoldNumber struct {
	Status int
	Data   HoldNumberData
}
type HoldNumberData struct {
	Hold_number int64
	Hang_number int64
}

func (g *Goldseed) parseJson() interface{} {
	return g
}

func (h *HoldNumber) parseJson() interface{} {
	return h
}

func (c *Cash) parseJson() interface{} {
	return c
}

type Parser interface {
	parseJson() interface{}
}

func getHttp(url string, parser Parser, work workers) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("错误的地址:" + url)
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("ReadAll is error")
	}
	//fmt.Println(string(out))

	result := parser.parseJson()
	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	if err := jsonIterator.Unmarshal(out, &result); err != nil {
		return nil, errors.New(fmt.Sprintf("jsoniterator %s", err))
	}

	//fmt.Println("====", result)

	work.in <- result
	work.done()

	return result, nil
}
