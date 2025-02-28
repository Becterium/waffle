package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	v1 "waffle/api/waffle/interface/v1"
)

func (w *WaffleInterface) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	return w.uc.Register(ctx, req)
}

func (w *WaffleInterface) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	return w.uc.Login(ctx, req)
}

func (w *WaffleInterface) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	return w.uc.Logout(ctx, req)
}

func (w *WaffleInterface) Ping(ctx context.Context, req *v1.PingReq) (*v1.PingReply, error) {
	if tokenClaims, ok := jwt.FromContext(ctx); ok {
		if claims, ook := tokenClaims.(jwtv5.MapClaims); ook {
			return &v1.PingReply{Message: claims["userId"].(string)}, nil
		}
		return nil, errors.New("jwtv5.MapClaims reflect tokenClaims failed")
	}
	return &v1.PingReply{Message: "token Unmarshell fail"}, nil
}

func (w *WaffleInterface) PingRPC(ctx context.Context, req *v1.PingRPCReq) (*v1.PingRPCReply, error) {
	return w.uc.Ping(ctx)
}
