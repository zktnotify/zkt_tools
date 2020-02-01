package user

import (
	"github.com/web_zktnotify/internal/app/model"
	"github.com/web_zktnotify/pkg/protobuf"
)

type IUser interface {
	Login(login *protobuf.Login) (*model.User, error)
}

type UserService struct {
}

func (*UserService) Login(login *protobuf.Login) (*model.User, error) {
	return &model.User{
		UserName: "zhangsan",
		Age:      10,
	}, nil
}
