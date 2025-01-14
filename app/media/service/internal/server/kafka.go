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
	addrs := make([]string, 0)
	srv := kafka.NewServer(
		kafka.WithAddress(append(addrs, c.Kafka.Broker.Addr)),
		kafka.WithCodec("json"),
	)

	registerKafkaSubscribers(ctx, srv, svc)

	return srv
}

// @param queue 订阅的分组
func registerKafkaSubscribers(ctx context.Context, srv *kafka.Server, svc *service.MediaService) {
	_ = srv.RegisterSubscriber(ctx,
		"image",
		"test",
		false,
		registerSensorHandler(svc.HandleKafkaImageSaveToElasticsearch),
		imageCreator,
	)
}
