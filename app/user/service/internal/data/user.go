package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"waffle/app/user/service/internal/biz"
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
	result := u.data.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
func (u *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	user := &biz.User{
		Id: id,
	}
	result := u.data.db.Find(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func (u *userRepo) VerifyPassword(ctx context.Context, user *biz.User) (bool, error) {
	usertar := &biz.User{
		Username: user.Username,
	}
	result := u.data.db.First(usertar)
	if result.Error != nil {
		return false, result.Error
	}
	if usertar.Password != user.Password {
		return false, errors.New("mistake password")
	}
	return true, nil
}
func (u *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	user := &biz.User{
		Username: username,
	}
	result := u.data.db.First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
