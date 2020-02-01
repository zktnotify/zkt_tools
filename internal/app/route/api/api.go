package api

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/spf13/viper"
	"github.com/web_zktnotify/internal/app/controller"
	"github.com/web_zktnotify/internal/app/service"
)

func Register(app *iris.Application) {
	hero.Register(service.NewService())
	app.Get("api/v1", banner)
	app.PartyFunc("api/v1", apiParty)
}

func apiParty(apiParty iris.Party) {
	ctr := controller.UserController{}
	// login
	apiParty.Post("/login", hero.Handler(ctr.Login))
}

func banner(ctx iris.Context) {
	ctx.ViewData("message", fmt.Sprintf(` _____                 _   
|  ___|               | |  
| |____  ___ __   ___ | |_ 
|  __\ \/ / '_ \ / _ \| __|
| |___>  <| |_) | (_) | |_ 
\____/_/\_\ .__/ \___/ \__|
          | |              
          |_|              Version:%s`, viper.GetString("version")))
	ctx.View("banner.html")
}
