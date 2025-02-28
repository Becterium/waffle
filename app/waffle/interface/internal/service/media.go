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
