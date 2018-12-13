package chat

import (
	"fmt"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/fnsq/nsqService"
	"github.com/nsqio/go-nsq"

	"context"
	"strconv"

	"code.tvmining.com/tvplay/tvmq/backend"
	"code.tvmining.com/tvplay/tvmq/fnsq/comm"
	"code.tvmining.com/tvplay/tvmq/lib"
	"code.tvmining.com/tvplay/tvmq/logs"
	"code.tvmining.com/tvplay/tvmq/logs/trace"
	"code.tvmining.com/tvplay/tvmq/models"
	"code.tvmining.com/tvplay/tvmq/models/comment"
	"code.tvmining.com/tvplay/tvmq/utils"
	"github.com/bitly/go-simplejson"
	"github.com/kataras/iris/core/errors"
)

type Chat struct {
	Topic   string
	Channel string
}

func (c Chat) Consumer() {
	go nsqService.Consumer(c.Topic, c.Channel, config.NsqHostProt, 2, c.Deal)
}

//处理消息
func (c Chat) Deal(msg *nsq.Message) error {

	body, _ := simplejson.NewJson(msg.Body)
	appid, _ := body.Get("Appid").String()
	tenantId, _ := body.Get("TenantId").String()

	//fmt.Println("接收到NSQ", appid, msg.NSQDAddress, "message:", string(msg.Body))

	if appid == "" {
		appid = config.Minappid
	}
	if tenantId == "" {
		tenantId = config.DefalutTenantId
	}

	var logger logs.Logger
	var (
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
			Ttype:       config.SocEventChat,
		}
		logs.TraceChan <- logger
	}

	rbMessage, _ := body.Get("Message").Bytes()
	json, _ := simplejson.NewJson(rbMessage)

	ms, _ := json.Get("message").String()
	headimgurl, _ := json.Get("headimgurl").String()
	nickname, _ := json.Get("nickname").String()
	from_userid, _ := json.Get("userid").String()

	coms := comment.Comment{
		Appid:      appid,
		TraceId:    traceId,
		TenantId:   tenantId,
		Userid:     from_userid,
		Headimgurl: headimgurl,
		Nickname:   nickname,
		Content:    ms,
	}
	var errStr string

	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()

	score := strconv.Itoa(int(utils.GetCurrentTime(13)))
	rmsg, _ := models.Marshal(coms)
	rekey := fmt.Sprintf("%s:%s:%s", config.RecentMsgKey, appid, tenantId)
	if _, err := lib.JudgeZadd(rekey, score, rmsg); err != nil {
		errStr = "save redis zset  failed"
		fmt.Println(errStr)
		return errors.New(errStr)
	}

	mout := comm.SaveMessageToMongodb(ctx, &coms)

	select {
	case <-ctx.Done():
		errStr = "超时SaveMessageToMongodb"
	case ok := <-mout:
		if ok != "ok" {
			errStr = "保存mongo failed"
		}
	}

	if errStr != "" {
		fmt.Println(errStr)
		return errors.New(errStr)
	}

	if ti != nil {
		traceId, _ = ti.Get("traceId").String()
		//log trace info to kafka
		logger = &trace.TraceInfo{
			TraceId:         traceId,
			UserId:          userId,
			RecvBySaveMongo: utils.GetCurrentTime(13),
			Ttype:           config.SocEventChat,
		}
		logs.TraceChan <- logger
	}

	message := string(msg.Body)
	out := backend.BroadcastByRedis(ctx, "", config.SocEventChat, message, appid, traceId, tenantId)
	select {
	case <-ctx.Done():
		errStr = "超时BroadcastByRedis"
	case ok := <-out:
		if !ok {
			errStr = "广播失败"
		}
	}

	if errStr != "" {
		fmt.Println(errStr)
		return errors.New(errStr)
	}

	return nil
}
