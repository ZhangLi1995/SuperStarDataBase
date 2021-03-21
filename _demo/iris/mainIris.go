package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

func main() {

	app := iris.New()
	app.Use(logger.New())

	htmlEngine := iris.HTML("./_demo/iris", ".html")
	app.RegisterView(htmlEngine)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello World! -- from iris")
	})
	app.Get("/hello", func(ctx iris.Context) {
		ctx.ViewData("Title", "测试页面")
		ctx.ViewData("Content", "Hello World! -- from iris")
		ctx.View("hello.html")
	})

	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}
