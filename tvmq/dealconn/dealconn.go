package dealconn

import (
	"context"
	"fmt"
	"strconv"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/lib"
	"code.tvmining.com/tvplay/tvmq/models"
	"code.tvmining.com/tvplay/tvmq/models/socuser"
	"code.tvmining.com/tvplay/tvmq/rpcserv"
	"code.tvmining.com/tvplay/tvmq/utils"
)

type SocketConn struct {
	WSchan chan *socuser.SoconnUser
}

func (sc *SocketConn) ReceConn(souser *socuser.SoconnUser) {
	go func() {
		sc.WSchan <- souser
	}()
}

func (sc *SocketConn) SaveUserToHashAndSet(ctx context.Context, setKey string) chan string {
	out := make(chan string)
	if souser, ok := <-sc.WSchan; ok {
		go func() {
			defer close(out)
			key := fmt.Sprintf("%s:%s", config.ChatPreixForUser, souser.UserId)

			m := souser.StructToMap()

			reply, r2 := lib.JudgeHmsetAndSet(key, m, setKey, souser.UserId, strconv.Itoa(3600*24*15))

			if reply != "OK" && r2 != 1 {
				out <- ""
				fmt.Printf("save user to redis failed %s %s", reply)
				return
			}
			out <- "OK"
		}()
	}
	return out
}

func (sc *SocketConn) SaveUserToHash(ctx context.Context, key string, souser *socuser.SoconnUser) chan bool {
	out := make(chan bool)
	go func() {
		defer close(out)
		m := souser.StructToMap()
		reply, err := lib.JudgeHmset(key, m, "0")
		if err != nil {
			out <- false
			fmt.Printf("save user to redis failed %s %s", reply, err.Error())
			return
		}
		out <- true
	}()
	return out
}

func (sc *SocketConn) SaveUserToSet(ctx context.Context, key, userid string) chan bool {
	out := make(chan bool)
	go func() {
		defer close(out)

		reply, err := lib.JudgeSadd(key, userid)

		if err != nil {
			out <- false
			fmt.Printf("save user to set failed %s %s", reply, err.Error())
			return
		}
		out <- true
	}()
	return out
}

type RedisSocketConns struct {
	Appid    string
	TraceId  string
	TenantId string
	Cid      string
	ChatFlag string
	Message  string
}

func (rsc RedisSocketConns) GetConnByRedisBroadcast(in chan bool) {
	reply, err := lib.JudgeSmembers(config.OnlineHostConnKey)
	if err != nil {
		in <- false
		fmt.Sprintf("hgetall is failed %s", err.Error())
		return
	}

	c := make(chan string)
	go func() {
		for _, v := range reply {
			c <- v
		}
		close(c)
	}()

	for host := range c {

		go func(host string) {
			args := models.RpcScArgs{
				Appid:    rsc.Appid,
				TraceId:  rsc.TraceId,
				TenantId: rsc.TenantId,
				Host:     host,
				Cid:      rsc.Cid, //myself socket connection is
				ChatFlag: rsc.ChatFlag,
				Message:  rsc.Message,
				From:     utils.GetIntranetIp(),
			}
			ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
			defer cancel()
			in := rpcserv.SendBroadCast(ctx, args)
			select {
			case <-ctx.Done():
				fmt.Println("超时", host)
				return
			case <-in:
				//fmt.Println(ok, host)
			}

		}(host)

	}

	in <- true
}
