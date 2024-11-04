package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "waffle/api/media/service/v1"
	"waffle/app/media/service/internal/biz"
)

var _ biz.MediaRepo = (*mediaRepo)(nil)

type mediaRepo struct {
	data *Data
	log  *log.Helper
}

func NewMediaRepo(data *Data, logger log.Logger) biz.MediaRepo {
	return &mediaRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/media")),
	}
}

func (m *mediaRepo) UploadVideo(ctx context.Context) (*v1.UploadImageReply, error) {
	return nil, nil
}

func (m *mediaRepo) GetVideo(ctx context.Context) (*v1.GetVideoReply, error) {
	return nil, nil
}
