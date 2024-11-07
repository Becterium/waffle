//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"waffle/app/waffle/interface/internal/biz"
	"waffle/app/waffle/interface/internal/conf"
	"waffle/app/waffle/interface/internal/data"
	"waffle/app/waffle/interface/internal/server"
	"waffle/app/waffle/interface/internal/service"
)

// wireApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Auth, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
