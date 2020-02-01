package web

import (
	"github.com/kataras/iris"
)

func Register(app *iris.Application) {
	// home
	app.Get("/", home)
	app.Get("/login", login)
}

func home(ctx iris.Context) {
	ctx.View("index.html")
}

func login(ctx iris.Context){
	ctx.View("login.html")
}
