package middler

import (
	"sync"

	"github.com/kataras/iris/websocket"
)

type WebSocketEntry struct {
	//Appid string
	//Id     string
	UserId string
	Conn   websocket.Connection
	//Close  bool
}

type WebSocketList struct {
	Mu    sync.RWMutex
	Items map[string]map[string]*WebSocketEntry
}

var (
	ImChatWebSockets = WebSocketList{ //全局map=map[appid:商户号id]map[连接id]具体连接信息体
		Mu:    sync.RWMutex{},
		Items: make(map[string]map[string]*WebSocketEntry),
	}
	err error
)

func (it *WebSocketList) Entry(mid, uid, wsid string, ws websocket.Connection) *WebSocketEntry {

	it.Mu.Lock()
	defer it.Mu.Unlock()

	wseMap := make(map[string]*WebSocketEntry)
	wse := &WebSocketEntry{
		Conn: ws,
		//Id:     wsid,
		UserId: uid,
	}
	wseMap[wsid] = wse

	if _, ok := it.Items[mid]; !ok {
		it.Items[mid] = wseMap
	} else {
		we := it.Items[mid]
		we[wsid] = wse
	}

	//for _, v := range it.Items {
	//	if v.Id == id {
	//		return v
	//	}
	//}

	return wseMap[wsid]
}

//func (it *WebSocketList) Sync(id string, ws websocket.Connection) {
//
//	it.mu.Lock()
//	defer it.mu.Unlock()
//
//	for _, v := range it.Items {
//		if v.Id == id {
//			return
//		}
//	}
//
//	it.Items = append(it.Items, &WebSocketEntry{
//		Id:   id,
//		Conn: ws,
//	})
//
//}

func (it *WebSocketList) Close(mid, id string) {

	it.Mu.Lock()
	defer it.Mu.Unlock()

	if _, ok := it.Items[mid]; ok {
		we := it.Items[mid]
		delete(we, id)
	}

	//for _, v := range it.Items {
	//	if v.Id == id {
	//		v.Close = true
	//		return
	//	}
	//}

}

//func (it *WebSocketList) Del(id string) {
//
//	it.mu.Lock()
//	defer it.mu.Unlock()
//
//	for i, v := range it.Items {
//		if v.Id == id {
//			it.Items = append(it.Items[:i], it.Items[i+1:]...)
//			return
//		}
//	}
//}

//func WebSocketListKeeper() {
//
//	for {
//		time.Sleep(10e9)
//		webSocketListKeeperAction()
//	}
//}

//func webSocketListKeeperAction() {
//
//	dels := []string{}
//
//	for _, v := range ImChatWebSockets.Items {
//		if v.Close {
//			dels = append(dels, v.Id)
//			continue
//		}
//
//		// v.Conn.Send(`{"type": "ping"}`)
//	}
//
//	for _, v := range dels {
//		ImChatWebSockets.Del(v)
//	}
//
//	hlog.Printf("info", "web sockets %d", len(ImChatWebSockets.Items))
//}
