package routes

import (
	"code.tvmining.com/tvplay/tvmq/handler"
	"github.com/kataras/iris"
)

func Index(app *iris.Application) {

	app.Get("/", IndexPage)

	app.Post("/rpcsend", handler.RpcsendHandler)

	app.Get("/test", handler.TextHandler)

	app.Get("/test1", handler.TextHandler1)

	app.Get("/messages/list", MessageList)

	app.Get("/wxauth/minlogin", Minlogin)

	app.Post("/wxauth/saveMinUser", SaveMinUser)

	app.Get("/wxauth/checktoken", ChcekTokenRoute)

	app.Get("/wxauth/getUserinfo", GetUserinfo)

}
