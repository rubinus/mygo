package main

import (
	"fmt"
	"mygo/redispool"

	"github.com/garyburd/redigo/redis"
	"github.com/json-iterator/go"
)

func main() {
	c, err := redispool.GetConn()
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()

	var p2 struct {
		Nickname   string `redis:"nickname" json:"nickname"`
		Province   string `redis:"province" json:"province"`
		Headimgurl string `redis:"headimgurl" json:"headimgurl"`
		City       string `redis:"city" json:"city"`
		Country    string `redis:"country" json:"country"`
	}

	valueGet, err := redis.Values(c.Do("HGETALL", "USER:5b92269fa33c300001007709"))
	if err != nil {
		fmt.Println(err)
	}
	if err := redis.ScanStruct(valueGet, &p2); err != nil {
		panic(err)
	}
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err1 := jsonIterator.Marshal(p2) //encoding/json
	if err1 != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf(string(b))
	fmt.Println("===")
	fmt.Println(redispool.DisplayStats())
}
