package handler

import (
	"fmt"
	"sync"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/logs"
	"code.tvmining.com/tvplay/tvmq/logs/trace"
	"code.tvmining.com/tvplay/tvmq/middler"
	"code.tvmining.com/tvplay/tvmq/utils"
	"github.com/hooto/hlog4g/hlog"
)

type rpcSendEntry struct {
	appid    string
	tenantId string
	traceId  string
	chatFlag string
	message  string
	from     string
	resType  string
}

type Servicer interface {
	Run()
}

var (
	rpcSendMu    sync.Mutex
	rpcSendQueue = make(chan *rpcSendEntry, 2000)
)

func NewHandlerService(servicer Servicer) {
	servicer.Run()
}

//
func BroadcastForOnlineUsersWorker() {

	for msgEntry := range rpcSendQueue {
		broadcastForOnlineUsersWorkerSendEntry(msgEntry)
	}
}

func broadcastForOnlineUsersWorkerSendEntry(msgEntry *rpcSendEntry) {

	go func() {
		if err := recover(); err != nil {
			hlog.Printf("fatal", "Worker/Panic %s", err)
		}
	}()

	if msgEntry.message == "" {
		return
	}

	message := msgEntry.message
	mid := fmt.Sprintf("%s:%s", msgEntry.appid, msgEntry.tenantId)
	middler.ImChatWebSockets.Mu.RLock()
	cmap := middler.ImChatWebSockets.Items[mid]
	middler.ImChatWebSockets.Mu.RUnlock()

	//fmt.Println(cmap, "---message---", mid)

	for _, wsEntry := range cmap {

		if err := wsEntry.Conn.Emit(msgEntry.chatFlag, message); err != nil {
			hlog.Printf("error", "ERR send to %s, msg %s", wsEntry.UserId, message)
		} else {
			// hlog.Printf("info", "OK send to %s, msg %s", wsEntry.UserId, message)
			if msgEntry.chatFlag == config.SocEventGift ||
				msgEntry.chatFlag == config.SocEventChat ||
				msgEntry.chatFlag == config.SocEventNewAuthUser {
				traceId := msgEntry.traceId
				var logger logs.Logger
				logger = &trace.TraceInfo{
					TraceId:      traceId,
					SendToClient: utils.GetCurrentTime(13),
				}
				logs.TraceChan <- logger
			}

		}
	}

}
