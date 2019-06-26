package redisclusterpool

import (
	"fmt"
	"time"

	"code.tvmining.com/tvplay/tvmq/config"
	"github.com/mediocregopher/radix"
)

var client *radix.Cluster

func init() {
	if config.UseRedisCluster != 1 {
		return
	}
	customConnFunc := func(network, addr string) (radix.Conn, error) {
		return radix.Dial(network, addr,
			radix.DialTimeout(10*time.Second),
			radix.DialAuthPass(""),
		)
	}
	poolFunc := func(network, addr string) (radix.Client, error) {
		return radix.NewPool(network, addr, 100, radix.PoolConnFunc(customConnFunc))
	}

	c, err := radix.NewCluster(config.RedisClusterIP,
		radix.ClusterPoolFunc(poolFunc), radix.ClusterSyncEvery(5*time.Second))
	if err != nil {
		fmt.Println("redis cluster ", err)
	}
	client = c
}

func GetConn() *radix.Cluster {
	return client
}
