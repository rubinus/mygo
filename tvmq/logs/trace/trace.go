package trace

import (
	"fmt"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/fnsq/nsqproduce"
	"code.tvmining.com/tvplay/tvmq/utils"
)

type TraceInfo struct {
	ErrMsg               string `json:"errMsg,omitempty"`
	Ttype                string `json:"tType"`
	TraceId              string `json:"traceId"`
	TenantId             string `json:"tenantId"`
	UserId               string `json:"userId"`
	RecvByClient         int64  `json:"recvByClient,omitempty"`
	FilterWord           int64  `json:"filterWord,omitempty"`
	GetUserInfo          int64  `json:"getUserInfo,omitempty"`
	SaveUserToHashAndSet int64  `json:"saveUserToHashAndSet,omitempty"`
	CheckIsActivity      int64  `json:"checkIsActivity,omitempty"`
	MsgLimit             int64  `json:"msgLimit,omitempty"`
	RecvByQueue          int64  `json:"recvByQueue,omitempty"`
	RecvByNode           int64  `json:"recvByNode,omitempty"`
	RecvByCallGift       int64  `json:"recvByCallGift,omitempty"`
	RecvBySaveMongo      int64  `json:"recvBySaveMongo,omitempty"`
	CallRPCClient        int64  `json:"callRPCClient,omitempty"`
	RecvByRPCServe       int64  `json:"recvByRPCServe,omitempty"`
	SendToClient         int64  `json:"sendToClient,omitempty"`
}

func (t *TraceInfo) Info() {
	if config.TraceLog == 0 {
		return
	}
	go func() {
		var errStr string
		msg, _ := utils.Marshal(t)
		ns := nsqproduce.Nservice{
			Topic: config.TraceInfo,
		}
		pout := ns.Producer(msg)
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

	}()
}
