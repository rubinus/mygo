package rpcserv

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"net/url"
	"strconv"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/logs"
	"code.tvmining.com/tvplay/tvmq/logs/trace"
	"code.tvmining.com/tvplay/tvmq/models"
	"code.tvmining.com/tvplay/tvmq/rpcserv/client"
	"code.tvmining.com/tvplay/tvmq/utils"
)

func RpcServe() {
	rpc.Register(SconnService{})
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(config.RpcServerPort))

	if err != nil {
		fmt.Println("rpc serve not start ... ", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("rpc serve start is error", err.Error())
			continue
		} else {
			//addr := conn.RemoteAddr()
			//fmt.Printf("rpc serve有数据 %s\n", addr)
		}

		go jsonrpc.ServeConn(conn)
	}
}

type SconnService struct {
}

//接收rpc client的调用，需要提前注册到rpc服务中，方便使用SconnService.Broadcast
func (SconnService) Broadcast(rpcArgs models.RpcScArgs, result *int) error {
	//fmt.Println(rpcArgs)
	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()

	//logger
	traceId := rpcArgs.TraceId
	var logger logs.Logger
	logger = &trace.TraceInfo{
		TraceId:        traceId,
		RecvByRPCServe: utils.GetCurrentTime(13),
	}
	logs.TraceChan <- logger

	in := SendBroadCastReal(ctx, &rpcArgs)
	if r := <-in; r {
		*result = 1
	} else {
		*result = 0
	}
	return nil
}

//发起rpc client调用
func SendBroadCast(ctx context.Context, args models.RpcScArgs) chan bool {

	if config.StartRpcServe == 1 {
		out := make(chan bool)
		go func() {
			defer close(out)

			if args.ChatFlag == config.SocEventAuth ||
				args.ChatFlag == config.SocEventChat ||
				args.ChatFlag == config.SocEventGift ||
				args.ChatFlag == config.SocEventNewAuthUser {
				traceId := args.TraceId
				var logger logs.Logger
				logger = &trace.TraceInfo{
					TraceId:       traceId,
					CallRPCClient: utils.GetCurrentTime(13),
				}
				logs.TraceChan <- logger
			}

			r := client.Call(&args)
			if r == 0 {
				out <- true
			} else {
				out <- false
			}
		}()
		return out
	} else {
		return SendBroadCastReal(ctx, &args)
	}

}

func SendBroadCastReal(ctx context.Context, args *models.RpcScArgs) chan bool {
	//play := strings.NewReader(fmt.Sprintf("cid=%s&chatFlag=%s&message=%s&from=%s", args.Cid, args.ChatFlag, args.Message,args.From))
	out := make(chan bool)
	go func() {
		defer close(out)
		u := fmt.Sprintf("http://%s:%s%s", args.Host, strconv.Itoa(config.SelfServePort), config.RpcCallPath)

		play := url.Values{
			"appid":    {args.Appid},
			"traceId":  {args.TraceId},
			"tenantId": {args.TenantId},
			"cid":      {args.Cid},
			"chatFlag": {args.ChatFlag},
			"message":  {args.Message},
			"from":     {args.From},
		}
		resp, err := http.PostForm(u, play)

		if err != nil {
			fmt.Printf("连接错误：%s\n", err.Error())
			out <- false
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("读取结果错误：%s\n", err.Error())
			out <- false
		}
		//fmt.Println(u, play)
		if string(body) == "1" {
			out <- true
		} else {
			out <- false
		}

	}()

	return out
}
