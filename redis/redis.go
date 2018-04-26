package redis

import (
	"fmt"

	"os"

	"github.com/garyburd/redigo/redis"
	"github.com/json-iterator/go"
)

func TestRedis() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	//easyjson:json
	var p2 struct {
		AdScheme       string `redis:"adScheme" json:"adScheme"`
		ActiveInterval string `redis:"activeInterval" json:"activeInterval"`
		BeginTime      string `redis:"beginTime" json:"beginTime"`
		Time1          string `redis:"time1" json:"time1"`
		Time2          string `redis:"time2" json:"time2"`
	}

	valueGet, err := redis.Values(c.Do("HGETALL", "tvchannel:1782"))
	if err != nil {
		fmt.Println(err)
	}
	if err := redis.ScanStruct(valueGet, &p2); err != nil {
		panic(err)
	}
	fmt.Printf("redis==========%+v\n", p2)
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err1 := jsonIterator.Marshal(p2) //encoding/json
	if err1 != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", b)
	os.Stdout.Write(b)
	fmt.Println("=======redis=====")
	//fmt.Println(imapGet["phonenumber"])

}
