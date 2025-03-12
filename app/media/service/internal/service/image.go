package service

import (
	"context"
	"github.com/tx7do/kratos-transport/broker"
	v1 "waffle/api/media/service/v1"
	"waffle/model/mq_kafka"
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

func (m *MediaService) AddImageTag(ctx context.Context, in *v1.AddImageTagReq) (*v1.AddImageTagReply, error) {
	return m.ic.AddImageTag(ctx, in)
}

func (m *MediaService) SearchImageTagByNameLike(ctx context.Context, in *v1.SearchImageTagByNameLikeReq) (*v1.SearchImageTagByNameLikeReply, error) {
	return m.ic.SearchImageTagByNameLike(ctx, in)
}

func (m *MediaService) ReloadCategoryRedisImageTag(ctx context.Context, req *v1.ReloadCategoryRedisImageTagReq) (*v1.ReloadCategoryRedisImageTagReply, error) {
	return m.ic.ReloadCategoryRedisImageTag(ctx, req)
}

func (m *MediaService) HandleKafkaImageSaveToElasticsearch(ctx context.Context, topic string, headers broker.Headers, msg *mq_kafka.Image) error {
	return m.ic.HandleKafkaImageSaveToElasticsearch(ctx, topic, headers, msg)
}
