package service

import (
	"github.com/google/wire"
	v1 "waffle/api/media/service/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewMediaService)

type MediaService struct {
	v1.UnimplementedMediaServer
}

func NewMediaService() MediaService {
	return MediaService{}
}
