package main

import (
	"fmt"
	"mygo/rpcdemo"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn is error")
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
