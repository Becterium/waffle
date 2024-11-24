package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"
	userv1 "waffle/api/user/service/v1"
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
	result, err, _ := r.sg.Do(fmt.Sprintf("find_user_by_id_%d", id), func() (interface{}, error) {
		user, err := r.data.uc.GetUser(ctx, &userv1.GetUserReq{
			Id: id,
		})
		if err != nil {
			return nil, biz.ErrUserNotFound
		}
		return &biz.User{
			Id:       user.Id,
			Username: user.Username,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*biz.User), nil
}

func (r *userRepo) FindByName(ctx context.Context, username string) (*biz.User, error) {
	result, err, _ := r.sg.Do(fmt.Sprintf("find_user_by_username_%s", username), func() (interface{}, error) {
		user, err := r.data.uc.GetUserByUsername(ctx, &userv1.GetUserByUsernameReq{Username: username})
		if err != nil {
			return nil, biz.ErrUserNotFound
		}
		return &biz.User{
			Id:       user.Id,
			Username: user.Username,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*biz.User), nil
}

func (r *userRepo) Save(ctx context.Context, u *biz.User) error {
	_, err := r.data.uc.Save(ctx, &userv1.SaveUserReq{
		Username: u.Username,
		Password: u.Password,
	})
	return err
}

func (r *userRepo) VerifyPassword(ctx context.Context, u *biz.User, password string) error {
	_, err := r.data.uc.VerifyPassword(ctx, &userv1.VerifyPasswordReq{
		Username: u.Username,
		Password: password,
	})
	if err != nil {
		return err
	}
	return nil
}