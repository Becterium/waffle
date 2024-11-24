package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"time"
	v1 "waffle/api/media/service/v1"
	"waffle/app/media/service/internal/biz"
)

type image struct {
	gorm.Model
	ImageUuid string
	ImageName string
	ImageUrl  string
	Category  string
	Purity    string
	Uploader  int64
	Size      int64
	Views     int64
}

type tag struct {
	gorm.Model
	TagName string
}

type imageTag struct {
	ImageID   uint           `gorm:"primaryKey"`
	TagID     uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type avatar struct {
	gorm.Model
	UserID     uint
	AvatarName string
	AvatarUuid string
	AvatarUrl  string
}

type imageRepo struct {
	data *Data
	log  *log.Helper
}

func NewImageRepo(data *Data, logger log.Logger) biz.ImageRepo {
	return &imageRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/image")),
	}
}

//func (m *imageRepo) UploadImage(ctx context.Context, images *biz.Images) (*v1.UploadImageReply, error) {
//	//todo: 在user实现jwt发布令牌
//	//token, ok := jwt.FromContext(ctx)
//	//if !ok {
//	//	return nil, errors.New("jwt.Parse fail, can`t get auth info")
//	//}
//	//subject, _ := token.GetSubject()
//	//return nil, errors.New(subject)
//	return nil, nil
//}

func (m *imageRepo) VerifyImageUpload(ctx context.Context, bucket string, imageUrl string) (bool, error) {
	_, err := m.data.mc.StatObject(ctx, bucket, imageUrl, minio.StatObjectOptions{})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *imageRepo) GetImage(ctx context.Context, imageUrl string) (*v1.GetImageReply, error) {
	return &v1.GetImageReply{
		Tags:      nil,
		Uploader:  "2",
		Category:  "2",
		Purity:    "2",
		Size:      0,
		Views:     0,
		Favorites: 0,
		Link:      "2",
		Thumbnail: "2",
	}, nil
}

func (m *imageRepo) ImageExist(ctx context.Context, redisKey string, imageUuid string) (bool, error) {
	result, err := m.data.rc.SIsMember(ctx, redisKey, imageUuid).Result()
	//TODO: handle error
	return result, err
}

func (m *imageRepo) GeneratePutImageURL(ctx context.Context, bucket string, imageName string, limitTime time.Duration) (string, error) {
	// Generates a presigned url
	presignedURL, err := m.data.mc.PresignedPutObject(context.Background(), bucket, imageName, limitTime)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

func (m *imageRepo) SaveImagesInfo(ctx context.Context, images *biz.Images) error {
	storgeImgs := make([]image, 0)
	for index, val := range *images {
		img := image{
			ImageUuid: val.ImageUuid,
			ImageName: val.ImageName,
			ImageUrl:  "http://192.168.37.100:30001/image/" + val.ImageName,
			Category:  val.Category,
			Purity:    val.Purity,
			Uploader:  int64(index), // todo: 如果有设置token的话，要从token中获得userID,参考/TODO/aim.md
			Size:      val.Size,
			Views:     0,
		}
		storgeImgs = append(storgeImgs, img)
	}
	err := m.data.db.Model(&storgeImgs).CreateInBatches(storgeImgs, len(storgeImgs)).Error
	if err != nil {
		return err
	}
	for index, val := range *images {
		for _, single := range val.Tags {
			imgtag := imageTag{
				ImageID: storgeImgs[index].ID,
				TagID:   uint(single),
			}
			errF := m.data.db.Model(&imgtag).Create(&imgtag).Error
			if errF != nil {
				return errF
			}
		}
	}
	return nil
}

func (m *imageRepo) SaveAvatarInfo(ctx context.Context, avatarName string) error {
	return nil
}
