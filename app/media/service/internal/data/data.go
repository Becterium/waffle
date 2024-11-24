package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"waffle/app/media/service/internal/conf"
)

const (
	DefalutUseSSL = false
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMediaRepo, NewImageRepo, NewMinioClient, NewMysqlClient, NewRedisClient)

type Data struct {
	mc  *minio.Client
	log *log.Helper
	db  *gorm.DB
	rc  *redis.Client
}

func NewData(minioClient *minio.Client, logger log.Logger, mysqlClient *gorm.DB, redisClient *redis.Client) (*Data, func(), error) {

	log := log.NewHelper(log.With(logger, "module", "media-service/data"))

	d := &Data{
		mc:  minioClient,
		log: log,
		db:  mysqlClient,
		rc:  redisClient,
	}

	return d, func() {
	}, nil
}

func NewMysqlClient(c *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "user-service/data/db"))

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	if err = db.AutoMigrate(&image{}, &tag{}, &imageTag{}, &avatar{}); err != nil {
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

func NewMinioClient(conf *conf.Minio, logger log.Logger) *minio.Client {
	log := log.NewHelper(log.With(logger, "module", "user-service/data/ent"))
	endpoint := conf.Client.GetEndpoint()
	accessKeyID := conf.Client.GetKeyId()
	secretAccessKey := conf.Client.GetAccessKey()
	useSSL := DefalutUseSSL
	// Initialize minio client object.

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("minio connect fail :%v", err.Error())
	}
	return minioClient
}
