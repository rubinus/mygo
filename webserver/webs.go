package main

import (
	"fmt"
	"io"
	_ "net/http/pprof"
	"time"

	jsoniter "github.com/json-iterator/go"

	"mygo/morerequest/do"
	"net/http"
)

const (
	mysql = iota
	mongo
	redis
	pika
	nsq
	kafka
	es
	etcd
)

func main() {
	fmt.Println(20000 * 400 / 30490)

	toBeCharge := "2015-01-01 00:00:00"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	fmt.Println("===", theTime)                                     //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println(sr)                                                 //打印输出时间戳 1420041600

	//时间戳转日期
	dataTimeStr := time.Unix(sr, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	fmt.Println(dataTimeStr)

	//str := `{"adid":"101111071031","redirect_platform":"2","city_id":"bj","house_type":3,"data":{"building_id":"117114","title":"\u5206\u4eab\u6807\u9898","abstract":"\u5206\u4eab\u6458\u8981:*","size_common":"360*270","size_common_url":"[\"http:\\\/\\\/img.zgsta.com\\\/adadmin\\\/2019-03-14\\\/e92fef1aec58b7d1ca374b1001da643b.png\"]","size_special":"360*270","size_special_url":null,"click_able":true,"click":"http\"\/\/ad.zhuge.combi=&u=TKrsd6pprjJQ68SVhhPqkVzmjW9j4389ZbI6ehhqHGo%3D&z=zgf"}}`

	switch 6 {
	case mysql:
		fmt.Println(mysql)
	case mongo:
		fmt.Println(mongo)
	case redis:
		fmt.Println(redis)
	case pika:
		fmt.Println(pika)
	case nsq:
		fmt.Println(nsq)
	case kafka:
		fmt.Println(kafka)
	case es:
		fmt.Println("es")
	case etcd:
		fmt.Println(etcd)
	}

	http.HandleFunc("/", One)
	//http.HandleFunc("/more", morequrest)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}

func One(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello go </h1")
}

func morequrest(w http.ResponseWriter, r *http.Request) {
	result := do.DoWork(1)
	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary

	be, _ := jsonIterator.Marshal(result)
	io.WriteString(w, string(be))
}
