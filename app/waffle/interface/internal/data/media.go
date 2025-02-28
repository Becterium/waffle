package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"
	v1Media "waffle/api/media/service/v1"
	v1 "waffle/api/waffle/interface/v1"
	"waffle/app/waffle/interface/internal/biz"
)

type mediaRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func NewMediaRepo(data *Data, logger log.Logger) biz.MediaRepo {
	return &mediaRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/media")),
		sg:   &singleflight.Group{},
	}
}

func (r *mediaRepo) GenerateUploadImgUrl(ctx context.Context, imgNames []string) (*v1.GenerateUploadImgUrlReply, error) {
	images, err := r.data.mc.UploadImages(ctx, &v1Media.UploadImagesReq{ImageName: imgNames})
	if err != nil {
		return nil, err
	}
	result := make([]*v1.GenerateUploadImgUrlReply_Image, 0)

	for _, value := range images.Result {
		img := v1.GenerateUploadImgUrlReply_Image{
			UploadUrl: value.UploadUrl,
			ImageName: value.ImageName,
			ImageUuid: value.ImageUuid,
		}

		result = append(result, &img)
	}
	return &v1.GenerateUploadImgUrlReply{Result: result}, nil
}

func (r *mediaRepo) VerifyImagesUpload(ctx context.Context, req *v1.VerifyImagesUploadReq) (*v1.VerifyImagesUploadReply, error) {

	infos := make([]*v1Media.VerifyImagesUploadReq_Info, 0)
	for _, value := range req.ImageInfo {
		info := v1Media.VerifyImagesUploadReq_Info{
			ImageName: value.ImageName,
			ImageUuid: value.ImageUuid,
			Category:  value.Category,
			Purity:    value.Purity,
			Tags:      value.Tags,
		}
		infos = append(infos, &info)
	}

	result, err := r.data.mc.VerifyImagesUpload(ctx, &v1Media.VerifyImagesUploadReq{ImageInfo: infos})

	if err != nil {
		return nil, err
	}

	return &v1.VerifyImagesUploadReply{Message: result.Message}, nil
}
