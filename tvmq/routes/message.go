package routes

import (
	"fmt"
	"sync"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/lib"
	"code.tvmining.com/tvplay/tvmq/utils"
	"github.com/bitly/go-simplejson"
	"github.com/kataras/iris"
)

type MsgList struct {
	mlist       []*simplejson.Json
	lastRequest int
}

type MapMsgList struct {
	mu   sync.RWMutex
	msgs map[string]MsgList
}

var Msgs = MapMsgList{
	msgs: make(map[string]MsgList),
}

func (ml *MapMsgList) getMsg(mid string, currTime int) (*MsgList, error) {
	var reply MsgList
	ml.mu.RLock()
	reply = ml.msgs[mid]
	ml.mu.RUnlock()

	if len(reply.mlist) > 0 && reply.lastRequest+3 >= currTime {
		//fmt.Println("have and 3 second")
		return &reply, nil
	}
	err := ml.setMsg(mid, currTime)
	reply = ml.msgs[mid]
	return &reply, err
}
func (ml *MapMsgList) setMsg(mid string, currTime int) error {
	ml.mu.Lock()
	defer ml.mu.Unlock()
	key := fmt.Sprintf("%s:%s", config.RecentMsgKey, mid)
	reply, err := lib.JudgeZrange(key, "0", "-1", -1)
	if err != nil {
		return err
	}
	var s []*simplejson.Json
	for _, v := range reply {
		body, _ := simplejson.NewJson([]byte(v))
		body.Del("appid")
		body.Del("traceId")
		body.Del("tenantId")
		s = append(s, body)
	}
	result := ml.msgs[mid]
	result.mlist = s
	result.lastRequest = currTime
	ml.msgs[mid] = result
	//fmt.Println("3 second after or no data in the mem ...")
	return nil
}

func MessageList(ctx iris.Context) {
	appid := ctx.URLParam("appid")
	tenantId := ctx.URLParam("tenantid")
	if appid == "" {
		appid = config.Minappid
	}
	if tenantId == "" {
		tenantId = config.DefalutTenantId
	}
	mid := fmt.Sprintf("%s:%s", appid, tenantId)
	if s, err := Msgs.getMsg(mid, int(utils.GetCurrentTime(10))); err != nil {
		e := fmt.Sprintf("zrange is failed: %s", err.Error())
		ctx.JSONP(iris.Map{"status": 401, "msg": e})
	} else {
		ctx.JSONP(iris.Map{"status": 200, "data": s.mlist})
	}

}
