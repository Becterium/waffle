package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	v1 "waffle/api/media/service/v1"
	"waffle/app/media/service/internal/biz"
)

type Images struct {
	gorm.Model
	ImageUuid string
	Category  string
	Purity    string
	UserId    int64
}

type Tags struct {
	gorm.Model
	TagName string
}

type imageRepo struct {
	data *Data
	log  *log.Helper
}

func NewImageRepo(data *Data, logger log.Logger) biz.ImageRepo {
	return &imageRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/image")),
	}
}

func (m *imageRepo) UploadImage(ctx context.Context, images *Images) (*v1.UploadImageReply, error) {
	return &v1.UploadImageReply{
		UploadUrl: "1",
		ImageUrl:  "2",
	}, nil
}

func (m *imageRepo) VerifyUploadImage(ctx context.Context, imageUrl string) (*v1.VerifyUploadImageReply, error) {
	return &v1.VerifyUploadImageReply{
		Success: false,
		Message: "s",
	}, nil
}

func (m *imageRepo) GetImage(ctx context.Context, imageUrl string) (*v1.GetImageReply, error) {
	return &v1.GetImageReply{
		Tags:      nil,
		Uploader:  "2",
		Category:  "2",
		Purity:    "2",
		Size:      0,
		Views:     0,
		Favorites: 0,
		Link:      "2",
		Thumbnail: "2",
	}, nil
}
