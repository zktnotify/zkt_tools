package service

import "github.com/web_zktnotify/internal/app/service/user"

type Service struct {
	user.UserService
}

func NewService() *Service {
	return &Service{}
}
