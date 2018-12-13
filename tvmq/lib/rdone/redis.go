package rdone

import (
	"fmt"

	"github.com/mediocregopher/radix.v3"
)

func DelKey(c *radix.Pool, key string) (int64, error) {
	var reply int64
	err := c.Do(radix.Cmd(&reply, "DEL", key))
	if err != nil {
		fmt.Println(err)
	}
	return reply, err //1,nil
}

func Hmset(c *radix.Pool, key string, values interface{}, ttl string) (string, error) {
	var reply string
	var expire int
	err := c.Do(radix.FlatCmd(&reply, "HMSET", key, values))
	if ttl != "" {
		err = c.Do(radix.Cmd(&expire, "EXPIRE", key, ttl))
	}
	if err != nil {
		fmt.Println(err)
	}
	return reply, err //OK,nil
}

func Hgetall(c *radix.Pool, key string, input map[string]string) (map[string]string, error) {
	err := c.Do(radix.FlatCmd(&input, "HGETALL", key))
	if err != nil {
		fmt.Println(err)
	}
	return input, err
}

func Setex(c *radix.Pool, key string, time int, value string) (string, error) {
	var reply string
	err := c.Do(radix.FlatCmd(&reply, "SETEX", key, time, value))
	if err != nil {
		fmt.Println(err)
	}
	return reply, err //OK,nil
}

func Get(c *radix.Pool, key string) (string, error) {
	var reply string
	err := c.Do(radix.Cmd(&reply, "GET", key))
	if err != nil {
		fmt.Println(err)
	}
	return reply, err
}

func Sadd(c *radix.Pool, key, member string) (int64, error) {
	var reply int64
	err := c.Do(radix.Cmd(&reply, "SADD", key, member))
	if err != nil {
		fmt.Println("sadd is err", err)
	}
	return reply, err
}

func Scard(c *radix.Pool, key string) (int64, error) {
	var reply int64
	err := c.Do(radix.Cmd(&reply, "SCARD", key))
	if err != nil {
		fmt.Println("Scard is err", err)
	}
	return reply, err
}

func Srem(c *radix.Pool, key, member string) (int64, error) {
	var reply int64
	err := c.Do(radix.Cmd(&reply, "SREM", key, member))
	if err != nil {
		fmt.Println("Srem is err", err)
	}
	return reply, err
}

func Smembers(c *radix.Pool, key string, reply []string) ([]string, error) {
	err := c.Do(radix.Cmd(&reply, "SMEMBERS", key))
	if err != nil {
		fmt.Println(err)
	}
	return reply, err
}

func HmsetAndSet(c *radix.Pool, key string, values interface{},
	cs *radix.Pool, setKey, member string, ttl string) (string, int64) {
	var reply string
	var r2 int64
	var expire string
	err := c.Do(radix.FlatCmd(&reply, "HMSET", key, values))
	if ttl != "" {
		err = c.Do(radix.Cmd(&expire, "EXPIRE", key, ttl))
	}
	err = cs.Do(radix.Cmd(&r2, "SADD", setKey, member))
	if err != nil {
		fmt.Println(err)
	}
	return reply, r2
}

func HmsetAndHmset(c *radix.Pool, key string, values interface{},
	cs *radix.Pool, key2 string, values2 interface{}, ttl string) (string, string) {
	var reply string
	var reply2 string
	var expire string
	err := c.Do(radix.FlatCmd(&reply, "HMSET", key, values))
	if ttl != "" {
		err = c.Do(radix.Cmd(&expire, "EXPIRE", key, ttl))
	}
	err = cs.Do(radix.FlatCmd(&reply2, "HMSET", key2, values2))
	if ttl != "" {
		err = c.Do(radix.Cmd(&expire, "EXPIRE", key2, ttl))
	}
	if err != nil {
		fmt.Println(err)
	}
	return reply, reply2
}

func DelKeyAndSetMember(c *radix.Pool, key string, cs *radix.Pool, setKey, member string) (int64, int64) {
	var reply int64
	var r2 int64
	err := c.Do(radix.Cmd(&reply, "DEL", key))
	err = cs.Do(radix.Cmd(&r2, "SREM", setKey, member))
	if err != nil {
		fmt.Println(err)
	}
	return reply, r2
}

func Zadd(c *radix.Pool, key, score, member string) (int64, error) {
	var reply int64
	err := c.Do(radix.Cmd(&reply, "ZADD", key, score, member))
	if err != nil {
		fmt.Println("zadd is err", err)
	}
	return reply, err
}

func Zrange(c *radix.Pool, key, start, stop string, st int) ([]string, error) {
	var reply []string
	var cmd string
	if st == 1 {
		cmd = "ZRANGE"
	} else {
		cmd = "ZREVRANGE"
	}
	err := c.Do(radix.Cmd(&reply, cmd, key, start, stop))
	if err != nil {
		fmt.Println("zrange is err :", err)
	}
	return reply, err
}

func Zcard(c *radix.Pool, key string) (int64, error) {
	var reply int64
	err := c.Do(radix.Cmd(&reply, "ZCARD", key))
	if err != nil {
		fmt.Println("Zcard is err", err)
	}
	return reply, err
}

func Zremrangebyrank(c *radix.Pool, key, start, stop string) (int64, error) {
	var reply int64
	err := c.Do(radix.Cmd(&reply, "ZREMRANGEBYRANK", key, start, stop))
	if err != nil {
		fmt.Println("ZREMRANGEBYRANK is err", err)
	}
	return reply, err
}
