package data

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	v1 "waffle/api/media/service/v1"
	"waffle/app/media/service/internal/biz"
	"waffle/app/media/service/internal/data/esDSL"
	util2 "waffle/app/media/service/internal/pkg/util"
	"waffle/model/mq_kafka"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/minio/minio-go/v7"
	"github.com/segmentio/kafka-go"
	"github.com/tx7do/kratos-transport/broker"
	"gorm.io/gorm"
)

const (
	RedisCashEmpty              = "null"
	RedisHashTagNameId          = "tag"
	RedisHashTagIdLevel         = "tag:level"
	RedisHashTagCategoryTree    = "tag:category"
	RedisHashUserIdThumbnailUrl = "tag:thumbnail_url"
	ElasticsearchTagsIndex      = "tags"
)

//	img := image{
//		Model:     gorm.Model{},
//		ImageUuid: "N1IT4IN",
//		ImageName: "wallpaper-N1IT4IN.png",
//		ImageUrl:  "http://192.168.37.100:30000/images/wallpaper-N1IT4IN.png",
//		Category:  "General",
//		Purity:    "SFW",
//		Uploader:  210,
//		Size:      15,
//		Views:     9582,
//	}
type image struct {
	gorm.Model
	ImageUuid string
	ImageName string
	ImageUrl  string
	Category  string
	Purity    string
	Uploader  uint
	Size      int64
	Views     int64
}

type tag struct {
	gorm.Model
	Name     string
	ParentId uint
	Level    int
}

type TagES struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	ParentId uint   `json:"parent_id"`
	Level    int    `json:"level"`
}

type imageTag struct {
	ImageID   uint           `gorm:"primaryKey"`
	TagID     uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type avatar struct {
	gorm.Model
	UserID        uint
	AvatarName    string
	AvatarUuid    string
	AvatarUrl     string
	ThumbnailsUrl string
}

type collection struct {
	gorm.Model
	UserId uint
}

type collectionImage struct {
	CollectionId uint
	ImageId      uint
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
	// Generates a pre-signed url
	presignedURL, err := m.data.mc.PresignedPutObject(context.Background(), bucket, imageName, limitTime)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

func (m *imageRepo) SaveImagesInfo(ctx context.Context, images *biz.Images) error {
	userId, errGetUserID := util2.MetadataGetUserIdFromMetaData(ctx)
	if errGetUserID != nil {
		return errGetUserID
	}
	storeImps := make([]image, 0)
	for _, val := range *images {
		img := image{
			ImageUuid: val.ImageUuid,
			ImageName: val.ImageName,
			ImageUrl:  "http://192.168.37.100:30000/image/" + val.ImageName,
			Category:  val.Category,
			Purity:    val.Purity,
			Uploader:  userId, // todo: check 如果有设置token的话，要从token中获得userID,参考/TODO/aim.md
			Size:      0,
			Views:     0,
		}
		storeImps = append(storeImps, img)
	}

	// 批量插入Image数据
	err := m.data.db.Model(&storeImps).CreateInBatches(storeImps, len(storeImps)).Error
	if err != nil {
		return errors.New(fmt.Sprintf("db CreateInBatches error: %s", err))
	}
	for index, val := range *images {
		for _, single := range val.Tags {
			imgtag := imageTag{
				ImageID: storeImps[index].ID,
				TagID:   uint(single),
			}
			errF := m.data.db.Model(&imgtag).Create(&imgtag).Error
			if errF != nil {
				return errors.New(fmt.Sprintf("db create imgtag error: %s", errF))
			}
		}

		_, errSAdd := m.data.rc.SAdd(ctx, "image", storeImps[index].ImageUuid).Result()
		if errSAdd != nil {
			return errors.New(fmt.Sprintf("redis set add imgUuid error: %s", errSAdd))
		}

		//向异步处理服务发送消息:向 elasticsearch 存储image信息，要求根据tag可以查到image
		ImageData := mq_kafka.Image{
			ImageUuid: storeImps[index].ImageUuid,
			ImageName: storeImps[index].ImageName,
			ImageUrl:  storeImps[index].ImageName,
			Category:  storeImps[index].Category,
			Purity:    storeImps[index].Purity,
			Uploader:  userId,
			Size:      storeImps[index].Size,
			Views:     storeImps[index].Views,
			Tags:      val.Tags,
		}
		msgContent, marshalErr := json.Marshal(ImageData)
		if marshalErr != nil {
			return errors.New(fmt.Sprintf("json marshal image err :%s", marshalErr))
		}
		msg := kafka.Message{
			Key:   []byte(val.ImageUuid),
			Value: msgContent,
			Time:  time.Now(),
			Topic: "image",
		}
		errKafkaSaveMessage := kafkaSaveMessage(m.data.kw, context.Background(), msg)
		if errKafkaSaveMessage != nil {
			return errKafkaSaveMessage
		}
	}
	return nil
}

func kafkaSaveMessage(writer *kafka.Writer, ctx context.Context, msg kafka.Message) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	err := writer.WriteMessages(ctx, msg)
	if err != nil {
		return errors.New(fmt.Sprintf("写入kafka失败, err: %v", err))
	}
	return nil
}

func (m *imageRepo) SaveAvatarInfo(ctx context.Context, avatarName string, avatarUuid string) (string, error) {
	userId, errGetUserID := util2.MetadataGetUserIdFromMetaData(ctx)
	if errGetUserID != nil {
		return "", errGetUserID
	}

	// 存入MySQL
	info := avatar{
		UserID:     userId,
		AvatarName: avatarName,
		AvatarUuid: avatarUuid,
		AvatarUrl:  "http://192.168.37.100:30000/avatar/" + avatarName,
	}
	result := m.data.db.Create(&info)
	if result.Error != nil {
		return "", errors.New(fmt.Sprintf("SaveAvatarInfo fail, save to mysql error: %s", result.Error))
	}

	// 缓存avatar的UUID
	_, errSAdd := m.data.rc.SAdd(ctx, "avatar", avatarUuid).Result()
	if errSAdd != nil {
		return "", errors.New(fmt.Sprintf("SaveAvatarInfo fail, save to redis error: %s", errSAdd))
	}

	avatarData := mq_kafka.Avatar{
		Id:         info.ID,
		UserID:     info.UserID,
		AvatarName: info.AvatarName,
		AvatarUuid: info.AvatarUuid,
		AvatarUrl:  info.AvatarUrl,
	}

	msgContent, marshalErr := json.Marshal(avatarData)
	if marshalErr != nil {
		return "", errors.New(fmt.Sprintf("SaveAvatarInfo json marshal avatar err :%s", marshalErr))
	}
	msg := kafka.Message{
		Key:   []byte(avatarData.AvatarUuid),
		Value: msgContent,
		Time:  time.Now(),
		Topic: "avatar",
	}
	errKafkaSaveMessage := kafkaSaveMessage(m.data.kw, context.Background(), msg)
	if errKafkaSaveMessage != nil {
		return "", errKafkaSaveMessage
	}
	return info.AvatarUrl, nil
}

func (m *imageRepo) AddImageTag(ctx context.Context, name, parentName string) (*v1.AddImageTagReply, error) {
	// 处理父tag
	// 当不是最大Tag分类时
	parent := tag{
		Name: parentName,
	}
	if parentName != "" {
		if result, errGetName := m.data.rc.HGet(ctx, RedisHashTagNameId, parentName).Result(); errGetName != nil || result == RedisCashEmpty {
			if result == RedisCashEmpty {
				return nil, errors.New("parent Name don't exist")
			}
			res := m.data.db.Model(&tag{}).Find(&parent)
			if res.Error != nil {
				// 此处证明Redis和数据库中都没有此数据，需要短时间缓存空值
				m.data.rc.HSet(ctx, RedisHashTagNameId, parentName, RedisCashEmpty)
				m.data.rc.HExpire(ctx, RedisHashTagNameId, util2.RedisCacheNullTime(), parentName)
				return nil, res.Error
			}
			// 此处证明数据库的数据有，redis中却没有，需要存入redis
			m.data.rc.HSet(ctx, RedisHashTagNameId, parent.Name, uint64(parent.ParentId))
		} else {
			m.data.rc.HExpire(ctx, RedisHashTagIdLevel, util2.RedisImageTagCreateRandTime(), parentName)
			parseUint, errParse := strconv.ParseUint(result, 10, 64)
			if errParse != nil {
				// todo : 这里表示Redis里面存入的ID数据有问题，应该添加要及时处理的错误日志
				return nil, errors.New(fmt.Sprintf("Save Tag Error,ParseUint error : %s", errParse))
			}
			parent.ID = uint(parseUint)
			if levelInfo, errGetLevel := m.data.rc.HGet(ctx, RedisHashTagIdLevel, result).Result(); errGetLevel != nil {
				// todo : 这里表示Redis里面存入了ID数据却没有指定的Level存入，应该增进处理
				return nil, errors.New(fmt.Sprintf("Save Tag Error,can't get level cashe after get id, error : %s", errGetLevel))
			} else {
				// 此处忽略了存储别的数据的情况
				level, _ := strconv.Atoi(levelInfo)
				parent.Level = level
				m.data.rc.HExpire(ctx, RedisHashTagIdLevel, util2.RedisImageTagCreateRandTime(), result)
			}
		}
	}

	// 开始处理存储tag
	if parentName == "" {
		parent.Level = 0
	}

	t := tag{
		Name:     name,
		ParentId: parent.ID,
		Level:    parent.Level + 1,
	}

	// 添加tag入数据库
	result := m.data.db.Model(&tag{}).Create(&t)
	if result.Error != nil {
		return nil, result.Error
	}

	// 处理tag的缓存
	m.data.rc.HSet(ctx, RedisHashTagNameId, t.Name, t.ID)
	m.data.rc.HExpire(ctx, RedisHashTagNameId, util2.RedisImageTagCreateRandTime(), t.Name)
	m.data.rc.HSet(ctx, RedisHashTagIdLevel, t.ID, t.Level)
	m.data.rc.HExpire(ctx, RedisHashTagIdLevel, util2.RedisImageTagCreateRandTime(), strconv.Itoa(int(t.ID)))

	//处理tag存入Elasticsearch
	ta := TagES{
		Id:       t.ID,
		Name:     t.Name,
		ParentId: t.ParentId,
		Level:    t.Level,
	}

	if err := esDSL.PutDoc[TagES](m.data.es, ta, ElasticsearchTagsIndex, strconv.Itoa(int(ta.Id))); err != nil {
		return nil, err
	}

	return &v1.AddImageTagReply{}, nil
}

func (m *imageRepo) SearchImageTagByNameLike(ctx context.Context, name string) (*v1.SearchImageTagByNameLikeReply, error) {
	reply := make([]*v1.SearchImageTagByNameLikeReply_Tags, 0)
	fuzzy, err := esDSL.QueryFuzzy[TagES](m.data.es, ElasticsearchTagsIndex, "name", name)
	if err != nil {
		return nil, err
	}
	for _, v := range fuzzy {
		reply = append(reply, &v1.SearchImageTagByNameLikeReply_Tags{
			Name: v.Name,
			Id:   int64(v.Id),
		})
	}
	return &v1.SearchImageTagByNameLikeReply{Tags: reply}, nil
}

func (m *imageRepo) ReloadCategoryRedisImageTag(ctx context.Context, req *v1.ReloadCategoryRedisImageTagReq) (*v1.ReloadCategoryRedisImageTagReply, error) {
	FirstCategory := make([]tag, 0)
	res := m.data.db.Model(&tag{Level: 1}).Find(&FirstCategory)
	if res.Error != nil {
		return nil, res.Error
	}
	for _, f := range FirstCategory {
		SecondCategory := make([]tag, 0)
		res := m.data.db.Model(&tag{ParentId: f.ParentId}).Find(&SecondCategory)
		if res.Error != nil {
			return nil, res.Error
		}
		nodes := make([]string, 0)
		for _, s := range SecondCategory {
			nodes = append(nodes, strconv.Itoa(int(s.ID)))
			nodes = append(nodes, s.Name)
		}
		_, err := m.data.rc.HMSet(ctx, RedisHashTagCategoryTree, nodes).Result()
		if err != nil {
			return nil, err
		}
	}
	return &v1.ReloadCategoryRedisImageTagReply{}, nil
}

func (m *imageRepo) CreateCollection(ctx context.Context, userId uint) (*v1.CreateCollectionReply, error) {
	clt := &collection{
		UserId: userId,
	}
	result := m.data.db.Create(clt)
	if result.Error != nil {
		return nil, result.Error
	}
	return &v1.CreateCollectionReply{}, nil
}

func (m *imageRepo) KafkaImageSaveToElasticsearch(ctx context.Context, topic string, headers broker.Headers, msg *mq_kafka.Image) error {
	data, _ := json.Marshal(msg)

	res, err := m.data.es.Index(
		"images",
		bytes.NewReader(data),
		m.data.es.Index.WithDocumentID(msg.ImageUuid),
		m.data.es.Index.WithRefresh("true"),
	)

	if err != nil {
		return errors.New(fmt.Sprintf("media/data/KafkaSaveToElasticsearch fail to save image to ES, error: %v", err))
	}

	defer res.Body.Close()

	if res.IsError() {
		return errors.New(fmt.Sprintf("KafkaSaveToElasticsearch Error response: %s", res.String()))
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			fmt.Printf("Document indexed with ID: %s\n", r["_id"])
		}
	}
	return nil
}

func (m *imageRepo) KafkaAvatarSaveToElasticsearch(ctx context.Context, topic string, headers broker.Headers, msg *mq_kafka.Avatar) error {
	object, err := m.data.mc.GetObject(context.Background(), "avatar", msg.AvatarName, minio.GetObjectOptions{})
	if err != nil {
		return errors.New(fmt.Sprintf("KafkaAvatarSaveToElasticsearch GetObject error: %s", err))
	}
	thumbnail, err := util2.Thumbnail(object)
	if err != nil {
		return errors.New(fmt.Sprintf("KafkaAvatarSaveToElasticsearch Thumbnail error: %s", err))
	}
	// 上传缩略图到 MinIO
	thumbObjectName := "thumb_avatar-" + msg.AvatarUuid + ".png"

	_, err = m.data.mc.PutObject(context.Background(), "thumbnails", thumbObjectName, thumbnail, int64(thumbnail.Len()), minio.PutObjectOptions{
		ContentType: "image/png",
	})
	if err != nil {
		return errors.New(fmt.Sprintf("KafkaAvatarSaveToElasticsearch PutObject error: %s", err))
	}

	ava := avatar{
		Model: gorm.Model{ID: msg.Id},
	}

	ThumbnailsUrl := "http://192.168.37.100:30000/thumbnails/" + thumbObjectName
	//更新avatar的ThumbnailsUrl
	result := m.data.db.Model(&ava).Update("thumbnails_url", ThumbnailsUrl)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("KafkaAvatarSaveToElasticsearch Update thumbnails_url error: %s", result.Error))
	}

	// 更新map[userId]ThumbnailsUrl到Redis
	_, err = m.data.rc.HMSet(ctx, RedisHashUserIdThumbnailUrl, msg.UserID, ThumbnailsUrl).Result()
	if err != nil {
		return errors.New(fmt.Sprintf("KafkaAvatarSaveToElasticsearch redis HMSet thumbnails_url error: %s", err))
	}
	return nil
}
