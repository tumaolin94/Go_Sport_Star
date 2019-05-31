package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	htmlEngine := iris.HTML("./", ".html")
	app.RegisterView(htmlEngine)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString( "Hello world!")

	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.ViewData("Title", "Test Page")
		ctx.ViewData("Content", "Hello World, from template")

		ctx.View("hello.html")


	})


	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}
