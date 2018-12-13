package main

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/hooto/hlog4g/hlog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
	"github.com/rcrowley/go-metrics"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/fnsq"
	"code.tvmining.com/tvplay/tvmq/handler"
	"code.tvmining.com/tvplay/tvmq/kafka"
	"code.tvmining.com/tvplay/tvmq/middler"
	"code.tvmining.com/tvplay/tvmq/routes"
	"code.tvmining.com/tvplay/tvmq/rpcserv"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			hlog.Printf("fatal", "Server/Panic %s", err)
		}
		hlog.Flush()
	}()

	metrics.UseNilMetrics = true

	app := iris.New()

	var pre string
	if config.UsePreAbsPath == 1 {
		prefix, _ := filepath.Abs(filepath.Dir(os.Args[0]) + "/")
		pre = prefix + "/templates"
	} else {
		pre = "./templates"
	}

	app.RegisterView(iris.HTML(pre, ".html")) // select the html engine to serve templates

	app.Any("/socket/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(websocket.ClientSource)
	})

	app.StaticWeb("/js", "./static/js")                 // serve our custom javascript code
	app.StaticWeb("/headimgurl", "./static/headimgurl") // serve our custom javascript code

	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})

	authEngine, chatEngine, giftEngine, dissEngine := middler.MiddRun() //启动中间件 队列调度引擎

	e := handler.ConnHandler{
		Rauth: authEngine,
		Rchat: chatEngine,
		Rgift: giftEngine,
		Rdiss: dissEngine,
	}
	ws.OnConnection(e.HandleFun)

	app.Get("/socket/broadcast", ws.Handler())

	//集中调用路由
	routes.Index(app)

	//发送心跳
	hb := handler.HeartBeat{
		Delay:    10 * time.Second,
		ChatFlag: config.SocEventHeartBeat,
	}
	handler.NewHandlerService(hb)

	//发送在线人数
	online := handler.Online{
		Delay:    5 * time.Second,
		ChatFlag: config.SocEventOnline,
	}
	handler.NewHandlerService(online)

	//最近发送的消息
	rm := handler.RecentMsg{
		Delay:    10 * time.Second,
		ChatFlag: config.SocEventHeartBeat,
	}
	handler.NewHandlerService(rm)

	//web服务器保存连接
	hostConn := handler.HostConn{
		Delay: 60 * time.Second,
		Key:   config.OnlineHostConnKey,
	}
	handler.NewHandlerService(hostConn)

	go handler.BroadcastForOnlineUsersWorker()
	//go middler.WebSocketListKeeper()

	if config.StartKafkaTestData == 1 {
		go handler.Start() //仅供测试数据
		//kafkaCons.GetKafkaInfo(config.KafkaHostPort)
	}

	if config.StartKafkaConsumer == 1 {
		kafka.Consumers()
	}
	if config.StartNsqConsumer == 1 {
		fnsq.Consumers()
	}

	if config.StartRpcServe == 1 {
		go func() {
			rpcserv.RpcServe() //异步启动一个rpc serve
		}()
	}

	app.Run(iris.Addr(":" + strconv.Itoa(config.SelfServePort)))
}
