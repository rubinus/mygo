package handler

import (
	"fmt"
	"time"

	"strings"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/lib"
	"code.tvmining.com/tvplay/tvmq/middler"
	"code.tvmining.com/tvplay/tvmq/models"
	"code.tvmining.com/tvplay/tvmq/utils"
)

type Online struct {
	Delay    time.Duration
	ChatFlag string
}

func (ol Online) Run() {
	ch := make(chan string, 1000)
	go func() {
		if ol.Delay < 1e9 {
			ol.Delay = 1e9
		}
		for {

			time.Sleep(ol.Delay)

			if len(rpcSendQueue) > 1000 {
				continue
			}

			//allmap.Appids.Mu.RLock()
			//ids := allmap.Appids.Ids
			//allmap.Appids.Mu.RUnlock()
			//for appid, _ := range ids {
			//	ch <- appid
			//}

			middler.ImChatWebSockets.Mu.RLock()
			items := middler.ImChatWebSockets.Items
			for k, _ := range items {
				ch <- k
			}
			middler.ImChatWebSockets.Mu.RUnlock()

		}
	}()
	go func() {
		for {
			for mid := range ch {
				s := strings.Split(mid, ":")
				appid := s[0]
				tenantId := s[1]
				num := ol.GetOnlineNumber(mid)

				if num == "" {
					continue
				}

				rse := &rpcSendEntry{
					appid:    appid,
					tenantId: tenantId,
					chatFlag: ol.ChatFlag,
					message:  num,
				}

				rpcSendQueue <- rse
			}
		}
	}()
}

func (ol Online) GetOnlineNumber(mid string) string {
	chatroom := fmt.Sprintf("%s:%s", config.ChatRoomName, mid)
	//fmt.Println("--chatroom--", chatroom)
	card, err := lib.JudgeScard(chatroom)
	if err != nil {
		fmt.Printf("redis 取所有在线用户失败 %s", err.Error())
		return ""
	}
	if card == 0 {
		//fmt.Printf("%s, no user \n",chatroom)
		return ""
	}
	reply := models.ResOnlineBody{
		Status:  200,
		Content: models.OnlineBody{Count: card},
		From:    utils.GetIntranetIp(),
	}
	res, err := models.Marshal(reply)
	if err == nil {
		return res
	}

	return ""
}
