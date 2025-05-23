package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"
	v1Media "waffle/api/media/service/v1"
	v1 "waffle/api/waffle/interface/v1"
	"waffle/app/waffle/interface/internal/biz"
)

type mediaRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func NewMediaRepo(data *Data, logger log.Logger) biz.MediaRepo {
	return &mediaRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/media")),
		sg:   &singleflight.Group{},
	}
}

// image

func (r *mediaRepo) GenerateUploadImgUrl(ctx context.Context, imgNames []string) (*v1.GenerateUploadImgUrlReply, error) {
	images, err := r.data.mc.UploadImages(ctx, &v1Media.UploadImagesReq{ImageName: imgNames})
	if err != nil {
		return nil, err
	}
	result := make([]*v1.GenerateUploadImgUrlReply_Image, 0)

	for _, value := range images.Result {
		img := v1.GenerateUploadImgUrlReply_Image{
			UploadUrl: value.UploadUrl,
			ImageName: value.ImageName,
			ImageUuid: value.ImageUuid,
		}

		result = append(result, &img)
	}
	return &v1.GenerateUploadImgUrlReply{Result: result}, nil
}

func (r *mediaRepo) VerifyImagesUpload(ctx context.Context, req *v1.VerifyImagesUploadReq) (*v1.VerifyImagesUploadReply, error) {
	infos := make([]*v1Media.VerifyImagesUploadReq_Info, 0)
	for _, value := range req.ImageInfo {
		info := v1Media.VerifyImagesUploadReq_Info{
			ImageName: value.ImageName,
			ImageUuid: value.ImageUuid,
			Category:  value.Category,
			Purity:    value.Purity,
			Tags:      value.Tags,
		}
		infos = append(infos, &info)
	}

	result, err := r.data.mc.VerifyImagesUpload(ctx, &v1Media.VerifyImagesUploadReq{ImageInfo: infos})

	if err != nil {
		return nil, err
	}

	return &v1.VerifyImagesUploadReply{Message: result.Message}, nil
}

// image - tag

func (r *mediaRepo) AddImageTag(ctx context.Context, req *v1.AddImageTagReq) (*v1.AddImageTagReply, error) {
	_, err := r.data.mc.AddImageTag(ctx, &v1Media.AddImageTagReq{
		Name:       req.Name,
		ParentName: req.ParentName,
	})
	if err != nil {
		return nil, err
	}
	return &v1.AddImageTagReply{}, nil
}

func (r *mediaRepo) SearchImageTagByNameLike(ctx context.Context, req *v1.SearchImageTagByNameLikeReq) (*v1.SearchImageTagByNameLikeReply, error) {
	reply, err := r.data.mc.SearchImageTagByNameLike(ctx, &v1Media.SearchImageTagByNameLikeReq{Name: req.Name})
	if err != nil {
		return nil, err
	}
	result := make([]*v1.SearchImageTagByNameLikeReply_Tags, 0)
	for _, val := range reply.Tags {
		tag := v1.SearchImageTagByNameLikeReply_Tags{
			Name: val.Name,
			Id:   val.Id,
		}
		result = append(result, &tag)
	}
	return &v1.SearchImageTagByNameLikeReply{Tags: result}, nil
}

func (r *mediaRepo) ReloadCategoryRedisImageTag(ctx context.Context, req *v1.ReloadCategoryRedisImageTagReq) (*v1.ReloadCategoryRedisImageTagReply, error) {
	_, err := r.data.mc.ReloadCategoryRedisImageTag(ctx, &v1Media.ReloadCategoryRedisImageTagReq{})
	if err != nil {
		return nil, err
	}
	return &v1.ReloadCategoryRedisImageTagReply{}, nil
}

func (r *mediaRepo) GenerateUploadAvatarUrl(ctx context.Context, req *v1.GenerateUploadAvatarUrlReq) (*v1.GenerateUploadAvatarUrlReply, error) {
	reply, err := r.data.mc.UploadUserImage(ctx, &v1Media.UploadUserImageReq{ImageName: req.ImageName})
	if err != nil {
		return nil, err
	}
	return &v1.GenerateUploadAvatarUrlReply{
		UploadUrl:  reply.UploadUrl,
		AvatarName: reply.AvatarName,
		AvatarUuid: reply.AvatarUuid,
	}, nil
}

func (r *mediaRepo) VerifyAvatarUpload(ctx context.Context, req *v1.VerifyAvatarUploadReq) (*v1.VerifyAvatarUploadReply, error) {
	reply, err := r.data.mc.VerifyUserImageUpload(ctx, &v1Media.VerifyUserImageUploadReq{
		AvatarName: req.AvatarName,
		AvatarUuid: req.AvatarUuid,
	})
	if err != nil {
		return nil, err
	}
	return &v1.VerifyAvatarUploadReply{
		UploadUrl: reply.AvatarUrl,
	}, nil
}

func (r *mediaRepo) GetImage(ctx context.Context, uid string) (*biz.UploaderInfo, *biz.ImageInfo, error) {
	res, err, _ := r.sg.Do(fmt.Sprintf("get_image_by_uid_%s", uid), func() (interface{}, error) {
		info, err := r.data.mc.GetImage(ctx, &v1Media.GetImageReq{ImageUrl: uid})
		if err != nil {
			return nil, err
		}
		tags := make([]biz.Tags, 0)
		for _, val := range info.Tags {
			tag := biz.Tags{
				Id:   val.TagId,
				Name: val.TagName,
			}
			tags = append(tags, tag)
		}
		return struct {
			uploader *biz.UploaderInfo
			image    *biz.ImageInfo
		}{
			uploader: &biz.UploaderInfo{
				Id:        info.UploaderId,
				AvatarUrl: info.UploaderUrl,
			},
			image: &biz.ImageInfo{
				ImageName: info.ImageName,
				ImageUUID: info.Thumbnail,
				Category:  info.Category,
				Purity:    info.Purity,
				Size:      info.Size,
				Views:     info.Views,
				ImageUrl:  info.Link,
				Tags:      tags,
				Id:        info.ImageId,
			},
		}, nil
	})

	if err != nil {
		return nil, nil, err
	}

	result := res.(struct {
		uploader *biz.UploaderInfo
		image    *biz.ImageInfo
	})
	return result.uploader, result.image, nil
}

func (r *mediaRepo) GetImageByQueryKVsAndPageAndOrderByDESC(ctx context.Context, req *v1.GetImageByQueryKVsAndPageAndOrderByDESCReq) (*v1.GetImageByQueryKVsAndPageAndOrderByDESCReply, error) {
	res, err, _ := r.sg.Do(fmt.Sprintf("get_sort_image_by_kv_%s", req.Query_KVs), func() (interface{}, error) {
		info, err := r.data.mc.GetImageByQueryKVsAndPageAndOrderByDESC(ctx, &v1Media.GetImageByQueryKVsAndPageAndOrderByDESCReq{
			Page:      req.Page,
			Size:      req.Size,
			Query_KVs: req.Query_KVs,
		})
		if err != nil {
			return nil, err
		}
		images := make([]*v1.GetImageByQueryKVsAndPageAndOrderByDESCReply_Images, 0)
		for _, val := range info.GetImages() {
			images = append(images, &v1.GetImageByQueryKVsAndPageAndOrderByDESCReply_Images{
				ImageId: val.GetImageId(),
				Link:    val.GetLink(),
			})
		}
		return &v1.GetImageByQueryKVsAndPageAndOrderByDESCReply{Images: images}, nil
	})

	if err != nil {
		return nil, err
	}

	return res.(*v1.GetImageByQueryKVsAndPageAndOrderByDESCReply), nil
}
