package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"waffle/app/media/service/internal/conf"
)

const (
	DefalutUseSSL = false
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMediaRepo, NewImageRepo, NewMinioClient)

type Data struct {
	mc  *minio.Client
	log *log.Helper
}

func NewData(minioClient *minio.Client, logger log.Logger) (*Data, func(), error) {

	log := log.NewHelper(log.With(logger, "module", "media-service/data"))

	d := &Data{
		mc:  minioClient,
		log: log,
	}

	return d, func() {
	}, nil
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
		log.Fatalf("minio connect fail :%v", err)
	}
	return minioClient
}
