package service

import (
	"github.com/google/wire"
	v1 "waffle/api/user/service/v1"
	"waffle/app/user/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	v1.UnimplementedUserServer

	uc *biz.UserUseCase
	ac *biz.AddressUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}
