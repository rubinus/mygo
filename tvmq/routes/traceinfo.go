package routes

import "github.com/kataras/iris"

type indexPage struct {
	Title string
}

func IndexPage(ctx iris.Context) {
	//ctx.WriteString("追踪 ...")
	ctx.ViewData("", indexPage{
		"追踪",
	})
	ctx.View("index.html")
}
