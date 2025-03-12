package service

import (
	"context"

	v1 "waffle/api/waffle/interface/v1"
)

func (w *WaffleInterface) GenerateUploadImgUrl(ctx context.Context, req *v1.GenerateUploadImgUrlReq) (*v1.GenerateUploadImgUrlReply, error) {
	return w.mc.GenerateUploadImgUrl(ctx, req)
}

func (w *WaffleInterface) VerifyImagesUpload(ctx context.Context, req *v1.VerifyImagesUploadReq) (*v1.VerifyImagesUploadReply, error) {
	return w.mc.VerifyImagesUpload(ctx, req)
}

func (w *WaffleInterface) AddImageTag(ctx context.Context, req *v1.AddImageTagReq) (*v1.AddImageTagReply, error) {
	return w.mc.AddImageTag(ctx, req)
}

func (w *WaffleInterface) SearchImageTagByNameLike(ctx context.Context, req *v1.SearchImageTagByNameLikeReq) (*v1.SearchImageTagByNameLikeReply, error) {
	return w.mc.SearchImageTagByNameLike(ctx, req)
}

func (w *WaffleInterface) ReloadCategoryRedisImageTag(ctx context.Context, req *v1.ReloadCategoryRedisImageTagReq) (*v1.ReloadCategoryRedisImageTagReply, error) {
	return w.mc.ReloadCategoryRedisImageTag(ctx, req)
}
