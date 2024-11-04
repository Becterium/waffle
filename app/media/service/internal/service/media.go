package service

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "waffle/api/media/service/v1"
	"waffle/app/media/service/internal/biz"
)

type MediaService struct {
	v1.UnimplementedMediaServer

	mc  *biz.MediaUseCase
	ic  *biz.ImageUseCase
	log *log.Helper
}

func NewMediaService(mc *biz.MediaUseCase, ic *biz.ImageUseCase, logger log.Logger) *MediaService {
	return &MediaService{
		mc:  mc,
		ic:  ic,
		log: log.NewHelper(log.With(logger, "module", "service/media")),
	}
}
