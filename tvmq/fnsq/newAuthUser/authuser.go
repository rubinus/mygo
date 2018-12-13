package newAuthUser

import (
	"fmt"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/fnsq/nsqService"
	"github.com/nsqio/go-nsq"

	"context"

	"code.tvmining.com/tvplay/tvmq/backend"
	"code.tvmining.com/tvplay/tvmq/logs"
	"code.tvmining.com/tvplay/tvmq/logs/trace"
	"code.tvmining.com/tvplay/tvmq/utils"
	"github.com/bitly/go-simplejson"
	"github.com/kataras/iris/core/errors"
)

type NewAuthUser struct {
	Topic   string
	Channel string
}

func (c NewAuthUser) Consumer() {
	go nsqService.Consumer(c.Topic, c.Channel, config.NsqHostProt, 2, c.Deal)
}

//处理消息
func (c NewAuthUser) Deal(msg *nsq.Message) error {
	//fmt.Println("接收到NSQ", msg.NSQDAddress, "message:", body)

	body, _ := simplejson.NewJson(msg.Body)
	appid, _ := body.Get("Appid").String()
	tenantId, _ := body.Get("TenantId").String()

	if appid == "" {
		appid = config.Minappid
	}
	if tenantId == "" {
		tenantId = config.DefalutTenantId
	}

	var (
		logger  logs.Logger
		traceId string
		userId  string
	)

	//get traceid
	ti := body.Get("TraceInfo")
	if ti != nil {
		traceId, _ = ti.Get("traceId").String()
		userId, _ = ti.Get("userId").String()
		//log trace info to kafka
		logger = &trace.TraceInfo{
			TraceId:     traceId,
			UserId:      userId,
			RecvByQueue: utils.GetCurrentTime(13),
			Ttype:       config.SocEventAuth,
		}
		logs.TraceChan <- logger
	}

	var errStr string
	var reply interface{}
	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()

	body.Set("Message", body.Get("UserInfo"))
	abody, _ := body.MarshalJSON()
	message := string(abody)
	out := backend.BroadcastByRedis(ctx, "", config.SocEventNewAuthUser, message, appid, traceId, tenantId)
	select {
	case <-ctx.Done():
		fmt.Println("超时BroadcastByRedis")
		return nil
	case reply = <-out:
		if reply == nil {
			errStr = "广播失败"
		}
	}

	if errStr != "" {
		fmt.Println(errStr)
		return errors.New(errStr)
	}

	return nil
}
