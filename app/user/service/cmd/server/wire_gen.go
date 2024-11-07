// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"waffle/app/user/service/internal/biz"
	"waffle/app/user/service/internal/conf"
	"waffle/app/user/service/internal/data"
	"waffle/app/user/service/internal/server"
	"waffle/app/user/service/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewMysqlClient(confData, logger)
	client := data.NewRedisClient(confData, logger)
	dataData, cleanup, err := data.NewData(db, client, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	userService := service.NewUserService(userUseCase)
	grpcServer := server.NewGRPCServer(confServer, userService, logger)
	registrar := server.NewConsulClient(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
