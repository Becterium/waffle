package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	v1 "waffle/api/user/service/v1"
	"waffle/app/user/service/internal/biz"
)

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	hashPassword, err2 := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err2 != nil {
		return nil, err2
	}

	rv, err := s.uc.Create(ctx, &biz.User{
		Username: req.Username,
		Password: string(hashPassword),
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserReply{
		Id:       rv.Id,
		Username: rv.Username,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	rv, err := s.uc.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetUserReply{
		Id:       rv.Id,
		Username: rv.Username,
	}, nil
}

func (s *UserService) GetUserByUsername(ctx context.Context, req *v1.GetUserByUsernameReq) (*v1.GetUserByUsernameReply, error) {
	rv, err := s.uc.FindByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}
	return &v1.GetUserByUsernameReply{
		Id:       rv.Id,
		Username: rv.Username,
	}, nil
}

func (s *UserService) Save(ctx context.Context, req *v1.SaveUserReq) (*v1.SaveUserReply, error) {
	rv, err := s.uc.Save(ctx, &biz.User{
		Id:       req.GetId(),
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
	rv, err := s.uc.VerifyPassword(ctx, &biz.User{Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, err
	}

	return &v1.VerifyPasswordReply{
		Ok: rv,
	}, nil
}
