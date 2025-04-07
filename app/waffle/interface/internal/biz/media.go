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
	Category  string
	Purity    string
	Size      int64
	Views     int64
	ImageUrl  string
	Tags      []string
}

type UploaderInfo struct {
	Id        uint64
	Name      string
	AvatarUrl string
}

type MediaRepo interface {
	GenerateUploadImgUrl(ctx context.Context, imgNames []string) (*v1.GenerateUploadImgUrlReply, error)
	VerifyImagesUpload(ctx context.Context, req *v1.VerifyImagesUploadReq) (*v1.VerifyImagesUploadReply, error)
	AddImageTag(ctx context.Context, req *v1.AddImageTagReq) (*v1.AddImageTagReply, error)
	SearchImageTagByNameLike(ctx context.Context, req *v1.SearchImageTagByNameLikeReq) (*v1.SearchImageTagByNameLikeReply, error)
	ReloadCategoryRedisImageTag(ctx context.Context, req *v1.ReloadCategoryRedisImageTagReq) (*v1.ReloadCategoryRedisImageTagReply, error)
	GenerateUploadAvatarUrl(ctx context.Context, req *v1.GenerateUploadAvatarUrlReq) (*v1.GenerateUploadAvatarUrlReply, error)
	VerifyAvatarUpload(ctx context.Context, req *v1.VerifyAvatarUploadReq) (*v1.VerifyAvatarUploadReply, error)
	GetImage(ctx context.Context, uid string) (*UploaderInfo, *ImageInfo, error)
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

func (m MediaUseCase) GenerateUploadAvatarUrl(ctx context.Context, req *v1.GenerateUploadAvatarUrlReq) (*v1.GenerateUploadAvatarUrlReply, error) {
	return m.repo.GenerateUploadAvatarUrl(ctx, req)
}

func (m MediaUseCase) VerifyAvatarUpload(ctx context.Context, req *v1.VerifyAvatarUploadReq) (*v1.VerifyAvatarUploadReply, error) {
	ctx, err := util.UnMarshalTokeToMetadata(ctx, util.PrefixLocalMetadata, util.MarshalUserId)
	if err != nil {
		return nil, err
	}
	return m.repo.VerifyAvatarUpload(ctx, req)
}

func (m MediaUseCase) GetImage(ctx context.Context, uid string) (*UploaderInfo, *ImageInfo, error) {
	return m.repo.GetImage(ctx, uid)
}
