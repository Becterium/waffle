package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
	v1 "waffle/api/waffle/interface/v1"
	"waffle/app/waffle/interface/internal/conf"
)

var (
	ErrPasswordInvalid = errors.New("password invalid")
	ErrUsernameInvalid = errors.New("username invalid")
	ErrUserNotFound    = errors.New("user not found")
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(
	username string,
	password string,
) (User, error) {

	// check username
	if len(username) <= 0 {
		return User{}, ErrUsernameInvalid
	}
	// check password
	if len(password) <= 0 {
		return User{}, ErrPasswordInvalid
	}

	return User{
		Username: username,
		Password: password,
	}, nil
}

type UserRepo interface {
	Find(ctx context.Context, id int64) (*User, error)
	FindByName(ctx context.Context, username string) (*User, error)
	Save(ctx context.Context, u *User) error
	VerifyPassword(ctx context.Context, u *User, password string) error
}

type UserUseCase struct {
	key  string
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(conf *conf.Auth, repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		key:  conf.ApiKey,
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/interface")),
	}
}

func (c UserUseCase) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	_, err := c.repo.FindByName(ctx, req.Username)
	if !errors.Is(err, ErrUserNotFound) {
		return nil, v1.ErrorRegisterFailed("username already exists")
	}
	user, err := NewUser(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	err = c.repo.Save(ctx, &user)
	if err != nil {
		return nil, v1.ErrorRegisterFailed("save user fail, error: %s", err.Error())
	}
	return &v1.RegisterReply{
		Id: user.Id,
	}, nil
}

func (c UserUseCase) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	user, err := c.repo.FindByName(ctx, req.Username)
	if errors.Is(err, ErrUserNotFound) {
		return nil, v1.ErrorRegisterFailed("user not found, error: %s", err.Error())
	}

	err = c.repo.VerifyPassword(ctx, user, req.Password)
	if err != nil {
		return nil, v1.ErrorLoginFailed("password not match")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
	})
	signedString, err := claims.SignedString([]byte(c.key))
	if err != nil {
		return nil, v1.ErrorLoginFailed("generate token failed,error: %s", err.Error())
	}
	return &v1.LoginReply{
		Token: signedString,
	}, nil
}

func (c UserUseCase) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	return &v1.LogoutReply{}, nil
}
