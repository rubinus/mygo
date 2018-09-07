package redispool

import (
	"time"

	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	p, err := newPool("10.20.80.132:6379", 9)
	if err != nil {
		fmt.Printf("db %d is create pool failed ", 9)
	}
	pool = p
}

func newPool(addr string, db int) (*redis.Pool, error) {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if db != 0 {
				if _, err := c.Do("SELECT", db); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}, nil
}

func GetConn() (redis.Conn, error) {
	conn := pool.Get()
	return conn, nil
}

func DisplayStats() redis.PoolStats {
	return pool.Stats()
}
