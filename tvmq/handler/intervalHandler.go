package handler

import (
	"strings"
	"time"

	"code.tvmining.com/tvplay/tvmq/middler"
)

type HeartBeat struct {
	Delay    time.Duration
	ChatFlag string
}

func (hb HeartBeat) Run() {
	ch := make(chan string, 1000)

	go func() {

		if hb.Delay < 1e9 {
			hb.Delay = 1e9
		}

		for {

			time.Sleep(hb.Delay)

			if len(rpcSendQueue) > 1000 {
				continue
			}

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

				rse := &rpcSendEntry{
					appid:    appid,
					tenantId: tenantId,
					chatFlag: hb.ChatFlag,
					message:  "1",
				}

				rpcSendQueue <- rse
			}
		}
	}()

}
