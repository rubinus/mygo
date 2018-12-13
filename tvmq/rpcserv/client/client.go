package client

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"
	"sync"

	"github.com/bitly/go-simplejson"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/models"
	"code.tvmining.com/tvplay/tvmq/utils"
)

var (
	rpc_mu      sync.Mutex
	rpc_clients = map[string]*rpc.Client{}
)

//真实的rpc client产生并调用
func Call(rpcArgs *models.RpcScArgs) int {

	var (
		reuse = false
		addr  = rpcArgs.Host + ":" + strconv.Itoa(config.RpcServerPort)
		err   error
	)

	rpc_mu.Lock()

	client, ok := rpc_clients[addr]
	if !ok {
		client, err = jsonrpc.Dial("tcp", addr)
		if err != nil {
			rpc_mu.Unlock()
			fmt.Println("RPC client : ", err.Error())
			return 0
		}
		rpc_clients[addr] = client
	} else {
		reuse = true
	}
	rpc_mu.Unlock()

	body, _ := simplejson.NewJson([]byte(rpcArgs.Message))
	ti := body.Get("TraceInfo")
	if ti != nil {
		ti.Set("callRPCClient", utils.GetCurrentTime(13))
	}
	abody, _ := body.MarshalJSON()

	rpcArgs.Message = string(abody)
	//fmt.Println("call rpc befor ..", string(abody))

	var result int
	err = client.Call("SconnService.Broadcast", rpcArgs, &result)
	if err != nil {
		if reuse {
			rpc_mu.Lock()
			if _, ok := rpc_clients[addr]; ok {
				client.Close()
				delete(rpc_clients, addr)
			}
			rpc_mu.Unlock()

			// retry
			client, err = jsonrpc.Dial("tcp", addr)
			if err != nil {
				fmt.Println("RPC client : ", err.Error())
				return 0
			}

			err = client.Call("SconnService.Broadcast", rpcArgs, &result)
			if err != nil {
				fmt.Println("RPC client : ", err.Error())
				return 0
			}

			rpc_mu.Lock()
			rpc_clients[addr] = client
			rpc_mu.Unlock()

		} else {
			fmt.Printf("错误%s", err.Error())
		}
	} else {
		//fmt.Println(result)
	}
	return result
}
