package biz

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
	"math/rand"
	"time"
	v1 "waffle/api/media/service/v1"
	"waffle/app/media/service/internal/data"
)

const (
	UuidLength  = 7
	ClusterNode = 1
)

type ImageRepo interface {
	UploadImage(ctx context.Context, images *data.Images) (*v1.UploadImageReply, error)
	VerifyUploadImage(ctx context.Context, imageUrl string) (*v1.VerifyUploadImageReply, error)
	GetImage(ctx context.Context, imageUrl string) (*v1.GetImageReply, error)
}

type ImageUseCase struct {
	ip  ImageRepo
	log *log.Helper
}

func NewImageUseCase(repo ImageRepo, logger log.Logger) *ImageUseCase {
	return &ImageUseCase{
		ip:  repo,
		log: log.NewHelper(log.With(logger, "module", "usecase/image")),
	}
}

func generateShortUUID() (string, error) {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(int64(ClusterNode))
	if err != nil {
		return "", err
	}
	// Generate a snowflake ID.
	charset := node.Generate().String()
	rand.NewSource(time.Now().UnixNano())
	b := make([]byte, UuidLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b), nil
}

func (c *ImageUseCase) Upload(ctx context.Context, req *v1.UploadImageReq) (*v1.UploadImageReply, error) {
	imageUrl, err := generateShortUUID()
	if err != nil {
		return nil, err
	}
	return c.ip.UploadImage(ctx, &data.Images{
		ImageUuid: imageUrl,
		Category:  req.Category,
		Purity:    req.Purity,
		UserId:    1,
	})
}

func (c *ImageUseCase) VerifyUpload(ctx context.Context, req *v1.VerifyUploadImageReq) (*v1.VerifyUploadImageReply, error) {
	return c.ip.VerifyUploadImage(ctx, req.ImageUrl)
}

func (c *ImageUseCase) Get(ctx context.Context, req *v1.GetImageReq) (*v1.GetImageReply, error) {
	return c.ip.GetImage(ctx, req.ImageUrl)
}
