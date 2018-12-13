package middler

import (
	"sync"
	"time"

	"code.tvmining.com/tvplay/tvmq/allmap"
	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/logs/trace"
	"code.tvmining.com/tvplay/tvmq/midpackage"
	"code.tvmining.com/tvplay/tvmq/models"
	"github.com/kataras/iris/websocket"
)

var am = allmap.AMap{}
var ConnMap = midpackage.GetMap(am)
var CMapMutex = new(sync.RWMutex)

var delay = time.Second * 10

type RequestBody struct {
	Appid     string
	TenantId  string
	Conn      *websocket.Connection
	Message   string
	EventType string
	UserInfo  *models.ResComment
	TraceInfo *trace.TraceInfo
}

type RequestAuth struct {
	ChanBody chan *RequestBody
	WorkChan chan chan *RequestBody
}
type RequestChat struct {
	ChanBody chan *RequestBody
	WorkChan chan chan *RequestBody
}
type RequestGift struct {
	ChanBody chan *RequestBody
	WorkChan chan chan *RequestBody
}
type RequestDiss struct {
	ChanBody chan *RequestBody
	WorkChan chan chan *RequestBody
}

func (qs *RequestAuth) Start() {
	//初始化2个chan chan集装箱
	qs.ChanBody = make(chan *RequestBody)
	qs.WorkChan = make(chan chan *RequestBody) //存放--执行收到的请求
	CreateStart(config.SocEventAuth, qs.ChanBody, qs.WorkChan)
}
func (r *RequestAuth) MakeWorkChan() chan *RequestBody {
	return make(chan *RequestBody)
}
func (r *RequestAuth) WorkReady(w chan *RequestBody) {
	r.WorkChan <- w

}
func (r *RequestAuth) Submit(c *RequestBody) {
	r.ChanBody <- c

}

func (qs *RequestChat) Start() {
	//初始化2个chan chan集装箱
	qs.ChanBody = make(chan *RequestBody)
	qs.WorkChan = make(chan chan *RequestBody) //存放--执行收到的请求
	CreateStart(config.SocEventChat, qs.ChanBody, qs.WorkChan)
}
func (r *RequestChat) MakeWorkChan() chan *RequestBody {
	return make(chan *RequestBody)
}
func (r *RequestChat) WorkReady(w chan *RequestBody) {
	r.WorkChan <- w

}
func (r *RequestChat) Submit(c *RequestBody) {
	r.ChanBody <- c

}

func (qs *RequestGift) Start() {
	//初始化2个chan chan集装箱
	qs.ChanBody = make(chan *RequestBody)
	qs.WorkChan = make(chan chan *RequestBody) //存放--执行收到的请求
	CreateStart(config.SocEventGift, qs.ChanBody, qs.WorkChan)
}
func (r *RequestGift) MakeWorkChan() chan *RequestBody {
	return make(chan *RequestBody)
}
func (r *RequestGift) WorkReady(w chan *RequestBody) {
	r.WorkChan <- w

}
func (r *RequestGift) Submit(c *RequestBody) {
	r.ChanBody <- c

}

func (qs *RequestDiss) Start() {
	//初始化2个chan chan集装箱
	qs.ChanBody = make(chan *RequestBody)
	qs.WorkChan = make(chan chan *RequestBody) //存放--执行收到的请求
	CreateStart(config.SocEventDiss, qs.ChanBody, qs.WorkChan)
}
func (r *RequestDiss) MakeWorkChan() chan *RequestBody {
	return make(chan *RequestBody)
}
func (r *RequestDiss) WorkReady(w chan *RequestBody) {
	r.WorkChan <- w
}
func (r *RequestDiss) Submit(c *RequestBody) {
	r.ChanBody <- c
}

func CreateStart(f string, cb chan *RequestBody, wc chan chan *RequestBody) {
	//在调度中开goruntine，只有一个并开始循环
	go func() {
		//创建队列
		var queueRequest []*RequestBody
		var queueWork []chan *RequestBody
		for {
			var activeRequest *RequestBody
			var activeWork chan *RequestBody
			if len(queueRequest) > 0 && len(queueWork) > 0 { //存放的，和执行的都有，各取出一个
				activeRequest = queueRequest[0]
				activeWork = queueWork[0]
			}
			select {
			case r := <-cb: //接收到有了请求，先放到队列中
				//fmt.Println("--从浏览器接到请求放到channel，并读到了channel中的值",r)
				queueRequest = append(queueRequest, r)
			case w := <-wc: //work接收到了，也放到队列中
				queueWork = append(queueWork, w)
			case activeWork <- activeRequest: //把当前的请求给处理的work chan 从队列中去掉
				queueRequest = queueRequest[1:]
				queueWork = queueWork[1:]
			case <-time.After(delay):
				//fmt.Println(f, "调度队列queueRequest=", len(queueRequest), "--queueWork=", len(queueWork))
			}
		}
	}()
}
