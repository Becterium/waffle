package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"time"
	"waffle/app/user/service/internal/biz"
	"waffle/app/user/service/internal/pkg/util"
)

var _ biz.UserRepo = (*userRepo)(nil)

var userCacheKey = func(username string) string {
	return "user_cache_key_" + username
}

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
	hashPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashPassword)

	result := u.data.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	cacheKey := userCacheKey(fmt.Sprintf("%d", id))
	target, err := u.getUserFromCache(ctx, cacheKey)
	if err != nil {
		//Todo handle error And get target from mysql
		return nil, err
	}
	user := &biz.User{
		Id: id,
	}
	result := u.data.db.Find(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{Id: target.Id, Username: target.Username}, nil
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
	result := u.data.db.Last(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *userRepo) getUserFromCache(ctx context.Context, key string) (*biz.User, error) {
	result, err := u.data.rc.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var cacheUser = &biz.User{}
	err = json.Unmarshal([]byte(result), cacheUser)
	if err != nil {
		return nil, err
	}
	return cacheUser, nil
}

func (u *userRepo) serUserCache(ctx context.Context, user *biz.User, key string) {
	marshal, err := json.Marshal(user)
	if err != nil {
		u.log.Errorf("fail to set user cache:json.Marshal(%v) error(%v)", user, err)
	}
	err = u.data.rc.Set(ctx, key, string(marshal), time.Minute*30).Err()
	if err != nil {
		u.log.Errorf("fail to set user cache:redis.Set(%v) error(%v)", user, err)
	}
}
