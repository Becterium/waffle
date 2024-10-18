package biz

import "github.com/go-kratos/kratos/v2/log"

type MediaRepo interface {
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
