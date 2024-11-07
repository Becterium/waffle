package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"
	"waffle/app/waffle/interface/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
		sg:   &singleflight.Group{},
	}
}

func (r *userRepo) Find(ctx context.Context, id int64) (*biz.User, error) {
	return nil, nil
}

func (r *userRepo) FindByName(ctx context.Context, name string) (*biz.User, error) {
	return nil, nil
}

func (r *userRepo) Save(ctx context.Context, u *biz.User) error {
	return nil
}
