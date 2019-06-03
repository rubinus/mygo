package do

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"time"

	jsoniter "github.com/json-iterator/go"
)

var tvmid = "sjh594b31593e800741dde24eae"
var wx_token = "46497107fa23"

var reqQ = []req{
	{
		url: "http://pmall.yaotv.tvm.cn/open/cashaccount?tvmid=" + tvmid,
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
	in     chan []byte
	parser Parser
	done   func()
}

func dealWork(url string, work workers) {
	fmt.Println(url)

	go func() {
		timeout := 3 * time.Second
		ctx, _ := context.WithTimeout(context.Background(), timeout)
		_, err := getHttp(ctx, url, work)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
	}()
}

//channel做为返回值
func createWorker(i int, group *sync.WaitGroup) workers {
	work := workers{
		in:     make(chan []byte, 1),
		parser: reqQ[i].parser,
		done: func() {
			group.Done()
		},
	}
	//go dealWork(reqQ[i].url, work)
	timeout := 3 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	go func() {
		_, err := getHttp(ctx, reqQ[i].url, work)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
	}()

	return work
}

func DoWork(j int) AllResult {
	fmt.Println(j, "并发开始...")
	works := [3]workers{}
	var group sync.WaitGroup

	for i, _ := range works {
		works[i] = createWorker(i, &group)
	}
	group.Add(3)

	all := AllResult{}
	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary

	for _, w := range works {
		select {
		case out := <-w.in:
			parser := w.parser
			result := parser.parseJson()
			//fmt.Println(out, "------")
			if err := jsonIterator.Unmarshal(out, &result); err != nil {
				fmt.Printf("%v", err)
			}
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

	}

	fmt.Printf("%+v == %v \n", all, j)
	return all
}

func getHttp(ctx context.Context, url string, work workers) ([]byte, error) {

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Deadline())
		}
	}()
	defer close(work.in)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("ReadAll is error")
	}
	//fmt.Println(string(out))

	//fmt.Println("====", result)

	work.in <- out
	work.done()

	return out, nil
}

type req struct {
	url    string
	parser Parser
}

type AllResult struct {
	Cash       int64 `json:"cash"`
	HoldNumber int64 `json:"holdNumber"`
	Seed       int64 `json:"seed"`
}

type Cash struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   int64  `json:"data"`
	ErrMsg string `json:"errMsg"`
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
