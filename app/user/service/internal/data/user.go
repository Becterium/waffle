package data

import (
	"context"
	"encoding/json"
	errors2 "errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
	"waffle/app/user/service/internal/biz"
	"waffle/app/user/service/internal/pkg/util"
)

var _ biz.UserRepo = (*userRepo)(nil)

var userCacheKey = func(username string) string {
	return "user_cache_key_" + username
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
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

	po := &User{
		Username: user.Username,
		Password: hashPassword,
	}

	result := u.data.db.Create(po)
	if result.Error != nil {
		return nil, result.Error
	}

	return &biz.User{Id: po.ID, Username: po.Username}, nil
}

func (u *userRepo) GetUser(ctx context.Context, id uint64) (*biz.User, error) {
	cacheKey := userCacheKey(fmt.Sprintf("%d", id))
	target, err := u.getUserFromCache(ctx, cacheKey)
	if err != nil {
		user := &User{
			Model: gorm.Model{
				ID: uint(id),
			},
		}
		result := u.data.db.Find(user)
		if result.RowsAffected == 0 {
			return nil, errors2.New("user not found")
		}
		if result.Error != nil {
			return nil, result.Error
		}

		u.setUserCache(ctx, user, cacheKey)
		return &biz.User{Id: user.ID, Username: user.Username}, err
	}

	return &biz.User{Id: target.Model.ID, Username: target.Username}, nil
}

func (u *userRepo) VerifyPassword(ctx context.Context, user *biz.User) error {
	findUser := &User{
		Username: user.Username,
	}

	result := u.data.db.First(&findUser)
	if result.Error != nil {
		return result.Error
	}
	if pass := util.CheckPasswordHash(findUser.Password, user.Password); pass == true {
		return nil
	} else {
		return errors.New(http.StatusUnauthorized, "REGISTER_FAILED", "wrong password")
	}

}

func (u *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	user := &User{
		Username: username,
	}
	result := u.data.db.Where("username = ?", username).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		Id:       user.ID,
		Username: user.Username,
	}, nil
}

func (u *userRepo) getUserFromCache(ctx context.Context, key string) (*User, error) {
	result, err := u.data.rc.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var cacheUser = &User{}
	err = json.Unmarshal([]byte(result), cacheUser)
	if err != nil {
		return nil, err
	}
	return cacheUser, nil
}

func (u *userRepo) setUserCache(ctx context.Context, user *User, key string) {
	marshal, err := json.Marshal(user)
	if err != nil {
		u.log.Errorf("fail to set user cache:json.Marshal(%v) error(%v)", user, err)
	}
	err = u.data.rc.Set(ctx, key, string(marshal), time.Minute*30).Err()
	if err != nil {
		u.log.Errorf("fail to set user cache:redis.Set(%v) error(%v)", user, err)
	}
}

func (u *userRepo) InitCache(ctx context.Context) (string, error) {
	users := make([]User, 0)
	result := u.data.db.Find(&users)
	// todo: handle users if users is nil
	for _, val := range users {
		key := userCacheKey(strconv.Itoa(int(val.ID)))
		u.setUserCache(ctx, &val, key)
	}
	if result.RowsAffected == 0 {
		return "", errors2.New("")
	}
	if result.Error != nil {
		return "", result.Error
	}
	return "", nil
}
