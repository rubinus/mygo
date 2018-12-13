package nsgift

import (
	"context"
	"errors"
	"fmt"

	"code.tvmining.com/tvplay/tvmq/backend"
	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/fnsq/comm"
	"code.tvmining.com/tvplay/tvmq/fnsq/nsqService"
	"code.tvmining.com/tvplay/tvmq/logs"
	"code.tvmining.com/tvplay/tvmq/logs/trace"
	"code.tvmining.com/tvplay/tvmq/models/gift"
	"code.tvmining.com/tvplay/tvmq/utils"
	"github.com/bitly/go-simplejson"
	"github.com/nsqio/go-nsq"
)

type Gift struct {
	Topic   string
	Channel string
}

func (g Gift) Consumer() {
	go nsqService.Consumer(g.Topic, g.Channel, config.NsqHostProt, 2, g.Deal)
}

//处理消息
func (g Gift) Deal(msg *nsq.Message) error {
	//fmt.Println("接收到NSQ", msg.NSQDAddress, "message:", message)

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
			Ttype:       config.SocEventGift,
		}
		logs.TraceChan <- logger
	}

	rbMessage, _ := body.Get("Message").Bytes()
	json, _ := simplejson.NewJson(rbMessage)
	headimgurl, _ := json.Get("headimgurl").String()
	nickname, _ := json.Get("nickname").String()
	from_userid, _ := json.Get("userid").String()
	token, _ := json.Get("token").String()
	icon, _ := json.Get("icon").String()
	pictures, _ := json.Get("pictures").String()
	giftid, _ := json.Get("giftid").String()
	giftname, _ := json.Get("giftname").String()
	count, _ := json.Get("count").Int()
	points, _ := json.Get("points").Int()
	gs := gift.Gift{
		Appid:      appid,
		Userid:     from_userid,
		TraceId:    traceId,
		Headimgurl: headimgurl,
		Nickname:   nickname,
		Giftid:     giftid,
		Giftname:   giftname,
		Icon:       icon,
		Pictures:   pictures,
		Count:      count,
		Points:     points,
	}

	//异步减掉积分
	var errStr string
	//url := fmt.Sprintf("%s%s",config.PointHost,config.PointPostPath)
	url := fmt.Sprintf("%s%s?userid=%s&token=%s", config.PointHost, config.PointPostPath, from_userid, token)

	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()

	sout := backend.SendPoint(ctx, url, &gs)
	select {
	case <-ctx.Done():
		errStr = "超时SendPoint"
	case ok := <-sout:
		if ok != "OK" {
			errStr = ok
		}
	}

	//logger
	//log trace info to kafka
	logger = &trace.TraceInfo{
		TraceId:        traceId,
		UserId:         userId,
		RecvByCallGift: utils.GetCurrentTime(13),
		Ttype:          config.SocEventGift,
		ErrMsg:         errStr,
	}
	logs.TraceChan <- logger

	if errStr != "" {
		fmt.Println(errStr)
		return errors.New(errStr)
	}

	if ti != nil {
		ti.Set("recvByCallGift", utils.GetCurrentTime(13))
	}

	mout := comm.SaveMessageToMongodb(ctx, &gs)

	select {
	case <-ctx.Done():
		errStr = "超时SaveMessageToMongodb"
	case ok := <-mout:
		if ok != "ok" {
			errStr = "保存mongo failed"
		}
	}
	//log trace info to kafka
	logger = &trace.TraceInfo{
		TraceId:         traceId,
		UserId:          userId,
		RecvBySaveMongo: utils.GetCurrentTime(13),
		Ttype:           config.SocEventGift,
		ErrMsg:          errStr,
	}
	logs.TraceChan <- logger

	if errStr != "" {
		fmt.Println(errStr)
		return errors.New(errStr)
	}

	message := string(msg.Body)
	out := backend.BroadcastByRedis(ctx, "", config.SocEventGift, message, appid, traceId, tenantId)
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
