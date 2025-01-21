package biz

import (
	"context"
	"errors"
	"math/rand"
	v1 "waffle/api/user/service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// GreeterRepo is a Greater repo.
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	VerifyPassword(ctx context.Context, user *User) (bool, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	InitCache(ctx context.Context) (string, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	out, err := uc.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (uc *UserUseCase) Get(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) FindByUsername(ctx context.Context, username string) (*User, error) {
	return uc.repo.FindByUsername(ctx, username)
}

func (uc *UserUseCase) Save(ctx context.Context, u *User) (*v1.SaveUserReply, error) {
	user := &User{
		Id:       rand.Int63(),
		Username: u.Username,
		Password: u.Password,
	}
	_, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		// todo: handle error
		return nil, err
	}
	return &v1.SaveUserReply{
		Id: user.Id,
	}, nil
}

func (uc *UserUseCase) VerifyPassword(ctx context.Context, u *User) (bool, error) {
	return uc.repo.VerifyPassword(ctx, u)
}

func (uc *UserUseCase) InitCache(ctx context.Context) (*v1.InitCacheReply, error) {
	message, err := uc.repo.InitCache(ctx)
	return &v1.InitCacheReply{
		Message: message,
	}, err
}
