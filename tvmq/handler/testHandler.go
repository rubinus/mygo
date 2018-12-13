package handler

import (
	"strconv"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/lib"
	"code.tvmining.com/tvplay/tvmq/utils"
	"github.com/kataras/iris"
)

type clientPage struct {
	Title  string
	Host   string
	Proto  string
	CIP    string
	Userid string
	Token  string
}

func Add() func() int {
	var demo int = 999
	return func() int {
		return demo + 1
	}
}

func GetAllUsers() chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		reply, _ := lib.JudgeScard(config.ChatRoomName)
		out <- int(reply)
	}()
	return out
}

func TextHandler(ctx iris.Context) {

	f := GetAllUsers()
	nc := strconv.Itoa(<-f + 1000)

	var socketHost string
	var socketProto string
	if config.EnvHost == true {
		socketHost = "localhost:" + strconv.Itoa(config.WebServerPort)
		socketProto = "ws"
	} else {
		socketHost = config.SocketHost
		socketProto = "wss"
	}
	ctx.ViewData("", clientPage{
		"Client Test Page",
		socketHost,
		socketProto,
		utils.GetIntranetIp(),
		nc,
		nc,
	})
	ctx.View("client.html")
}

func TextHandler1(ctx iris.Context) {

	f := GetAllUsers()
	nc := strconv.Itoa(<-f + 1000)

	var socketHost string
	var socketProto string
	if config.EnvHost == true {
		socketHost = "localhost:" + strconv.Itoa(config.WebServerPort)
		socketProto = "ws"
	} else {
		socketHost = config.SocketHost
		socketProto = "wss"
	}
	ctx.ViewData("", clientPage{
		"Client Test Page",
		socketHost,
		socketProto,
		utils.GetIntranetIp(),
		nc,
		nc,
	})
	ctx.View("client1.html")
}
