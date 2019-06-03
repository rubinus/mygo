package main

import (
	"fmt"
	"time"

	//"code.tvmining.com/tvplay/tvmq/config"
	"github.com/mediocregopher/radix"
)

var client *radix.Pool

func init() {

	customConnFunc := func(network, addr string) (radix.Conn, error) {
		return radix.Dial(network, addr,
			radix.DialTimeout(10*time.Second), radix.DialSelectDB(9), radix.DialAuthPass(""),
		)
	}
	c, err := radix.NewPool("tcp", "127.0.0.1:6379", 100, radix.PoolConnFunc(customConnFunc))
	if err != nil {
		fmt.Println("redis ", err)
	}
	client = c
}

func GetConn() *radix.Pool {
	return client
}

func main() {

}

//var pool *redis.Pool
//
//var redisHostStr string
//
//func init() {
//	redisHost := os.Getenv("REDIS_HOST")
//	if redisHost != "" {
//		redisHostStr = fmt.Sprintf("%s:%d", redisHost, config.DefaultRedisPort)
//	}else{
//		redisHostStr = fmt.Sprintf("%s:%d", config.DefaultRedisHost, config.DefaultRedisPort)
//	}
//	p, err := newPool(redisHostStr, "", 9)
//	if err != nil {
//		fmt.Printf("db %d is create pool failed ", 9)
//	}
//
//	//go func() {
//	//	for range time.Tick(1 * time.Second){
//	//		fmt.Printf("%+v===\n",p.Stats())
//	//	}
//	//}()
//
//	pool = p
//}
//
//func newPool(addr string, password string, db int) (*redis.Pool, error) {
//	return &redis.Pool{
//		MaxIdle:     5000,             //表示连接池空闲连接列表的长度限制
//		MaxActive:   10000,            //表示连接池中最大连接数限制
//		IdleTimeout: 10 * time.Second, //空闲连接的超时设置，一旦超时，将会从空闲列表中摘除
//		Dial: func() (redis.Conn, error) {
//			c, err := redis.Dial("tcp", addr)
//			if err != nil {
//				return nil, err
//			}
//			if password != "" {
//				if _, err := c.Do("AUTH", password); err != nil {
//					c.Close()
//					return nil, err
//				}
//			}
//			if db != 0 {
//				if _, err := c.Do("SELECT", db); err != nil {
//					c.Close()
//					return nil, err
//				}
//			}
//			return c, err
//		},
//		TestOnBorrow: func(c redis.Conn, t time.Time) error {
//			if time.Since(t) < time.Minute {
//				return nil
//			}
//			_, err := c.Do("PING")
//			return err
//		},
//	}, nil
//}
//
//func GetConn() (redis.Conn, error) {
//	conn := pool.Get()
//	return conn, nil
//}
