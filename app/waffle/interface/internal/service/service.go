package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"time"
	v1 "waffle/api/waffle/interface/v1"
	"waffle/app/waffle/interface/internal/biz"
	"waffle/app/waffle/interface/internal/conf"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewWaffleInterface, NewRedisClient)

type WaffleInterface struct {
	v1.UnimplementedWaffleInterfaceServer

	cash *redis.Client

	log *log.Helper
	uc  *biz.UserUseCase
	mc  *biz.MediaUseCase
}

func NewWaffleInterface(logger log.Logger, uc *biz.UserUseCase, mc *biz.MediaUseCase) *WaffleInterface {
	return &WaffleInterface{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		uc:  uc,
		mc:  mc,
	}
}

func NewRedisClient(c *conf.Data, logger log.Logger) *redis.Client {
	log := log.NewHelper(log.With(logger, "module", "waffle-interface/service/redisClient"))

	client := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     "", // 没有密码，默认值
		DB:           2,  // 默认DB 0
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}
	return client
}
