package rdone

import (
	"fmt"
	"strconv"
	"testing"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/models/socuser"
	"code.tvmining.com/tvplay/tvmq/redispool"
	"github.com/mitchellh/mapstructure"
)

func TestHmsetAndSet(t *testing.T) {

	tt := socuser.Userinfo{
		Token:    "1000",
		Nickname: "1000",
	}
	m := map[string]string{}
	mapstructure.Decode(tt, &m)
	HmsetAndSet(redispool.GetConn(), "USER:0000", m,
		redispool.GetConn(), config.ChatRoomName, "0000", strconv.Itoa(100))
}

func TestHmset(t *testing.T) {
	tt := socuser.Userinfo{
		Token:    "1001",
		Nickname: "1001",
	}
	m := tt.StructToMap()

	r, err := Hmset(redispool.GetConn(), "USER:1001", m, strconv.Itoa(100))

	fmt.Println(r, "====", err)
}

func TestHgetall(t *testing.T) {
	var in map[string]string

	reply, _ := Hgetall(redispool.GetConn(), "USER:1001", in)
	var su = socuser.Userinfo{}

	mapstructure.Decode(reply, &su)

	fmt.Println(su, "====")
}
