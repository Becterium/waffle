package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"waffle/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (u *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	return &biz.User{
		Id:       0,
		Username: "nani",
		Password: "bsdnaklj",
	}, nil
}
func (u *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	return &biz.User{
		Id:       0,
		Username: "heheh",
		Password: "andjka",
	}, nil
}
func (u *userRepo) VerifyPassword(ctx context.Context, user *biz.User) (bool, error) {
	return false, nil
}
func (u *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	return nil, nil
}
