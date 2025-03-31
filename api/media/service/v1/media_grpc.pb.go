// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: v1/media.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Media_UploadImages_FullMethodName                 = "/waffle.media.v1.Media/UploadImages"
	Media_UploadUserImage_FullMethodName              = "/waffle.media.v1.Media/UploadUserImage"
	Media_VerifyUserImageUpload_FullMethodName        = "/waffle.media.v1.Media/VerifyUserImageUpload"
	Media_VerifyImagesUpload_FullMethodName           = "/waffle.media.v1.Media/VerifyImagesUpload"
	Media_GetImage_FullMethodName                     = "/waffle.media.v1.Media/GetImage"
	Media_AddImageTag_FullMethodName                  = "/waffle.media.v1.Media/AddImageTag"
	Media_SearchImageTagByNameLike_FullMethodName     = "/waffle.media.v1.Media/SearchImageTagByNameLike"
	Media_ReloadCategoryRedisImageTag_FullMethodName  = "/waffle.media.v1.Media/ReloadCategoryRedisImageTag"
	Media_CreateCollection_FullMethodName             = "/waffle.media.v1.Media/CreateCollection"
	Media_StarImage_FullMethodName                    = "/waffle.media.v1.Media/StarImage"
	Media_UnStarImage_FullMethodName                  = "/waffle.media.v1.Media/UnStarImage"
	Media_FindCollectionByImageId_FullMethodName      = "/waffle.media.v1.Media/FindCollectionByImageId"
	Media_FindCollectionByCollectionId_FullMethodName = "/waffle.media.v1.Media/FindCollectionByCollectionId"
	Media_UploadVideo_FullMethodName                  = "/waffle.media.v1.Media/UploadVideo"
	Media_GetVideo_FullMethodName                     = "/waffle.media.v1.Media/GetVideo"
)

// MediaClient is the client API for Media service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MediaClient interface {
	// image
	UploadImages(ctx context.Context, in *UploadImagesReq, opts ...grpc.CallOption) (*UploadImagesReply, error)
	UploadUserImage(ctx context.Context, in *UploadUserImageReq, opts ...grpc.CallOption) (*UploadUserImageReply, error)
	VerifyUserImageUpload(ctx context.Context, in *VerifyUserImageUploadReq, opts ...grpc.CallOption) (*VerifyUserImageUploadReply, error)
	VerifyImagesUpload(ctx context.Context, in *VerifyImagesUploadReq, opts ...grpc.CallOption) (*VerifyImagesUploadReply, error)
	GetImage(ctx context.Context, in *GetImageReq, opts ...grpc.CallOption) (*GetImageReply, error)
	// image - tag
	AddImageTag(ctx context.Context, in *AddImageTagReq, opts ...grpc.CallOption) (*AddImageTagReply, error)
	SearchImageTagByNameLike(ctx context.Context, in *SearchImageTagByNameLikeReq, opts ...grpc.CallOption) (*SearchImageTagByNameLikeReply, error)
	ReloadCategoryRedisImageTag(ctx context.Context, in *ReloadCategoryRedisImageTagReq, opts ...grpc.CallOption) (*ReloadCategoryRedisImageTagReply, error)
	// collection
	CreateCollection(ctx context.Context, in *CreateCollectionReq, opts ...grpc.CallOption) (*CreateCollectionReply, error)
	StarImage(ctx context.Context, in *StarImageReq, opts ...grpc.CallOption) (*StarImageReply, error)
	UnStarImage(ctx context.Context, in *UnStarImageReq, opts ...grpc.CallOption) (*UnStarImageReply, error)
	FindCollectionByImageId(ctx context.Context, in *FindCollectionByImageIdReq, opts ...grpc.CallOption) (*FindCollectionByImageIdReply, error)
	FindCollectionByCollectionId(ctx context.Context, in *FindCollectionByCollectionIdReq, opts ...grpc.CallOption) (*FindCollectionByCollectionIdReply, error)
	// video
	UploadVideo(ctx context.Context, in *UpLoadVideoReq, opts ...grpc.CallOption) (*UpLoadVideoReply, error)
	GetVideo(ctx context.Context, in *GetVideoReq, opts ...grpc.CallOption) (*GetVideoReply, error)
}

type mediaClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaClient(cc grpc.ClientConnInterface) MediaClient {
	return &mediaClient{cc}
}

func (c *mediaClient) UploadImages(ctx context.Context, in *UploadImagesReq, opts ...grpc.CallOption) (*UploadImagesReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadImagesReply)
	err := c.cc.Invoke(ctx, Media_UploadImages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) UploadUserImage(ctx context.Context, in *UploadUserImageReq, opts ...grpc.CallOption) (*UploadUserImageReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadUserImageReply)
	err := c.cc.Invoke(ctx, Media_UploadUserImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) VerifyUserImageUpload(ctx context.Context, in *VerifyUserImageUploadReq, opts ...grpc.CallOption) (*VerifyUserImageUploadReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyUserImageUploadReply)
	err := c.cc.Invoke(ctx, Media_VerifyUserImageUpload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) VerifyImagesUpload(ctx context.Context, in *VerifyImagesUploadReq, opts ...grpc.CallOption) (*VerifyImagesUploadReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyImagesUploadReply)
	err := c.cc.Invoke(ctx, Media_VerifyImagesUpload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) GetImage(ctx context.Context, in *GetImageReq, opts ...grpc.CallOption) (*GetImageReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetImageReply)
	err := c.cc.Invoke(ctx, Media_GetImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) AddImageTag(ctx context.Context, in *AddImageTagReq, opts ...grpc.CallOption) (*AddImageTagReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddImageTagReply)
	err := c.cc.Invoke(ctx, Media_AddImageTag_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) SearchImageTagByNameLike(ctx context.Context, in *SearchImageTagByNameLikeReq, opts ...grpc.CallOption) (*SearchImageTagByNameLikeReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchImageTagByNameLikeReply)
	err := c.cc.Invoke(ctx, Media_SearchImageTagByNameLike_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) ReloadCategoryRedisImageTag(ctx context.Context, in *ReloadCategoryRedisImageTagReq, opts ...grpc.CallOption) (*ReloadCategoryRedisImageTagReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReloadCategoryRedisImageTagReply)
	err := c.cc.Invoke(ctx, Media_ReloadCategoryRedisImageTag_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) CreateCollection(ctx context.Context, in *CreateCollectionReq, opts ...grpc.CallOption) (*CreateCollectionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCollectionReply)
	err := c.cc.Invoke(ctx, Media_CreateCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) StarImage(ctx context.Context, in *StarImageReq, opts ...grpc.CallOption) (*StarImageReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StarImageReply)
	err := c.cc.Invoke(ctx, Media_StarImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) UnStarImage(ctx context.Context, in *UnStarImageReq, opts ...grpc.CallOption) (*UnStarImageReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnStarImageReply)
	err := c.cc.Invoke(ctx, Media_UnStarImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) FindCollectionByImageId(ctx context.Context, in *FindCollectionByImageIdReq, opts ...grpc.CallOption) (*FindCollectionByImageIdReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindCollectionByImageIdReply)
	err := c.cc.Invoke(ctx, Media_FindCollectionByImageId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) FindCollectionByCollectionId(ctx context.Context, in *FindCollectionByCollectionIdReq, opts ...grpc.CallOption) (*FindCollectionByCollectionIdReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindCollectionByCollectionIdReply)
	err := c.cc.Invoke(ctx, Media_FindCollectionByCollectionId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) UploadVideo(ctx context.Context, in *UpLoadVideoReq, opts ...grpc.CallOption) (*UpLoadVideoReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpLoadVideoReply)
	err := c.cc.Invoke(ctx, Media_UploadVideo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) GetVideo(ctx context.Context, in *GetVideoReq, opts ...grpc.CallOption) (*GetVideoReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVideoReply)
	err := c.cc.Invoke(ctx, Media_GetVideo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaServer is the server API for Media service.
// All implementations must embed UnimplementedMediaServer
// for forward compatibility.
type MediaServer interface {
	// image
	UploadImages(context.Context, *UploadImagesReq) (*UploadImagesReply, error)
	UploadUserImage(context.Context, *UploadUserImageReq) (*UploadUserImageReply, error)
	VerifyUserImageUpload(context.Context, *VerifyUserImageUploadReq) (*VerifyUserImageUploadReply, error)
	VerifyImagesUpload(context.Context, *VerifyImagesUploadReq) (*VerifyImagesUploadReply, error)
	GetImage(context.Context, *GetImageReq) (*GetImageReply, error)
	// image - tag
	AddImageTag(context.Context, *AddImageTagReq) (*AddImageTagReply, error)
	SearchImageTagByNameLike(context.Context, *SearchImageTagByNameLikeReq) (*SearchImageTagByNameLikeReply, error)
	ReloadCategoryRedisImageTag(context.Context, *ReloadCategoryRedisImageTagReq) (*ReloadCategoryRedisImageTagReply, error)
	// collection
	CreateCollection(context.Context, *CreateCollectionReq) (*CreateCollectionReply, error)
	StarImage(context.Context, *StarImageReq) (*StarImageReply, error)
	UnStarImage(context.Context, *UnStarImageReq) (*UnStarImageReply, error)
	FindCollectionByImageId(context.Context, *FindCollectionByImageIdReq) (*FindCollectionByImageIdReply, error)
	FindCollectionByCollectionId(context.Context, *FindCollectionByCollectionIdReq) (*FindCollectionByCollectionIdReply, error)
	// video
	UploadVideo(context.Context, *UpLoadVideoReq) (*UpLoadVideoReply, error)
	GetVideo(context.Context, *GetVideoReq) (*GetVideoReply, error)
	mustEmbedUnimplementedMediaServer()
}

// UnimplementedMediaServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMediaServer struct{}

func (UnimplementedMediaServer) UploadImages(context.Context, *UploadImagesReq) (*UploadImagesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadImages not implemented")
}
func (UnimplementedMediaServer) UploadUserImage(context.Context, *UploadUserImageReq) (*UploadUserImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadUserImage not implemented")
}
func (UnimplementedMediaServer) VerifyUserImageUpload(context.Context, *VerifyUserImageUploadReq) (*VerifyUserImageUploadReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUserImageUpload not implemented")
}
func (UnimplementedMediaServer) VerifyImagesUpload(context.Context, *VerifyImagesUploadReq) (*VerifyImagesUploadReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyImagesUpload not implemented")
}
func (UnimplementedMediaServer) GetImage(context.Context, *GetImageReq) (*GetImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedMediaServer) AddImageTag(context.Context, *AddImageTagReq) (*AddImageTagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddImageTag not implemented")
}
func (UnimplementedMediaServer) SearchImageTagByNameLike(context.Context, *SearchImageTagByNameLikeReq) (*SearchImageTagByNameLikeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchImageTagByNameLike not implemented")
}
func (UnimplementedMediaServer) ReloadCategoryRedisImageTag(context.Context, *ReloadCategoryRedisImageTagReq) (*ReloadCategoryRedisImageTagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReloadCategoryRedisImageTag not implemented")
}
func (UnimplementedMediaServer) CreateCollection(context.Context, *CreateCollectionReq) (*CreateCollectionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollection not implemented")
}
func (UnimplementedMediaServer) StarImage(context.Context, *StarImageReq) (*StarImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StarImage not implemented")
}
func (UnimplementedMediaServer) UnStarImage(context.Context, *UnStarImageReq) (*UnStarImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnStarImage not implemented")
}
func (UnimplementedMediaServer) FindCollectionByImageId(context.Context, *FindCollectionByImageIdReq) (*FindCollectionByImageIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCollectionByImageId not implemented")
}
func (UnimplementedMediaServer) FindCollectionByCollectionId(context.Context, *FindCollectionByCollectionIdReq) (*FindCollectionByCollectionIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCollectionByCollectionId not implemented")
}
func (UnimplementedMediaServer) UploadVideo(context.Context, *UpLoadVideoReq) (*UpLoadVideoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadVideo not implemented")
}
func (UnimplementedMediaServer) GetVideo(context.Context, *GetVideoReq) (*GetVideoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideo not implemented")
}
func (UnimplementedMediaServer) mustEmbedUnimplementedMediaServer() {}
func (UnimplementedMediaServer) testEmbeddedByValue()               {}

// UnsafeMediaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MediaServer will
// result in compilation errors.
type UnsafeMediaServer interface {
	mustEmbedUnimplementedMediaServer()
}

func RegisterMediaServer(s grpc.ServiceRegistrar, srv MediaServer) {
	// If the following call pancis, it indicates UnimplementedMediaServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Media_ServiceDesc, srv)
}

func _Media_UploadImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadImagesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).UploadImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_UploadImages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).UploadImages(ctx, req.(*UploadImagesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_UploadUserImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadUserImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).UploadUserImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_UploadUserImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).UploadUserImage(ctx, req.(*UploadUserImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_VerifyUserImageUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUserImageUploadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).VerifyUserImageUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_VerifyUserImageUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).VerifyUserImageUpload(ctx, req.(*VerifyUserImageUploadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_VerifyImagesUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyImagesUploadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).VerifyImagesUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_VerifyImagesUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).VerifyImagesUpload(ctx, req.(*VerifyImagesUploadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_GetImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).GetImage(ctx, req.(*GetImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_AddImageTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddImageTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).AddImageTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_AddImageTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).AddImageTag(ctx, req.(*AddImageTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_SearchImageTagByNameLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchImageTagByNameLikeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).SearchImageTagByNameLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_SearchImageTagByNameLike_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).SearchImageTagByNameLike(ctx, req.(*SearchImageTagByNameLikeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_ReloadCategoryRedisImageTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadCategoryRedisImageTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).ReloadCategoryRedisImageTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_ReloadCategoryRedisImageTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).ReloadCategoryRedisImageTag(ctx, req.(*ReloadCategoryRedisImageTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCollectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_CreateCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).CreateCollection(ctx, req.(*CreateCollectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_StarImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StarImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).StarImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_StarImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).StarImage(ctx, req.(*StarImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_UnStarImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnStarImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).UnStarImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_UnStarImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).UnStarImage(ctx, req.(*UnStarImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_FindCollectionByImageId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCollectionByImageIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).FindCollectionByImageId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_FindCollectionByImageId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).FindCollectionByImageId(ctx, req.(*FindCollectionByImageIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_FindCollectionByCollectionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCollectionByCollectionIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).FindCollectionByCollectionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_FindCollectionByCollectionId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).FindCollectionByCollectionId(ctx, req.(*FindCollectionByCollectionIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_UploadVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpLoadVideoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).UploadVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_UploadVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).UploadVideo(ctx, req.(*UpLoadVideoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_GetVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).GetVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_GetVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).GetVideo(ctx, req.(*GetVideoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Media_ServiceDesc is the grpc.ServiceDesc for Media service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Media_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "waffle.media.v1.Media",
	HandlerType: (*MediaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadImages",
			Handler:    _Media_UploadImages_Handler,
		},
		{
			MethodName: "UploadUserImage",
			Handler:    _Media_UploadUserImage_Handler,
		},
		{
			MethodName: "VerifyUserImageUpload",
			Handler:    _Media_VerifyUserImageUpload_Handler,
		},
		{
			MethodName: "VerifyImagesUpload",
			Handler:    _Media_VerifyImagesUpload_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _Media_GetImage_Handler,
		},
		{
			MethodName: "AddImageTag",
			Handler:    _Media_AddImageTag_Handler,
		},
		{
			MethodName: "SearchImageTagByNameLike",
			Handler:    _Media_SearchImageTagByNameLike_Handler,
		},
		{
			MethodName: "ReloadCategoryRedisImageTag",
			Handler:    _Media_ReloadCategoryRedisImageTag_Handler,
		},
		{
			MethodName: "CreateCollection",
			Handler:    _Media_CreateCollection_Handler,
		},
		{
			MethodName: "StarImage",
			Handler:    _Media_StarImage_Handler,
		},
		{
			MethodName: "UnStarImage",
			Handler:    _Media_UnStarImage_Handler,
		},
		{
			MethodName: "FindCollectionByImageId",
			Handler:    _Media_FindCollectionByImageId_Handler,
		},
		{
			MethodName: "FindCollectionByCollectionId",
			Handler:    _Media_FindCollectionByCollectionId_Handler,
		},
		{
			MethodName: "UploadVideo",
			Handler:    _Media_UploadVideo_Handler,
		},
		{
			MethodName: "GetVideo",
			Handler:    _Media_GetVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/media.proto",
}
