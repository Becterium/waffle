// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"waffle/app/media/service/internal/biz"
	"waffle/app/media/service/internal/conf"
	"waffle/app/media/service/internal/data"
	"waffle/app/media/service/internal/server"
	"waffle/app/media/service/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func initApp(confServer *conf.Server, minio *conf.Minio, auth *conf.Auth, confData *conf.Data, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	client := data.NewMinioClient(minio, logger)
	db := data.NewMysqlClient(confData, logger)
	redisClient := data.NewRedisClient(confData, logger)
	writer := data.NewKafkaWriter(confServer, logger)
	elasticsearchClient := data.NewElasticsearchClient(confData, logger)
	dataData, cleanup, err := data.NewData(client, logger, db, redisClient, writer, elasticsearchClient)
	if err != nil {
		return nil, nil, err
	}
	mediaRepo := data.NewMediaRepo(dataData, logger)
	mediaUseCase := biz.NewMediaUseCase(mediaRepo, logger)
	imageRepo := data.NewImageRepo(dataData, logger)
	imageUseCase := biz.NewImageUseCase(imageRepo, logger)
	mediaService := service.NewMediaService(mediaUseCase, imageUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, auth, logger, mediaService)
	registrar := server.NewRegistrar(registry)
	kafkaServer := server.NewKafkaServer(confServer, logger, mediaService)
	cronWorker := server.NewCronServer(imageRepo)
	app := newApp(logger, grpcServer, registrar, kafkaServer, cronWorker)
	return app, func() {
		cleanup()
	}, nil
}
