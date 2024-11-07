package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "waffle/api/waffle/interface/v1"
)

type User struct {
	Id       int64
	Name     string
	Password string
}

type UserRepo interface {
	Find(ctx context.Context, id int64) (*User, error)
	FindByName(ctx context.Context, name string) (*User, error)
	Save(ctx context.Context, u *User) error
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func (c UserUseCase) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	return nil, nil
}

func (c UserUseCase) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	return nil, nil
}

func (c UserUseCase) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	return nil, nil
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/interface")),
	}
}
