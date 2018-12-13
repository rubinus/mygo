package middler

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/hooto/hlog4g/hlog"
	"github.com/kataras/iris/websocket"

	"code.tvmining.com/tvplay/tvmq/backend"
	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/dealconn"
	"code.tvmining.com/tvplay/tvmq/fnsq/nsqproduce"
	"code.tvmining.com/tvplay/tvmq/lib"
	"code.tvmining.com/tvplay/tvmq/logs"
	"code.tvmining.com/tvplay/tvmq/logs/trace"
	"code.tvmining.com/tvplay/tvmq/models"
	"code.tvmining.com/tvplay/tvmq/models/socuser"
	"code.tvmining.com/tvplay/tvmq/utils"
	"github.com/mitchellh/mapstructure"
)

func DoWorker(in chan *RequestBody) {
	rb := <-in
	//fmt.Printf("msg from socket :%s %s\n",rb.Message,rb.EventType)
	switch rb.EventType {
	case config.SocEventAuth:
		rb.DealAuth()
	case config.SocEventChat:
		rb.DealChat()
	case config.SocEventGift:
		rb.DealGift()
	case config.SocEventDiss:
		rb.DealDiss()
	}

}

func (rb *RequestBody) DealAuth() {
	c := *rb.Conn
	msg := rb.Message

	var logger logs.Logger
	var tenantId string
	var traceId string
	var appid string
	var errStr string

	//验证用户的请求输入
	aid, userid, nickname, headimgurl, _, tid, taid, err := rb.checkAuthInput(msg)

	if tid == "" {
		traceId = strconv.Itoa((int)(utils.GetCurrentTime(19)))
	} else {
		traceId = tid
	}

	if aid == "" {
		appid = config.Minappid
	} else {
		appid = aid
	}

	if taid == "" {
		tenantId = config.DefalutTenantId
	} else {
		tenantId = taid
	}

	if err != nil {
		//log trace info to kafka
		logger = &trace.TraceInfo{
			TraceId:      traceId,
			UserId:       userid,
			TenantId:     tenantId,
			RecvByClient: utils.GetCurrentTime(13),
			Ttype:        config.SocEventAuth,
			ErrMsg:       err.Error(),
		}
		logs.TraceChan <- logger

		res := models.NewResErrBody(201, err.Error(), "")
		SendMessage(c, config.SocEventAuth, res)
		return
	}

	//log trace info to kafka
	logger = &trace.TraceInfo{
		TraceId:      traceId,
		UserId:       userid,
		TenantId:     tenantId,
		RecvByClient: utils.GetCurrentTime(13),
		Ttype:        config.SocEventAuth,
	}
	logs.TraceChan <- logger

	if v, ok := logger.(*trace.TraceInfo); ok {
		rb.TraceInfo = v
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()

	c.SetValue("userid", userid)
	c.SetValue("appid", appid)
	c.SetValue("tenantId", tenantId)
	rb.Appid = appid
	rb.TenantId = tenantId
	rsuser := &models.ResComment{
		Userid:     userid,
		Nickname:   nickname,
		Headimgurl: headimgurl,
	}
	rb.UserInfo = rsuser

	mid := fmt.Sprintf("%s:%s", appid, tenantId)
	if ws := ImChatWebSockets.Entry(mid, userid, c.ID(), c); ws == nil {
		return
	}

	//存储用户信息到redis
	sc := dealconn.SocketConn{
		make(chan *socuser.SoconnUser),
	}
	suser := socuser.SoconnUser{
		UserId: userid,
		Host:   utils.GetIntranetIp(),
		Cid:    c.ID(),
	}
	sc.ReceConn(&suser)
	chatroom := fmt.Sprintf("%s:%s:%s", config.ChatRoomName, rb.Appid, tenantId)
	ch := sc.SaveUserToHashAndSet(ctx, chatroom)
	select {
	case <-ctx.Done():
		errStr = "超时SaveToCache"
		fmt.Println(errStr)
	case reply := <-ch:
		if reply != "OK" {
			errStr = reply
		}
	}

	if errStr != "" {
		//logger
		logger = &trace.TraceInfo{
			TraceId:              traceId,
			UserId:               userid,
			TenantId:             tenantId,
			SaveUserToHashAndSet: utils.GetCurrentTime(13),
			Ttype:                config.SocEventAuth,
			ErrMsg:               errStr,
		}
		logs.TraceChan <- logger

		res := models.NewResErrBody(201, errStr, userid)
		SendMessage(c, config.SocEventAuth, res)
		return
	}

	//logger
	logger = &trace.TraceInfo{
		TraceId:              traceId,
		UserId:               userid,
		TenantId:             tenantId,
		SaveUserToHashAndSet: utils.GetCurrentTime(13),
		Ttype:                config.SocEventAuth,
	}
	logs.TraceChan <- logger

	//发送auth事件
	go func(rb *RequestBody) {
		res := models.ResAuthBody{
			Status:  200,
			Content: models.AuthBody{rb.UserInfo.Userid},
			From:    utils.GetIntranetIp(),
		}
		reply, _ := models.Marshal(res)
		e := SendMessage(c, config.SocEventAuth, reply)

		if e != nil {
			//log trace info to kafka
			logger = &trace.TraceInfo{
				TraceId:      rb.TraceInfo.TraceId,
				UserId:       userid,
				TenantId:     tenantId,
				SendToClient: utils.GetCurrentTime(13),
				Ttype:        config.SocEventAuth,
				ErrMsg:       e.Error(),
			}
		} else {
			//log trace info to kafka
			logger = &trace.TraceInfo{
				TraceId:      rb.TraceInfo.TraceId,
				UserId:       userid,
				SendToClient: utils.GetCurrentTime(13),
				Ttype:        config.SocEventAuth,
			}
		}
		logs.TraceChan <- logger

	}(rb)

	//发送到nsq队列
	ns := nsqproduce.Nservice{
		Topic: config.TopicNewAuthUser,
	}
	newMsg, _ := models.Marshal(&rb)
	pout := ns.Producer(newMsg)
	select {
	case r := <-pout:
		if !r {
			errStr = "发送NSQ失败"
		}
	}
	if errStr != "" {
		fmt.Println(errStr, "=======")
		return
	}

	//check当前是否是互动
	out := checkIsActivity(ctx, appid, tenantId)
	select {
	case <-ctx.Done():
		errStr = "超时checkIsActivity"
		fmt.Println(errStr)
	case reply := <-out:
		if reply != "" {
			errStr = reply
		}
	}

	if errStr != "" {
		//logger
		logger = &trace.TraceInfo{
			TraceId:         traceId,
			UserId:          userid,
			CheckIsActivity: utils.GetCurrentTime(13),
			Ttype:           config.SocEventAuth,
			ErrMsg:          errStr,
		}
		logs.TraceChan <- logger

		fmt.Println(errStr, "=======")
		SendMessage(c, config.SocEventActivity, errStr)
		return
	}

	//logger
	logger = &trace.TraceInfo{
		TraceId:         traceId,
		UserId:          userid,
		CheckIsActivity: utils.GetCurrentTime(13),
		Ttype:           config.SocEventAuth,
	}
	logs.TraceChan <- logger

}

func (rb *RequestBody) DealChat() {
	c := *rb.Conn
	msg := rb.Message
	var logger logs.Logger
	var errStr string
	var traceId string

	if c.GetValue("userid") == nil || c.GetValue("appid") == nil {
		//logger
		errStr = "userid/appid is empty"
		return
	}
	userid := c.GetValue("userid").(string)
	appid := c.GetValue("appid").(string)
	tenantId := c.GetValue("tenantId").(string)
	rb.Appid = appid
	rb.TenantId = tenantId
	//fmt.Println("--DealChat--",c.ID(),userid)

	if userid == "" {
		errStr := "请先进行auth认证"
		//logger
		logger = &trace.TraceInfo{
			TraceId:      traceId,
			UserId:       userid,
			TenantId:     tenantId,
			RecvByClient: utils.GetCurrentTime(13),
			Ttype:        config.SocEventChat,
			ErrMsg:       errStr,
		}
		logs.TraceChan <- logger

		res := models.NewResErrBody(202, errStr, "")
		SendMessage(c, config.SocEventChatReply, res)
		return
	}

	json, _ := simplejson.NewJson([]byte(msg))
	message, _ := json.Get("message").String()
	from_userid, _ := json.Get("userid").String()
	tid, _ := json.Get("traceid").String()
	if from_userid == "" || message == "" {
		errStr := "发送人或消息不能为空"
		//logger
		logger = &trace.TraceInfo{
			TraceId:      traceId,
			UserId:       userid,
			RecvByClient: utils.GetCurrentTime(13),
			Ttype:        config.SocEventChat,
			ErrMsg:       errStr,
		}
		logs.TraceChan <- logger

		res := models.NewResErrBody(201, errStr, message)
		SendMessage(c, config.SocEventChatReply, res)
		return
	}

	if tid == "" {
		traceId = strconv.Itoa((int)(utils.GetCurrentTime(19)))
	} else {
		traceId = tid
	}

	//logger
	logger = &trace.TraceInfo{
		TraceId:      traceId,
		UserId:       userid,
		RecvByClient: utils.GetCurrentTime(13),
		Ttype:        config.SocEventChat,
	}
	logs.TraceChan <- logger

	if v, ok := logger.(*trace.TraceInfo); ok {
		rb.TraceInfo = v
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()

	//是否敏感
	if config.UseFilterWord == 1 {
		body := backend.FilterBody{
			Items: []backend.Items{
				{
					Id:     from_userid,
					UserId: from_userid,
					Text:   message,
				},
			},
		}
		fch := backend.SendFilter(ctx, config.FilterHost, body)
		select {
		case <-ctx.Done():
			errStr = "超时SendFilter"
		case result := <-fch:
			if result < 0.90 && result > 0 {
				errStr = "请你文明点!"
			}
		}
		if errStr != "" {
			//logger
			logger = &trace.TraceInfo{
				TraceId:    traceId,
				UserId:     userid,
				FilterWord: utils.GetCurrentTime(13),
				Ttype:      config.SocEventChat,
				ErrMsg:     errStr,
			}
			logs.TraceChan <- logger

			res := models.NewResErrBody(201, errStr, message)
			fmt.Println(errStr)
			SendMessage(c, config.SocEventChatReply, res)
			return
		}

		//logger
		logger = &trace.TraceInfo{
			TraceId:    traceId,
			UserId:     userid,
			FilterWord: utils.GetCurrentTime(13),
			Ttype:      config.SocEventChat,
		}
		logs.TraceChan <- logger

	}

	//是否可以发送
	if config.UseSendMsgLimit == 1 {
		out := checkCanSend(ctx, from_userid)
		select {
		case <-ctx.Done():
			errStr = "提前退出"
		case ok := <-out:
			if !ok {
				errStr = "发送太频繁了"
			}
		}
		if errStr != "" {
			//logger
			logger = &trace.TraceInfo{
				TraceId:  traceId,
				UserId:   userid,
				MsgLimit: utils.GetCurrentTime(13),
				Ttype:    config.SocEventChat,
				ErrMsg:   errStr,
			}
			logs.TraceChan <- logger

			res := models.NewResErrBody(201, errStr, message)
			c.Emit(config.SocEventChatReply, res)
			return
		}

		//logger
		logger = &trace.TraceInfo{
			TraceId:  traceId,
			UserId:   userid,
			MsgLimit: utils.GetCurrentTime(13),
			Ttype:    config.SocEventChat,
		}
		logs.TraceChan <- logger
	}

	ns := nsqproduce.Nservice{
		Topic: config.TopicChat,
	}

	newMsg, _ := models.Marshal(&rb)
	pout := ns.Producer(newMsg)
	select {
	case r := <-pout:
		if !r {
			errStr = "发送NSQ失败"
		}
	}
	if errStr != "" {
		fmt.Println(errStr, "=======")
	} else {
		//fmt.Println("发送NSQ successful=========")

	}
}

func (rb *RequestBody) DealGift() {
	c := *rb.Conn
	msg := rb.Message

	var logger logs.Logger
	var traceId string
	var errStr string

	_, _, _, _, _, tid, _, err := rb.checkAuthInput(msg)
	if err != nil {
		res := models.NewResErrBody(201, err.Error(), "")
		SendMessage(c, config.SocEventChatReply, res)
		return
	}

	if tid == "" {
		traceId = strconv.Itoa((int)(utils.GetCurrentTime(19)))
	} else {
		traceId = tid
	}

	userid := c.GetValue("userid").(string)
	appid := c.GetValue("appid").(string)
	tenantId := c.GetValue("tenantId").(string)

	if userid == "" {
		errStr = "请先进行auth认证"
		//logger
		logger = &trace.TraceInfo{
			TraceId:      traceId,
			UserId:       userid,
			RecvByClient: utils.GetCurrentTime(13),
			Ttype:        config.SocEventGift,
			ErrMsg:       errStr,
		}
		logs.TraceChan <- logger

		res := models.NewResErrBody(202, errStr, "")
		SendMessage(c, config.SocEventChatReply, res)
		return
	}
	rb.Appid = appid
	rb.TenantId = tenantId

	json, err := simplejson.NewJson([]byte(msg))
	if err != nil {
		errStr := "消息解析错误"
		//logger
		logger = &trace.TraceInfo{
			TraceId:      traceId,
			UserId:       userid,
			RecvByClient: utils.GetCurrentTime(13),
			Ttype:        config.SocEventGift,
			ErrMsg:       errStr,
		}
		logs.TraceChan <- logger

		res := models.NewResErrBody(201, errStr, "")
		SendMessage(c, config.SocEventChatReply, res)
		return
	}

	from_userid, _ := json.Get("userid").String()
	token, _ := json.Get("token").String()
	giftid, _ := json.Get("giftid").String()
	count, _ := json.Get("count").Int()
	points, _ := json.Get("points").Int()

	if from_userid == "" || giftid == "" || token == "" {
		errStr = "发送人、giftid、token不能为空"
		//logger
		logger = &trace.TraceInfo{
			TraceId:      traceId,
			UserId:       userid,
			RecvByClient: utils.GetCurrentTime(13),
			Ttype:        config.SocEventGift,
			ErrMsg:       errStr,
		}
		logs.TraceChan <- logger

		res := models.NewResErrBody(201, errStr, giftid)
		SendMessage(c, config.SocEventGiftReply, res)
		return
	}

	//logger
	logger = &trace.TraceInfo{
		TraceId:      traceId,
		UserId:       userid,
		RecvByClient: utils.GetCurrentTime(13),
		Ttype:        config.SocEventGift,
	}
	logs.TraceChan <- logger

	if v, ok := logger.(*trace.TraceInfo); ok {
		rb.TraceInfo = v
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()

	if config.UseSendMsgLimit == 1 {
		out := checkCanSend(ctx, from_userid)
		select {
		case <-ctx.Done():
			errStr = "超时提前退出"
		case ok := <-out:
			if !ok {
				errStr = "发送太频繁了"
			}
		}
		if errStr != "" {
			//logger
			logger = &trace.TraceInfo{
				TraceId:  traceId,
				UserId:   userid,
				MsgLimit: utils.GetCurrentTime(13),
				Ttype:    config.SocEventGift,
				ErrMsg:   errStr,
			}
			logs.TraceChan <- logger

			res := models.NewResErrBody(201, errStr, "")
			c.Emit(config.SocEventGiftReply, res)
			return
		}

		//logger
		logger = &trace.TraceInfo{
			TraceId:  traceId,
			UserId:   userid,
			MsgLimit: utils.GetCurrentTime(13),
			Ttype:    config.SocEventGift,
		}
		logs.TraceChan <- logger

	}

	url := fmt.Sprintf("%s%s%s&token=%s", config.PointHost, config.PointGetPath, from_userid, token)
	pointCh := backend.GetPoint(ctx, url)
	select {
	case <-ctx.Done():
		errStr = "超时GetPoint"
		fmt.Println(errStr)
	case r := <-pointCh:
		if r != -1 && r < count*points {
			errStr = "积分不够"
		}
	}

	if errStr != "" {
		//logger
		logger = &trace.TraceInfo{
			TraceId:        traceId,
			UserId:         userid,
			RecvByCallGift: utils.GetCurrentTime(13),
			Ttype:          config.SocEventGift,
			ErrMsg:         errStr,
		}
		logs.TraceChan <- logger

		res := models.NewResErrBody(201, errStr, "")
		//c.Emit(config.SocEventGiftReply, res)
		SendMessage(c, config.SocEventGiftReply, res)
		return
	}

	ns := nsqproduce.Nservice{
		Topic: config.TopicGift,
	}
	newMsg, _ := models.Marshal(&rb)
	pout := ns.Producer(newMsg)
	select {
	case r := <-pout:
		if !r {
			errStr = "发送NSQ失败"
		}
	}
	if errStr != "" {
		fmt.Println(errStr, "=======")
	}
}

func (rb *RequestBody) DealDiss() {
	c := *rb.Conn

	var userid, appid, tenantId string

	if c.GetValue("userid") != nil {
		userid = c.GetValue("userid").(string)
	}
	if c.GetValue("appid") != nil {
		appid = c.GetValue("appid").(string)
	}
	if c.GetValue("tenantId") != nil {
		tenantId = c.GetValue("tenantId").(string)
	}

	mid := fmt.Sprintf("%s:%s", appid, tenantId)
	ImChatWebSockets.Close(mid, c.ID())

	if userid == "" && appid == "" {
		if wsEntry := ImChatWebSockets.Entry(mid, userid, c.ID(), c); wsEntry != nil {
			//userid = wsEntry.UserId
			//appid = wsEntry.Appid
		}
	}

	if userid != "" && appid != "" {
		//准备删除redis中的连接
		key := fmt.Sprintf("%s:%s:%s", config.ChatPreixForUser, userid, tenantId)
		key2 := fmt.Sprintf("%s:%s:%s", config.ChatRoomName, appid, tenantId)
		if !delChatKeyAndSetMember(key, key2, userid) {
			hlog.Printf("warn", "ERR Close WebSocket %s", c.ID())
		}
	}
}

func SendMessage(c websocket.Connection, chatFlag, message string) error {
	//if entry := ImChatWebSockets.Entry(c.ID()); entry != nil {
	//	if !entry.Close {
	//		if err := entry.Conn.Emit(chatFlag, message); err != nil {
	//			entry.Close = true
	//			return err
	//		}
	//	}
	//}
	//return nil
	return c.Emit(chatFlag, message)
}

func (r *RequestBody) checkAuthInput(msg string) (string, string, string, string, string, string, string, error) {
	//fmt.Println("msg form socket : ", msg)

	json, err := simplejson.NewJson([]byte(msg))
	if err != nil {
		errStr := "消息解析错误"
		return "", "", "", "", "", "", "", errors.New(errStr)
	}
	userId, _ := json.Get("userid").String()
	token, _ := json.Get("token").String()

	uid, nickname, headimgurl, err := utils.TokenValid(token, config.TokenSecretKey)
	if err != nil {
		return "", "", "", "", "", "", "", err
	}
	if uid != userId {
		errStr := "id and token not match"
		return "", "", "", "", "", "", "", errors.New(errStr)
	}

	appid, _ := json.Get("appid").String()
	traceId, _ := json.Get("traceid").String()
	tenantId, _ := json.Get("tenantid").String()

	return appid, userId, nickname, headimgurl, token, traceId, tenantId, nil
}

func (r *RequestBody) checkUserIdAndToken(ctx context.Context, key, token string) chan socuser.Userinfo {
	out := make(chan socuser.Userinfo)
	go func() {
		defer close(out)

		var userinfo socuser.Userinfo
		reply, err := lib.JudgeHgetall(key)
		mapstructure.Decode(reply, &userinfo)
		if err != nil {
			out <- userinfo
			fmt.Printf("获取redis数据失败%s%s", reply, err.Error())
		}

		if token == userinfo.Token {
			out <- userinfo
		} else {
			out <- socuser.Userinfo{}
		}

	}()

	return out
}

func checkIsActivity(ctx context.Context, appid, tenantId string) chan string {
	key := fmt.Sprintf("%s:%s:%s", config.CurrentActivity, appid, tenantId)
	out := make(chan string)
	go func() {
		reply, err := lib.JudgeGet(key)

		//time.Sleep(5 * time.Second)
		if reply == "" || err != nil {
			out <- ""
			//fmt.Printf("当前无互动%s%s\n", reply, err)
		} else {
			json, err := simplejson.NewJson([]byte(reply))
			if err != nil {
				out <- ""
				return
			}
			content := json.Get("content")
			if content == nil {
				out <- ""
				return
			}
			start_time, err := content.Get("start_time").Int()
			if err != nil {
				out <- ""
				return
			}
			countdown, _ := json.Get("content").Get("countdown").String()
			countdownI, _ := strconv.Atoi(countdown)
			sub := time.Now().UnixNano()/1e9 - int64(start_time)
			if int64(countdownI)-sub > 0 {

				content.Set("countdown", strconv.Itoa(int(int64(countdownI)-sub)))
				r, _ := json.MarshalJSON()

				out <- string(r)
			} else {
				out <- ""
			}
		}

	}()
	return out
}

func checkCanSend(ctx context.Context, userid string) chan bool {
	key := fmt.Sprintf("%s:%s", config.LastSendPreix, userid)

	out := make(chan bool)

	go func() {
		reply, err := lib.JudgeGet(key)
		now := time.Now().UnixNano() / 1e6

		if reply == "" && err == nil {
			//out <- true
			saveLastSendInfo(out, userid, config.SendMsgLimitSecond, now)
			//fmt.Printf("最近没有发送过%s%s\n", reply, err)
		} else {
			d, _ := strconv.Atoi(reply)
			if now-int64(d) < int64(1000*config.SendMsgLimitSecond) { //3秒内不能频繁发送
				out <- false
			} else {
				saveLastSendInfo(out, userid, config.SendMsgLimitSecond, now)
			}
		}

	}()
	return out
}

func saveLastSendInfo(in chan bool, userid string, ttl int, lasttime int64) {

	key := fmt.Sprintf("%s:%s", config.LastSendPreix, userid)

	value := strconv.Itoa(int(lasttime))

	reply, err := lib.JudgeSetex(key, ttl, value)

	if err != nil {
		in <- false
		fmt.Printf("hmset失败 %s", err.Error())
	}
	if reply == "OK" {
		in <- true
	} else {
		in <- false
	}
}

func delChatKeyAndSetMember(key, setKey, member string) bool {

	reply, r2 := lib.JudgeDelKeyAndSetMember(key, setKey, member)

	if reply != 1 && r2 != 1 {
		return false
	}
	return true
}

func delChatKey(key string) chan bool {
	out := make(chan bool)
	go func() {

		reply, err := lib.JudgeDelKey(key)

		if err != nil {
			out <- false
			fmt.Printf("断开连接删除sokcet失败%s%s", reply, err.Error())
		} else {
			out <- true
		}

	}()
	return out
}

func delChatSet(key, userid string) chan bool {
	out := make(chan bool)
	go func() {

		reply, err := lib.JudgeSrem(key, userid)

		if err != nil {
			out <- false
			fmt.Printf("断开连接删除%s失败%s%s", key, reply, err.Error())
		} else {
			out <- true
		}
	}()
	return out
}
