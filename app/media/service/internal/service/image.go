package service

import (
	"context"
	v1 "waffle/api/media/service/v1"
)

func (m *MediaService) UploadImage(ctx context.Context, req *v1.UploadImageReq) (*v1.UploadImageReply, error) {
	return m.ic.Upload(ctx, req)
}

func (m *MediaService) VerifyUploadImage(ctx context.Context, req *v1.VerifyUploadImageReq) (*v1.VerifyUploadImageReply, error) {
	return m.ic.VerifyUpload(ctx, req)
}

func (m *MediaService) GetImage(ctx context.Context, req *v1.GetImageReq) (*v1.GetImageReply, error) {
	return m.ic.Get(ctx, req)
}
