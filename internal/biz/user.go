package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	Id       int64
	Username string
	Password string
}

// GreeterRepo is a Greater repo.
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	VerifyPassword(ctx context.Context, user *User) (bool, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
}

type UserUsecase struct {
	ur  UserRepo
	log *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: repo, log: log.NewHelper(logger)}
}
