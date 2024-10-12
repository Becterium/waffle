package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Address struct {
	Id       int64
	Name     string
	Mobile   string
	Address  string
	PostCode string
}

type AddressRepo interface {
	CreateAddress(ctx *context.Context, a *Address) (*Address, error)
	GetAddress(ctx *context.Context, id int64) (*Address, error)
	ListAddress(ctx *context.Context, uid int64) ([]*Address, error)
}

type AddressUseCase struct {
	repo *AddressRepo
	log  *log.Helper
}

func NewAddressUseCase(repo *AddressRepo, logger log.Logger) *AddressUseCase {
	return &AddressUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/address")),
	}
}
