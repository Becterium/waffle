package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"math/rand"
	"strings"
	"time"
	v1 "waffle/api/media/service/v1"
	"waffle/model/mq_kafka"
)

const (
	UuidLength              = 7
	ClusterNode             = 1
	RedisLibImageKey        = "LibImage"
	RedisAvatarKey          = "Avatar"
	MinioBucketLibImageName = "images"
	MinioBucketAvatarName   = "avatar"
	MinioTimeToPreSignedPut = time.Minute * 10
)

type Image struct {
	ImageName string
	ImageUuid string
	Category  string
	Purity    string
	Size      int64
	Tags      []uint64
}

type Images []Image

type ImageRepo interface {
	VerifyImageUpload(ctx context.Context, bucket string, imageUrl string) (bool, error)
	GetImage(ctx context.Context, imageUrl string) (*v1.GetImageReply, error)
	// ImageExist 由于存储在minio中的图片信息会同步到MySQL再同步到redis，所以直接在redis中查看进行防重复的索引
	ImageExist(ctx context.Context, redisKey string, imageUuid string) (bool, error)
	GeneratePutImageURL(ctx context.Context, bucket string, imageName string, limitTime time.Duration) (string, error)
	SaveImagesInfo(ctx context.Context, images *Images) error
	SaveAvatarInfo(ctx context.Context, avatarName string, avatarUuid string) (string, error)
	AddImageTag(ctx context.Context, name, parentName string) (*v1.AddImageTagReply, error)
	SearchImageTagByNameLike(ctx context.Context, name string) (*v1.SearchImageTagByNameLikeReply, error)
	ReloadCategoryRedisImageTag(ctx context.Context, req *v1.ReloadCategoryRedisImageTagReq) (*v1.ReloadCategoryRedisImageTagReply, error)
	GetImageByQueryKVsAndPageAndOrderByDESC(ctx context.Context, req *v1.GetImageByQueryKVsAndPageAndOrderByDESCReq) (*v1.GetImageByQueryKVsAndPageAndOrderByDESCReply, error)
	// collection
	CreateCollection(ctx context.Context, userId uint) (*v1.CreateCollectionReply, error)
	// KafkaImageSaveToElasticsearch kafka consume message
	KafkaImageSaveToElasticsearch(ctx context.Context, topic string, headers broker.Headers, msg *mq_kafka.Image) error
	KafkaAvatarSaveToElasticsearch(ctx context.Context, topic string, headers broker.Headers, msg *mq_kafka.Avatar) error
	CronSynchronizeImageViewFromRedis() func()
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

// generate uuid like this "zkgz5yy"
func generateShortUUID() (string, error) {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(int64(ClusterNode))
	if err != nil {
		return "", err
	}
	// Generate a snowflake ID.
	charset := node.Generate().Base64()
	rand.NewSource(time.Now().UnixNano())
	b := make([]byte, UuidLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset)-2)]
	}
	return string(b), nil
}

// ImagesUpload 承担预签名URL的作用
func (c *ImageUseCase) ImagesUpload(ctx context.Context, req *v1.UploadImagesReq) (*v1.UploadImagesReply, error) {
	var result v1.UploadImagesReply
	for _, name := range req.ImageName {
		exist := true
		var err error
		var uid string
		for exist {
			uid, _ = generateShortUUID()
			exist, err = c.ip.ImageExist(ctx, RedisLibImageKey, uid)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("function: ImagesUpload, check image exist from redis error: %s", err))
			}
		}
		split := strings.Split(name, ".")
		imageName := "wallpaper-" + uid + "." + split[len(split)-1]
		url, err := c.ip.GeneratePutImageURL(ctx, MinioBucketLibImageName, imageName, MinioTimeToPreSignedPut)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("function: ImagesUpload, generate pre-signed imageURL error: %s", err))
		}
		reply := v1.UploadImagesReply_Image{
			UploadUrl: url,
			ImageName: imageName,
			ImageUuid: uid,
		}
		result.Result = append(result.Result, &reply)
	}
	return &result, nil
}

func (c *ImageUseCase) UserImageUpload(ctx context.Context, req *v1.UploadUserImageReq) (*v1.UploadUserImageReply, error) {
	exist := true
	var err error
	var uid string
	for exist {
		uid, _ = generateShortUUID()
		exist, err = c.ip.ImageExist(ctx, RedisAvatarKey, uid)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("function: ImagesUpload, check avatar exist from redis error: %s", err))
		}
	}
	split := strings.Split(req.ImageName, ".")
	name := "avatar-" + uid + "." + split[len(split)-1]
	url, err := c.ip.GeneratePutImageURL(ctx, MinioBucketAvatarName, name, MinioTimeToPreSignedPut)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("function: ImagesUpload, generate pre-signed avatarURL error: %s", err))
	}
	return &v1.UploadUserImageReply{
		UploadUrl:  url,
		AvatarName: name,
		AvatarUuid: uid,
	}, nil
}

func (c *ImageUseCase) VerifyImagesUpload(ctx context.Context, req *v1.VerifyImagesUploadReq) (*v1.VerifyImagesUploadReply, error) {
	for _, image := range req.ImageInfo {
		exist, err := c.ip.VerifyImageUpload(ctx, MinioBucketLibImageName, image.ImageName)
		if err != nil {
			return &v1.VerifyImagesUploadReply{
				Success: false,
				Message: err.Error(),
			}, err
		}
		if !exist {
			return &v1.VerifyImagesUploadReply{
				Success: false,
				Message: fmt.Sprintf("LibImage did`t exist, image name : %s, image uuid: %s", image.ImageName, image.ImageUuid),
			}, err
		}
	}
	var images Images
	for _, info := range req.ImageInfo {
		images = append(images, Image{
			ImageName: info.ImageName,
			ImageUuid: info.ImageUuid,
			Category:  info.Category,
			Purity:    info.Purity,
			Tags:      info.Tags,
		})
	}
	err := c.ip.SaveImagesInfo(ctx, &images)
	if err != nil {
		return nil, err
	}
	return &v1.VerifyImagesUploadReply{
		Success: true,
		Message: "LibImages Upload success",
	}, nil
}

func (c *ImageUseCase) VerifyUserImageUpload(ctx context.Context, req *v1.VerifyUserImageUploadReq) (*v1.VerifyUserImageUploadReply, error) {
	exist, err := c.ip.VerifyImageUpload(ctx, MinioBucketAvatarName, req.AvatarName)
	if err != nil {
		return &v1.VerifyUserImageUploadReply{
			Success: false,
			Message: err.Error(),
		}, err
	}
	if !exist {
		return &v1.VerifyUserImageUploadReply{
			Success: false,
			Message: "Avatar did`t exist",
		}, err
	}
	url, err := c.ip.SaveAvatarInfo(ctx, req.AvatarName, req.AvatarUuid)
	if err != nil {
		return &v1.VerifyUserImageUploadReply{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return &v1.VerifyUserImageUploadReply{
		Success:   true,
		Message:   "success upload user Avatar",
		AvatarUrl: url,
	}, nil
}

func (c *ImageUseCase) Get(ctx context.Context, req *v1.GetImageReq) (*v1.GetImageReply, error) {
	return c.ip.GetImage(ctx, req.ImageUrl)
}

func (c *ImageUseCase) AddImageTag(ctx context.Context, in *v1.AddImageTagReq) (*v1.AddImageTagReply, error) {
	return c.ip.AddImageTag(ctx, in.GetName(), in.GetParentName())
}

func (c *ImageUseCase) SearchImageTagByNameLike(ctx context.Context, in *v1.SearchImageTagByNameLikeReq) (*v1.SearchImageTagByNameLikeReply, error) {
	return c.ip.SearchImageTagByNameLike(ctx, in.GetName())
}

func (c *ImageUseCase) ReloadCategoryRedisImageTag(ctx context.Context, req *v1.ReloadCategoryRedisImageTagReq) (*v1.ReloadCategoryRedisImageTagReply, error) {
	return c.ip.ReloadCategoryRedisImageTag(ctx, req)
}

func (c *ImageUseCase) CreateCollection(ctx context.Context, req *v1.CreateCollectionReq) (*v1.CreateCollectionReply, error) {
	return c.ip.CreateCollection(ctx, uint(req.UserId))
}

func (c *ImageUseCase) StarImage(ctx context.Context, req *v1.StarImageReq) (*v1.StarImageReply, error) {
	return nil, nil
}
func (c *ImageUseCase) UnStarImage(ctx context.Context, req *v1.UnStarImageReq) (*v1.UnStarImageReply, error) {
	return nil, nil
}
func (c *ImageUseCase) FindCollectionByImageId(ctx context.Context, req *v1.FindCollectionByImageIdReq) (*v1.FindCollectionByImageIdReply, error) {
	return nil, nil
}
func (c *ImageUseCase) FindCollectionByCollectionId(ctx context.Context, req *v1.FindCollectionByCollectionIdReq) (*v1.FindCollectionByCollectionIdReply, error) {
	return nil, nil
}

// kafka consume message

func (c *ImageUseCase) HandleKafkaImageSaveToElasticsearch(ctx context.Context, topic string, headers broker.Headers, msg *mq_kafka.Image) error {
	return c.ip.KafkaImageSaveToElasticsearch(ctx, topic, headers, msg)
}

func (c *ImageUseCase) HandleKafkaAvatarSaveToElasticsearch(ctx context.Context, topic string, headers broker.Headers, msg *mq_kafka.Avatar) error {
	return c.ip.KafkaAvatarSaveToElasticsearch(ctx, topic, headers, msg)
}

func (c *ImageUseCase) CronSynchronizeImageViewFromRedis() func() {
	return c.ip.CronSynchronizeImageViewFromRedis()
}

func (c *ImageUseCase) GetImageByQueryKVsAndPageAndOrderByDESC(ctx context.Context, req *v1.GetImageByQueryKVsAndPageAndOrderByDESCReq) (*v1.GetImageByQueryKVsAndPageAndOrderByDESCReply, error) {
	return c.ip.GetImageByQueryKVsAndPageAndOrderByDESC(ctx, req)
}
