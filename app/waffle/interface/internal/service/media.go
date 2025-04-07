package service

import (
	"context"
	v1 "waffle/api/waffle/interface/v1"
	"waffle/app/waffle/interface/internal/pkg/util"
)

func (w *WaffleInterface) GenerateUploadImgUrl(ctx context.Context, req *v1.GenerateUploadImgUrlReq) (*v1.GenerateUploadImgUrlReply, error) {
	key, err := util.UnMarshalIdempotentKey(ctx)
	if err != nil {
		return nil, err
	}
	_, err = util.VerifyIdempotencyByRedis(ctx, w.cash, key)
	if err != nil {
		return nil, err
	}
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

func (w *WaffleInterface) GenerateUploadAvatarUrl(ctx context.Context, req *v1.GenerateUploadAvatarUrlReq) (*v1.GenerateUploadAvatarUrlReply, error) {
	return w.mc.GenerateUploadAvatarUrl(ctx, req)
}

func (w *WaffleInterface) VerifyAvatarUpload(ctx context.Context, req *v1.VerifyAvatarUploadReq) (*v1.VerifyAvatarUploadReply, error) {
	return w.mc.VerifyAvatarUpload(ctx, req)
}

func (w *WaffleInterface) GetImage(ctx context.Context, req *v1.GetImageReq) (*v1.GetImageReply, error) {
	upInfo, imageInfo, err := w.mc.GetImage(ctx, req.GetUid())
	if err != nil {
		return nil, err
	}

	return &v1.GetImageReply{
		Uploader: &v1.GetImageReply_Uploader{
			Id:        upInfo.Id,
			AvatarUrl: upInfo.AvatarUrl,
		},
		Info: &v1.GetImageReply_Info{
			Category: imageInfo.Category,
			Purity:   imageInfo.Purity,
			Size:     imageInfo.Size,
			Views:    imageInfo.Views,
			Url:      imageInfo.ImageUrl,
			Uid:      imageInfo.ImageUUID,
			Tags:     imageInfo.Tags,
		},
	}, nil
}
