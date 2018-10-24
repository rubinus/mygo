package main

import (
	"fmt"

	"time"

	"github.com/mediocregopher/radix.v3"
)

func main() {
	customConnFunc := func(network, addr string) (radix.Conn, error) {
		return radix.Dial(network, addr,
			radix.DialTimeout(10*time.Second),
			radix.DialAuthPass(""),
		)
	}

	poolFunc := func(network, addr string) (radix.Client, error) {
		return radix.NewPool(network, addr, 10, radix.PoolConnFunc(customConnFunc))
	}

	//var RedisClusterIP = []string{"106.15.228.49:7001", "106.15.228.49:7002", "106.15.228.49:7003"}
	var RedisClusterIP = []string{
		"127.0.0.1:7000",
		"127.0.0.1:7001",
		"127.0.0.1:7002",
		"127.0.0.1:7006",
	}

	client, err := radix.NewCluster(RedisClusterIP, radix.ClusterPoolFunc(poolFunc))

	//client, err := radix.NewPool("tcp", "106.15.228.49:7003", 10)
	//if err != nil {
	//	// handle error
	//}

	if err != nil {
		// handle error
		fmt.Println("===", err)
		return
	}
	var fooVal map[string]string
	err = client.Do(radix.Cmd(&fooVal, "HGETALL", "hh"))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fooVal)

	time.Sleep(10 * time.Second)

}
