package main

import (
	"fmt"

	"time"

	"github.com/mediocregopher/radix.v3"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Unionid    string        `json:"unionid,omitempty" bson:"unionid,omitempty"`
	MinAppid   string        `json:"minappid" bson:"minappid"`
	MinOpenid  string        `json:"minopenid" bson:"minopenid"`
	Nickname   string        `json:"nickname,omitempty" bson:"nickname,omitempty"`
	Headimgurl string        `json:"headimgurl,omitempty" bson:"headimgurl,omitempty"`
	City       string        `json:"city,omitempty" bson:"city,omitempty"`
	Sex        string        `json:"sex,omitempty" bson:"sex,omitempty"`
	Province   string        `json:"province,omitempty" bson:"province,omitempty"`
	Country    string        `json:"country,omitempty" bson:"country,omitempty"`
	CreateTime string        `json:"createTime" bson:"createTime"`
	UpdateTime int64         `json:"updateTime,omitempty" bson:"updateTime,omitempty"`
}

func main() {
	customConnFunc := func(network, addr string) (radix.Conn, error) {
		return radix.Dial(network, addr,
			radix.DialTimeout(10*time.Second),
			radix.DialAuthPass(""),
			radix.DialSelectDB(9),
		)
	}

	//poolFunc := func(network, addr string) (radix.Client, error) {
	//	return radix.NewPool(network, addr, 10, radix.PoolConnFunc(customConnFunc))
	//}

	//var RedisClusterIP = []string{"106.15.228.49:7001", "106.15.228.49:7002", "106.15.228.49:7003"}
	//var RedisClusterIP = []string{
	//	"127.0.0.1:6379",
	//}

	//client, err := radix.NewCluster(RedisClusterIP, radix.ClusterPoolFunc(poolFunc))

	client, err := radix.NewPool("tcp", "localhost:6379", 10, radix.PoolConnFunc(customConnFunc))
	if err != nil {
		// handle error
	}

	if err != nil {
		// handle error
		fmt.Println("===", err)
		return
	}
	var fooVal map[string]string

	err = client.Do(radix.FlatCmd(&fooVal, "HGETALL", "USER:5bee7bdc28ddf5e68bcf3c45"))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v", fooVal)
	var result User
	err = mapstructure.Decode(fooVal, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Printf("%#v", result)

	time.Sleep(10 * time.Second)

}
