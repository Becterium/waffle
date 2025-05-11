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

func (m *MediaService) CreateCollection(ctx context.Context, req *v1.CreateCollectionReq) (*v1.CreateCollectionReply, error) {
	return m.ic.CreateCollection(ctx, req)
}

func (m *MediaService) StarImage(ctx context.Context, req *v1.StarImageReq) (*v1.StarImageReply, error) {
	return m.ic.StarImage(ctx, req)
}
func (m *MediaService) UnStarImage(ctx context.Context, req *v1.UnStarImageReq) (*v1.UnStarImageReply, error) {
	return m.ic.UnStarImage(ctx, req)
}
func (m *MediaService) FindCollectionByImageId(ctx context.Context, req *v1.FindCollectionByImageIdReq) (*v1.FindCollectionByImageIdReply, error) {
	return m.ic.FindCollectionByImageId(ctx, req)
}
func (m *MediaService) FindCollectionByCollectionId(ctx context.Context, req *v1.FindCollectionByCollectionIdReq) (*v1.FindCollectionByCollectionIdReply, error) {
	return m.ic.FindCollectionByCollectionId(ctx, req)
}

func (m *MediaService) GetImageByQueryKVsAndPageAndOrderByDESC(ctx context.Context, req *v1.GetImageByQueryKVsAndPageAndOrderByDESCReq) (*v1.GetImageByQueryKVsAndPageAndOrderByDESCReply, error) {
	return m.ic.GetImageByQueryKVsAndPageAndOrderByDESC(ctx, req)
}

func (m *MediaService) HandleKafkaImageSaveToElasticsearch(ctx context.Context, topic string, headers broker.Headers, msg *mq_kafka.Image) error {
	return m.ic.HandleKafkaImageSaveToElasticsearch(ctx, topic, headers, msg)
}

func (m *MediaService) HandleKafkaAvatarSaveToElasticsearch(ctx context.Context, topic string, headers broker.Headers, msg *mq_kafka.Avatar) error {
	return m.ic.HandleKafkaAvatarSaveToElasticsearch(ctx, topic, headers, msg)
}
