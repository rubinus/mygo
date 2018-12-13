package allmap

import (
	"sync"
	"time"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/models/appid"
	"github.com/kataras/iris/websocket"
)

type AMap struct {
	WebsConn *websocket.Connection
	UserId   string
}

func (c AMap) Reader() map[string]AMap {
	return make(map[string]AMap)
}

type AllAppids struct {
	Mu  sync.RWMutex
	Ids map[string]string
}

var Appids AllAppids

//init appids
func init() {
	initAppids()
	go func() {
		for {
			select {
			case <-time.Tick(3 * time.Minute):
				initAppids()
			}
		}
	}()
}

func initAppids() {
	ch := GetAllAppids()
	ids := <-ch
	Appids.Mu.Lock()
	Appids.Ids = make(map[string]string)
	aids := Appids.Ids
	if len(ids) > 0 {
		for _, v := range ids {
			aids[v.Appid] = v.Secret
		}
	} else {
		//init appid
		aids[config.Minappid] = config.Minsecret
	}
	//fmt.Println(Appids.Ids)
	Appids.Mu.Unlock()
}

func GetAllAppids() chan []appid.Appid {
	out := make(chan []appid.Appid)
	go func() {
		defer close(out)
		var u []appid.Appid
		u, err := appid.GetAllAppids()
		if err != nil {
			out <- u
			return
		}
		out <- u
	}()
	return out
}
