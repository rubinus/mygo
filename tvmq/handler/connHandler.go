package handler

import (
	"github.com/kataras/iris/websocket"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/middler"
	"code.tvmining.com/tvplay/tvmq/utils"
)

type ConnHandler struct {
	C     *middler.RequestBody
	Rauth middler.Submiter //处理auth事件调度器
	Rchat middler.Submiter //处理chat事件调度器
	Rgift middler.Submiter //处理gift事件调度器
	Rdiss middler.Submiter //处理diss事件调度器
}

func (conn *ConnHandler) NewConnHandle(c *websocket.Connection) *ConnHandler {
	return &ConnHandler{
		C: &middler.RequestBody{
			Conn: c,
		},
		Rauth: conn.Rauth,
		Rchat: conn.Rchat,
		Rgift: conn.Rgift,
		Rdiss: conn.Rdiss,
	}
}

func (conn *ConnHandler) HandleFun(c websocket.Connection) {

	conn = conn.NewConnHandle(&c)

	c.OnDisconnect(conn.DealDissConn)

	//middler.ImChatWebSockets.Sync(c.ID(), c)

	c.Emit(config.SocEventIP, utils.GetIntranetIp())

	c.On(config.SocEventAuth, conn.Judge(config.SocEventAuth))

	c.On(config.SocEventChat, conn.Judge(config.SocEventChat))

	c.On(config.SocEventGift, conn.Judge(config.SocEventGift))

	c.On(config.SocEventHeartBeat, conn.DealHeartBeatEvent)

}

func (conn *ConnHandler) DealHeartBeatEvent(msg string) {
	//fmt.Println(config.SocEventHeartBeat,conn.Userid,msg)
}

func (conn *ConnHandler) Judge(t string) func(msg string) {
	switch t {
	case config.SocEventAuth:
		return conn.DealAuthEvent
	case config.SocEventChat:
		return conn.DealChatEvent
	case config.SocEventGift:
		return conn.DealGiftEvent
	}
	return nil
}

func (conn *ConnHandler) DealAuthEvent(msg string) {
	rb := middler.RequestBody{
		Conn:      conn.C.Conn,
		Message:   msg,
		EventType: config.SocEventAuth,
	}
	conn.Rauth.Submit(&rb) //把请求的body(socket connection及message提交到channel中)

}

func (conn *ConnHandler) DealChatEvent(msg string) {
	rb := middler.RequestBody{
		Conn:      conn.C.Conn,
		Message:   msg,
		EventType: config.SocEventChat,
	}

	//c := *(rb.Conn)
	//fmt.Println("\n--DealChatEvent--",c.ID(),c.GetValue("userid"))

	conn.Rchat.Submit(&rb) //把请求的body(socket connection及message提交到channel中)
}

func (conn *ConnHandler) DealGiftEvent(msg string) {
	//fmt.Println("msg from socket ",msg,conn.C.ID(),conn.Userid)
	rb := middler.RequestBody{
		Conn:      conn.C.Conn,
		Message:   msg,
		EventType: config.SocEventGift,
	}
	conn.Rgift.Submit(&rb) //把请求的body(socket connection及message提交到channel中)
}

func SendMessage(c websocket.Connection, chatFlag, message string) {
	//c.To(config.ChatRoomName).Emit(chatFlag,message)
	c.Emit(chatFlag, message)
}

func (conn *ConnHandler) DealDissConn() {
	rb := middler.RequestBody{
		Conn:      conn.C.Conn,
		EventType: config.SocEventDiss,
	}
	conn.Rdiss.Submit(&rb) //把请求的body(socket connection及message提交到channel中)
}
