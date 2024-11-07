//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"waffle/app/user/service/internal/biz"
	"waffle/app/user/service/internal/conf"
	"waffle/app/user/service/internal/data"
	"waffle/app/user/service/internal/server"
	"waffle/app/user/service/internal/service"
)

// wireApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
