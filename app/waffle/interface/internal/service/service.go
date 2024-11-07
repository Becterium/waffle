package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "waffle/api/waffle/interface/v1"
	"waffle/app/waffle/interface/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewWaffleInterface)

type WaffleInterface struct {
	v1.UnimplementedWaffleInterfaceServer

	log *log.Helper
	uc  *biz.UserUseCase
}

func NewWaffleInterface(logger log.Logger, uc *biz.UserUseCase) *WaffleInterface {
	return &WaffleInterface{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		uc:  uc,
	}
}
