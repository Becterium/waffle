package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"waffle/app/user/service/internal/biz"
)

var _ biz.AddressRepo = (*addressRepo)(nil)

type addressRepo struct {
	data *Data
	log  *log.Helper
}

func NewAddressRepo(data *Data, logger log.Logger) biz.AddressRepo {
	return addressRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/address")),
	}
}

func (a2 addressRepo) CreateAddress(ctx *context.Context, a *biz.Address) (*biz.Address, error) {
	//TODO implement me
	panic("implement me")
}

func (a2 addressRepo) GetAddress(ctx *context.Context, id int64) (*biz.Address, error) {
	//TODO implement me
	panic("implement me")
}

func (a2 addressRepo) ListAddress(ctx *context.Context, uid int64) ([]*biz.Address, error) {
	//TODO implement me
	panic("implement me")
}
