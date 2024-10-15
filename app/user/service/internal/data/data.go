package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"waffle/app/user/service/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMysqlClient, NewUserRepo, NewAddressRepo, NewRedisClient)

// Data .
type Data struct {
	db *gorm.DB
	rc *redis.Client
}

// NewData .
func NewData(mysqlClient *gorm.DB, redisClient *redis.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "user-service/data"))
	data := &Data{
		db: mysqlClient,
		rc: redisClient,
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
	if err = db.AutoMigrate(); err != nil {
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
