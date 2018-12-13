package handler

import (
	"github.com/kataras/iris"

	"code.tvmining.com/tvplay/tvmq/msgtype"
)

func RpcsendHandler(ctx iris.Context) {

	rpcSendMu.Lock()
	defer rpcSendMu.Unlock()

	if len(rpcSendQueue) > 1000 {
		ctx.Writef("0")
	} else {

		appid := ctx.PostValue("appid")
		tenantId := ctx.PostValue("tenantId")
		chatFlag := ctx.PostValue("chatFlag")
		message := ctx.PostValue("message")
		from := ctx.PostValue("from")
		//组织返回消息体
		newMsg, traceId := msgtype.CheckMessageType(chatFlag, message, from)

		rse := &rpcSendEntry{
			appid:    appid,
			chatFlag: chatFlag,
			message:  newMsg,
			from:     from,
			traceId:  traceId,
			tenantId: tenantId,
			//cid := ctx.PostValue("cid")
		}

		rpcSendQueue <- rse

		ctx.Writef("1")
	}
}
