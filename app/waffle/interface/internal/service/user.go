package service

import (
	"context"
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
