package server

import (
	"context"
	"waffle/app/media/service/internal/conf"
	"waffle/app/media/service/internal/service"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/kratos-transport/transport/kafka"
)

// NewKafkaServer create a kafka server.
func NewKafkaServer(c *conf.Server, _ log.Logger, svc *service.MediaService) *kafka.Server {
	ctx := context.Background()
	srv := kafka.NewServer(
		kafka.WithAddress(c.Kafka.Addrs),
		kafka.WithCodec("json"),
	)

	registerKafkaSubscribers(ctx, srv, svc)

	return srv
}

// @param queue 订阅的分组
func registerKafkaSubscribers(ctx context.Context, srv *kafka.Server, svc *service.MediaService) {
	err := srv.RegisterSubscriber(ctx,
		"image",
		"test",
		false,
		registerImageHandler(svc.HandleKafkaImageSaveToElasticsearch),
		imageCreator,
	)

	if err != nil {
		panic(err)
	}

	err = srv.RegisterSubscriber(ctx,
		"avatar",
		"test",
		false,
		registerAvatarHandler(svc.HandleKafkaAvatarSaveToElasticsearch),
		avatarCreator,
	)

	if err != nil {
		panic(err)
	}
}
