package rdcluster

import (
	"fmt"
	"testing"

	"code.tvmining.com/tvplay/tvmq/models/socuser"
	"code.tvmining.com/tvplay/tvmq/redisclusterpool"
	"github.com/mitchellh/mapstructure"
)

var client = redisclusterpool.GetConn()
var client2 = redisclusterpool.GetConn()

func TestHmset(t *testing.T) {
	su := socuser.SoconnUser{
		UserId: "123",
		Host:   "123",
		Cid:    "123",
	}
	m := su
	r, err := Hmset(client, "abc", m)

	fmt.Println(r, "====", err)
}

func TestDelKey(t *testing.T) {
	r, err := DelKey(client, "ab")

	fmt.Println(r, "====", err)
}

func TestHgetall(t *testing.T) {
	var in map[string]string

	reply, err := Hgetall(client, "USER:1000", in)
	var su = socuser.Userinfo{}

	mapstructure.Decode(reply, &su)

	fmt.Println(su, "====", err)
}
func TestSetex(t *testing.T) {
	r, err := Setex(client, "se", 300, "123")
	fmt.Println(r, err)
}
func TestGet(t *testing.T) {
	r, err := Get(client, "se")
	fmt.Println(r, err)
}
func TestSadd(t *testing.T) {
	r, err := Sadd(client, "sadd", "abc")
	fmt.Println(r, err)
}

func TestScard(t *testing.T) {
	r, err := Scard(client, "sadd")
	fmt.Println(r, err)
}

func TestSmembers(t *testing.T) {
	var reply []string
	r, err := Smembers(client, "sadd", reply)
	fmt.Println(r, err)
}

func TestSrem(t *testing.T) {
	r, err := Srem(client, "sadd", "123")
	fmt.Println(r, err)
}

func TestHmsetAndSet(t *testing.T) {
	su := socuser.SoconnUser{
		UserId: "123",
		Host:   "123",
		Cid:    "123",
	}
	m := su.StructToMap()
	r, r2 := HmsetAndSet(client, "USER:123", m, client2, "sadd", "1234")
	fmt.Println(r, r2)
}

func TestDelKeyAndSetMember(t *testing.T) {
	r, r2 := DelKeyAndSetMember(client, "USER:123", client2, "sadd", "1234")
	fmt.Println(r, r2)
}
