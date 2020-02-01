package middleware

import (
	"fmt"
	"github.com/kataras/iris/context"
	"github.com/spf13/viper"
	"time"
)

func NewPreHandler() context.Handler {
	return func(ctx context.Context) {
		startAt := time.Now().UnixNano() / 1000000
		ctx.Values().Set("startAt", startAt)
		ctx.Header("X-Powered-By", fmt.Sprintf("web_zktnotify/v%s", viper.GetString("version")))
		ctx.Gzip(viper.GetBool("server.gzip"))
		ctx.Next()
	}
}
