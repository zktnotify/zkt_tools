package middleware

import (
	"fmt"
	"github.com/go-errors/errors"
	"github.com/kataras/iris/context"
)

func NewRecover() context.Handler {
	return func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				msg := fmt.Sprintf("发生panic异常: %v\n", errors.Wrap(err, 2).ErrorStack())
				fmt.Println(msg)
			}
		}()
		ctx.Next()
	}
}
