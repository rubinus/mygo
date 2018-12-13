package handler

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"code.tvmining.com/tvplay/tvmq/lib"
	"code.tvmining.com/tvplay/tvmq/models/socuser"
)

func TestPing(t *testing.T) {
	//host := "172.25.0.12"
	//host := "32a76d396792"
	host := "192.168.2.100"
	var errStr string
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	out := RunCMD(ctx, host)
	select {
	case <-ctx.Done():
		errStr = "超时"
	case r := <-out:
		fmt.Println(r)
	}
	if errStr != "" {
		fmt.Println(errStr, "===")
		return
	}
}

func TestRunCMD(t *testing.T) {
	ch := make(chan string)
	go func() {
		for i := 1000; i < 2000; i++ {
			k := strconv.Itoa(i)
			ch <- k
		}
		close(ch)
	}()

	for k := range ch {
		key := "USER:" + k

		tt := socuser.Userinfo{
			Token:    k,
			Nickname: k,
		}
		m := tt.StructToMap()
		reply, _ := lib.JudgeHmset(key, m, strconv.Itoa(100*36000))
		fmt.Println(reply, key)
	}

	time.Sleep(30 * time.Second)
}
