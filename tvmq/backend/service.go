package backend

import (
	"context"
	"fmt"

	"code.tvmining.com/tvplay/tvmq/dealconn"
	"code.tvmining.com/tvplay/tvmq/lib"
	"code.tvmining.com/tvplay/tvmq/models"
	"code.tvmining.com/tvplay/tvmq/models/socuser"
	"github.com/mitchellh/mapstructure"
)

func BroadcastByRedis(ctx context.Context, cid, chatFlag, message, appid, traceId, tenantId string) chan bool {
	out := make(chan bool)
	go func() {
		defer close(out)
		rsc := dealconn.RedisSocketConns{
			Appid:    appid,
			TraceId:  traceId,
			TenantId: tenantId,
			Cid:      cid,
			ChatFlag: chatFlag,
			Message:  message,
		}
		rsc.GetConnByRedisBroadcast(out)
	}()

	return out
}

func GetUserinfo(ctx context.Context, userid string) chan *models.ResComment {
	//先从redis中取出userid的信息
	out := make(chan *models.ResComment)
	go func() {
		defer close(out)

		var userinfo socuser.Userinfo
		reply, err := lib.JudgeHgetall("USER:" + userid)
		mapstructure.Decode(reply, &userinfo)

		if err != nil {
			out <- nil
			fmt.Printf("redis hgetall失败 %s", err.Error())
			return
		}

		c := models.ResComment{
			Userid:     userid,
			Headimgurl: userinfo.Headimgurl,
			Nickname:   userinfo.Nickname,
			City:       userinfo.City,
			Province:   userinfo.Province,
			Country:    userinfo.Country,
		}
		out <- &c
	}()
	return out
}

func SaveActvityToRedis(ctx context.Context, key, value string, ttl int) chan bool {
	out := make(chan bool)
	go func() {

		reply, err := lib.JudgeSetex(key, ttl, value)

		if err != nil {
			fmt.Printf("setex失败 %s", err.Error())
		}
		if reply == "OK" {
			out <- true
		} else {
			out <- false
		}
	}()
	return out
}
