package data

import (
	"github.com/go-kratos/kratos/v2/log"
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
