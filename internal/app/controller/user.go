package controller

import (
	"errors"
	"github.com/kataras/iris"
	"github.com/web_zktnotify/internal/app/errcode"
	"github.com/web_zktnotify/internal/app/middleware"
	"github.com/web_zktnotify/internal/app/model"
	"github.com/web_zktnotify/internal/app/service/user"
	"github.com/web_zktnotify/pkg/protobuf"
)

type UserController struct {
}

func (*UserController) Login(ctx iris.Context, iUser user.IUser) *model.Result {
	login := new(protobuf.Login)
	if err := ctx.ReadForm(login); err != nil {
		// TODO log
		return errcode.ERR_PARAM_ERROR.Result()
	}
	if err := checkUserLoginParam(login); err != nil {
		// TODO log
		return errcode.ERR_PARAM_ERROR.Result()
	}
	user, err := iUser.Login(login)
	if err != nil {
		// TODO log
		return errcode.ERR_SYSTEM_ERROR.Result()
	}
	token := middleware.CreateJWTToken(user)
	return &model.Result{
		Status: 0,
		Data:   model.Authorization{Token: token},
		Msg:    "ok",
	}
}

func checkUserLoginParam(login *protobuf.Login) error {
	if login.UserName == "" {
		return errors.New("invalid username")
	}
	if login.Password == "" {
		return errors.New("invalid password")
	}
	return nil
}
