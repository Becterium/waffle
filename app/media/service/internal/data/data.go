package data

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"waffle/app/media/service/internal/conf"
)

const (
	DefalutUseSSL = false
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMediaRepo, NewImageRepo, NewMinioClient, NewMysqlClient, NewRedisClient, NewKafkaWriter, NewElasticsearchClient)

type Data struct {
	mc  *minio.Client
	log *log.Helper
	db  *gorm.DB
	rc  *redis.Client
	kw  *kafka.Writer
	es  *elasticsearch.Client
}

func NewData(minioClient *minio.Client, logger log.Logger, mysqlClient *gorm.DB, redisClient *redis.Client, kafkaWriter *kafka.Writer) (*Data, func(), error) {

	log := log.NewHelper(log.With(logger, "module", "media-service/data"))

	d := &Data{
		mc:  minioClient,
		log: log,
		db:  mysqlClient,
		rc:  redisClient,
		kw:  kafkaWriter,
	}

	return d, func() {
		// todo: 优雅close
		d.kw.Close()
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
	log := log.NewHelper(log.With(logger, "module", "media-service/data/redis"))

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
	log := log.NewHelper(log.With(logger, "module", "media-service/data/minio"))
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

func NewKafkaWriter(conf *conf.Server, logger log.Logger) *kafka.Writer {
	writer := kafka.Writer{
		Addr:                   kafka.TCP(conf.Kafka.Broker.Addr),
		Topic:                  conf.Kafka.Topic,
		Balancer:               nil,
		MaxAttempts:            0,
		WriteBackoffMin:        0,
		WriteBackoffMax:        0,
		BatchSize:              0,
		BatchBytes:             0,
		BatchTimeout:           0,
		ReadTimeout:            0,
		WriteTimeout:           time.Second,
		RequiredAcks:           kafka.RequireNone,
		Async:                  false,
		Completion:             nil,
		Compression:            0,
		Logger:                 nil,
		ErrorLogger:            nil,
		Transport:              nil,
		AllowAutoTopicCreation: false,
	}
	return &writer
}

func NewElasticsearchClient(conf *conf.Data, logger log.Logger) *elasticsearch.Client {
	log := log.NewHelper(log.With(logger, "module", "media-service/data/elasticsearch"))
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{conf.Elasticsearch.Addr},
	})
	if err != nil {
		log.Fatalf("Elasticsearch connect fail :%v", err.Error())
	}
	return client
}
