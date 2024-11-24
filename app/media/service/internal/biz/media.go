package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "waffle/api/media/service/v1"
)

type MediaRepo interface {
	UploadVideo(ctx context.Context) (*v1.UpLoadVideoReply, error)
	GetVideo(ctx context.Context) (*v1.GetVideoReply, error)
}

type MediaUseCase struct {
	repo MediaRepo
	log  *log.Helper
}

func NewMediaUseCase(repo MediaRepo, logger log.Logger) *MediaUseCase {
	return &MediaUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/media")),
	}
}
