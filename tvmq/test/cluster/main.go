package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gitstliu/go-redis-cluster"
)

func main() {
	cluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes:   []string{"106.15.228.49:7001", "106.15.228.49:7002", "106.15.228.49:7003"},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})
	if err != nil {
		log.Fatalf("redis.New error: %s", err.Error())
	}

	reply, err := redis.String(cluster.Do("GET", "abc"))
	fmt.Println(reply, err)

	//_, err = cluster.Do("MSET", "myfoo1", "mybar1", "myfoo2", "mybar2", "myfoo3", "mybar3")
	//if err != nil {
	//	log.Fatalf("MSET error: %s", err.Error())
	//}
	//
	//values, err := redis.Strings(cluster.Do("MGET", "myfoo1", "myfoo5", "myfoo2", "myfoo3", "myfoo4"))
	//if err != nil {
	//	log.Fatalf("MGET error: %s", err.Error())
	//}
	//
	//for i := range values {
	//	fmt.Printf("reply[%d]: %s\n", i, values[i])
	//}
}
