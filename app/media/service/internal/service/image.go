package service

import (
	"context"
	v1 "waffle/api/media/service/v1"
)

func (m *MediaService) UploadImages(ctx context.Context, req *v1.UploadImagesReq) (*v1.UploadImagesReply, error) {
	return m.ic.ImagesUpload(ctx, req)
}

func (m *MediaService) UploadUserImage(ctx context.Context, req *v1.UploadUserImageReq) (*v1.UploadUserImageReply, error) {
	return m.ic.UserImageUpload(ctx, req)
}

func (m *MediaService) VerifyImagesUpload(ctx context.Context, req *v1.VerifyImagesUploadReq) (*v1.VerifyImagesUploadReply, error) {
	return m.ic.VerifyImagesUpload(ctx, req)
}

func (m *MediaService) VerifyUserImageUpload(ctx context.Context, req *v1.VerifyUserImageUploadReq) (*v1.VerifyUserImageUploadReply, error) {
	return m.ic.VerifyUserImageUpload(ctx, req)
}

func (m *MediaService) GetImage(ctx context.Context, req *v1.GetImageReq) (*v1.GetImageReply, error) {
	return m.ic.Get(ctx, req)
}
