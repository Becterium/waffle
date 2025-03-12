package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"waffle/app/waffle/interface/internal/pkg/util"

	v1 "waffle/api/waffle/interface/v1"
)

type ImageInfo struct {
	ImageName string
	ImageUUID string
	category  string
	purity    string
	tags      []uint64
}

type MediaRepo interface {
	GenerateUploadImgUrl(ctx context.Context, imgNames []string) (*v1.GenerateUploadImgUrlReply, error)
	VerifyImagesUpload(ctx context.Context, req *v1.VerifyImagesUploadReq) (*v1.VerifyImagesUploadReply, error)
	AddImageTag(ctx context.Context, req *v1.AddImageTagReq) (*v1.AddImageTagReply, error)
	SearchImageTagByNameLike(ctx context.Context, req *v1.SearchImageTagByNameLikeReq) (*v1.SearchImageTagByNameLikeReply, error)
	ReloadCategoryRedisImageTag(ctx context.Context, req *v1.ReloadCategoryRedisImageTagReq) (*v1.ReloadCategoryRedisImageTagReply, error)
}

type MediaUseCase struct {
	repo MediaRepo
	log  *log.Helper
}

func NewMediaUseCase(repo MediaRepo, logger log.Logger) *MediaUseCase {
	return &MediaUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/waffle")),
	}
}

func (m MediaUseCase) GenerateUploadImgUrl(ctx context.Context, req *v1.GenerateUploadImgUrlReq) (*v1.GenerateUploadImgUrlReply, error) {

	//将存储在token中的user_id转存进metadata中
	ctx, err := util.UnMarshalTokeToMetadata(ctx, util.PrefixLocalMetadata, util.MarshalUserId)
	if err != nil {
		return nil, err
	}

	return m.repo.GenerateUploadImgUrl(ctx, req.ImageName)
}

func (m MediaUseCase) VerifyImagesUpload(ctx context.Context, req *v1.VerifyImagesUploadReq) (*v1.VerifyImagesUploadReply, error) {

	//将存储在token中的user_id转存进metadata中
	ctx, err := util.UnMarshalTokeToMetadata(ctx, util.PrefixLocalMetadata, util.MarshalUserId)
	if err != nil {
		return nil, err
	}

	return m.repo.VerifyImagesUpload(ctx, req)
}

func (m MediaUseCase) AddImageTag(ctx context.Context, req *v1.AddImageTagReq) (*v1.AddImageTagReply, error) {
	return m.repo.AddImageTag(ctx, req)
}

func (m MediaUseCase) SearchImageTagByNameLike(ctx context.Context, req *v1.SearchImageTagByNameLikeReq) (*v1.SearchImageTagByNameLikeReply, error) {
	return m.repo.SearchImageTagByNameLike(ctx, req)
}

func (m MediaUseCase) ReloadCategoryRedisImageTag(ctx context.Context, req *v1.ReloadCategoryRedisImageTagReq) (*v1.ReloadCategoryRedisImageTagReply, error) {
	return m.repo.ReloadCategoryRedisImageTag(ctx, req)
}
