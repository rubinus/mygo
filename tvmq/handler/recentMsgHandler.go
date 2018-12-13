package handler

import (
	"fmt"
	"time"

	"strconv"

	"code.tvmining.com/tvplay/tvmq/allmap"
	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/lib"
)

type RecentMsg struct {
	Delay    time.Duration
	ChatFlag string
}

func (rm RecentMsg) Run() {
	ch := make(chan string, 100)
	go func() {
		for range time.Tick(rm.Delay) {
			allmap.Appids.Mu.RLock()
			ids := allmap.Appids.Ids
			allmap.Appids.Mu.RUnlock()
			for appid, _ := range ids {
				ch <- appid
			}
		}
	}()
	go func() {
		for {
			for appid := range ch {
				c := int(rm.GetRecentMsgCard(appid))
				if c <= config.RecentMsgCount {
					continue
				}
				stop := strconv.Itoa(c - config.RecentMsgCount - 1)
				key := fmt.Sprintf("%s:%s", config.RecentMsgKey, appid)
				if _, err := lib.JudgeZremrangebyrank(key, "0", stop); err != nil {
					fmt.Println("Zremrangebyrank is err :", err)
				}
			}
		}
	}()
}

func (rm RecentMsg) GetRecentMsgCard(appid string) int64 {
	key := fmt.Sprintf("%s:%s", config.RecentMsgKey, appid)
	//fmt.Println(key,"----key----")
	card, err := lib.JudgeZcard(key)
	if err != nil {
		fmt.Printf(key, "redis 取所有最近消息失败 %s", err.Error())
		return 0
	}
	return card
}
