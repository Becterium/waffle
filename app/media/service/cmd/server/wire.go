//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"waffle/app/media/service/internal/biz"
	"waffle/app/media/service/internal/conf"
	"waffle/app/media/service/internal/data"
	"waffle/app/media/service/internal/server"
	"waffle/app/media/service/internal/service"
)

// wireApp init kratos application.
func initApp(*conf.Server, *conf.Minio, *conf.Auth, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
