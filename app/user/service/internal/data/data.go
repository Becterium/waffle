package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/redis/go-redis/v9"
	grpcx "google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	mediav1 "waffle/api/media/service/v1"
	"waffle/app/user/service/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMysqlClient, NewUserRepo, NewAddressRepo, NewRedisClient, NewDiscovery, NewMediaServiceClient)

// Data .
type Data struct {
	db *gorm.DB
	rc *redis.Client
	mc mediav1.MediaClient
}

// NewData .
func NewData(mysqlClient *gorm.DB, redisClient *redis.Client, logger log.Logger, mediaClient mediav1.MediaClient) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "user-service/data"))
	data := &Data{
		db: mysqlClient,
		rc: redisClient,
		mc: mediaClient,
	}
	db, _ := data.db.DB()
	return data, func() {
		if err := db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

func NewMysqlClient(c *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "user-service/data/db"))

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	if err = db.AutoMigrate(&User{}); err != nil {
		log.Fatal("failed to Database AutoMigrate")
	}
	return db
}

func NewRedisClient(c *conf.Data, logger log.Logger) *redis.Client {
	log := log.NewHelper(log.With(logger, "module", "user-service/data/rc"))

	client := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     "", // 没有密码，默认值
		DB:           0,  // 默认DB 0
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

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewMediaServiceClient(r registry.Discovery) mediav1.MediaClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///waffle.media.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			metadata.Client(),
		),
		grpc.WithTimeout(3*time.Second),
		// 设置空闲连接超时时间
		grpc.WithOptions(grpcx.WithIdleTimeout(0)),
	)
	if err != nil {
		panic(err)
	}
	c := mediav1.NewMediaClient(conn)
	return c
}
