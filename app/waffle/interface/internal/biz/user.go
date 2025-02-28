package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"strconv"
	v1 "waffle/api/waffle/interface/v1"
	"waffle/app/waffle/interface/internal/conf"
	"waffle/app/waffle/interface/internal/pkg/util"
)

const (
	JwtMapClaimsUserID = "userId"
)

var (
	ErrPasswordInvalid = errors.New("password invalid")
	ErrUsernameInvalid = errors.New("username invalid")
	ErrUserNotFound    = errors.New("user not found")
)

type User struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(username string,
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
	Find(ctx context.Context, id uint64) (*User, error)
	FindByName(ctx context.Context, username string) (*User, error)
	Save(ctx context.Context, u *User) (uint64, error)
	VerifyPassword(ctx context.Context, username, password string) error
	Ping(ctx context.Context) (string, error)
}

type UserUseCase struct {
	key  string
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(conf *conf.Auth, repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		key:  conf.GetJwtKey(),
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/interface")),
	}
}

func (c UserUseCase) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	//todo:是否要考虑幂等性问题呢
	_, err := c.repo.FindByName(ctx, req.Username)
	if err == nil {
		return nil, v1.ErrorRegisterFailed("this username has exist")
	}
	user, err := NewUser(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	userId, err := c.repo.Save(ctx, &user)
	if err != nil {
		return nil, v1.ErrorRegisterFailed("save user fail, error: %s", err.Error())
	}
	return &v1.RegisterReply{
		Id: userId,
	}, nil
}

func (c UserUseCase) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	user, err := c.repo.FindByName(ctx, req.Username)

	if err != nil {
		return nil, v1.ErrorLoginFailed(err.Error())
	}

	err = c.repo.VerifyPassword(ctx, user.Username, req.Password)
	if err != nil {
		return nil, v1.ErrorLoginFailed("password not match ,error : %s", err.Error())
	}

	claims := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, jwtv5.MapClaims{
		JwtMapClaimsUserID: strconv.FormatUint(user.Id, 10),
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

func (c UserUseCase) Ping(ctx context.Context) (*v1.PingRPCReply, error) {

	ctx, err := util.UnMarshalTokeToMetadata(ctx, util.PrefixLocalMetadata, util.MarshalUserId)
	if err != nil {
		return nil, err
	}

	reply, err := c.repo.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &v1.PingRPCReply{Message: reply}, nil
}
