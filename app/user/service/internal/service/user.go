package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/metadata"
	v1 "waffle/api/user/service/v1"
	"waffle/app/user/service/internal/biz"
)

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	rv, err := s.uc.Create(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserReply{
		Id:       uint64(rv.Id),
		Username: rv.Username,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	rv, err := s.uc.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetUserReply{
		Id:       uint64(rv.Id),
		Username: rv.Username,
	}, nil
}

func (s *UserService) GetUserByUsername(ctx context.Context, req *v1.GetUserByUsernameReq) (*v1.GetUserByUsernameReply, error) {
	rv, err := s.uc.FindByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}
	return &v1.GetUserByUsernameReply{
		Id:       uint64(rv.Id),
		Username: rv.Username,
	}, nil
}

func (s *UserService) Save(ctx context.Context, req *v1.SaveUserReq) (*v1.SaveUserReply, error) {
	rv, err := s.uc.Save(ctx, &biz.User{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.SaveUserReply{
		Id: rv.Id,
	}, nil
}

func (s *UserService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordReq) (*v1.VerifyPasswordReply, error) {
	err := s.uc.VerifyPassword(ctx, &biz.User{Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, err
	}

	return &v1.VerifyPasswordReply{}, nil
}

func (s *UserService) InitCache(ctx context.Context, req *v1.InitCacheReq) (*v1.InitCacheReply, error) {
	return s.uc.InitCache(ctx)
}

func (s *UserService) Ping(ctx context.Context, req *v1.PingReq) (*v1.PingReply, error) {
	if md, ok := metadata.FromServerContext(ctx); ok {
		userId := md.Get("x-md-local-userId")
		return &v1.PingReply{Message: userId}, nil
	}
	return nil, errors.New("metadata can't get message from server context")
}
