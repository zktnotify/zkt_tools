package initialize

import (
	"context"
	"fmt"
	"github.com/kataras/iris"
	"github.com/spf13/viper"
	"github.com/web_zktnotify/internal/app/middleware"
	"github.com/web_zktnotify/internal/app/route/api"
	"github.com/web_zktnotify/internal/app/route/web"
	"time"
)

func SetupServer() {
	// new app and register api
	app := newApp(web.Register, api.Register)
	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// Shutdown
		app.Shutdown(ctx)
	})

	port := viper.GetString("server.port")
	host := viper.GetString("server.host")
	// run
	app.Run(
		iris.Addr(fmt.Sprintf("%s:%s", host, port)),
		//按下CTRL/CMD+C时错误的服务器：
		iris.WithoutServerError(iris.ErrServerClosed),
		//启用更快的json序列化和优化：
		iris.WithOptimizations,
	)
}

func newApp(appRouteFunc ...func(app *iris.Application)) *iris.Application {
	app := iris.New()
	// before handler
	app.Use(middleware.NewPreHandler())
	// global error handle
	app.Use(middleware.NewRecover())
	app.RegisterView(iris.HTML("./static/html", ".html"))
	app.StaticWeb("/css", "./static/css")
	app.StaticWeb("/images", "./static/images")
	app.StaticWeb("/js", "./static/js")
	app.Favicon("./static/favicon.ico", "/favicon.ico")
	app.Logger().SetLevel(viper.GetString("logLevel"))
	if appRouteFunc != nil && len(appRouteFunc) != 0 {
		for _, item := range appRouteFunc {
			item(app)
		}
	}
	return app
}
