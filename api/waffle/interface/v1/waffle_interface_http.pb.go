// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.28.2
// source: v1/waffle_interface.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationWaffleInterfaceAddImageTag = "/waffle.interface.v1.WaffleInterface/AddImageTag"
const OperationWaffleInterfaceGenerateUploadAvatarUrl = "/waffle.interface.v1.WaffleInterface/GenerateUploadAvatarUrl"
const OperationWaffleInterfaceGenerateUploadImgUrl = "/waffle.interface.v1.WaffleInterface/GenerateUploadImgUrl"
const OperationWaffleInterfaceGetImage = "/waffle.interface.v1.WaffleInterface/GetImage"
const OperationWaffleInterfaceGetImageByQueryKVsAndPageAndOrderByDESC = "/waffle.interface.v1.WaffleInterface/GetImageByQueryKVsAndPageAndOrderByDESC"
const OperationWaffleInterfaceLogin = "/waffle.interface.v1.WaffleInterface/Login"
const OperationWaffleInterfaceLogout = "/waffle.interface.v1.WaffleInterface/Logout"
const OperationWaffleInterfacePing = "/waffle.interface.v1.WaffleInterface/Ping"
const OperationWaffleInterfacePingRPC = "/waffle.interface.v1.WaffleInterface/PingRPC"
const OperationWaffleInterfaceRegister = "/waffle.interface.v1.WaffleInterface/Register"
const OperationWaffleInterfaceReloadCategoryRedisImageTag = "/waffle.interface.v1.WaffleInterface/ReloadCategoryRedisImageTag"
const OperationWaffleInterfaceSearchImageTagByNameLike = "/waffle.interface.v1.WaffleInterface/SearchImageTagByNameLike"
const OperationWaffleInterfaceVerifyAvatarUpload = "/waffle.interface.v1.WaffleInterface/VerifyAvatarUpload"
const OperationWaffleInterfaceVerifyImagesUpload = "/waffle.interface.v1.WaffleInterface/VerifyImagesUpload"

type WaffleInterfaceHTTPServer interface {
	// AddImageTag image - tag
	AddImageTag(context.Context, *AddImageTagReq) (*AddImageTagReply, error)
	GenerateUploadAvatarUrl(context.Context, *GenerateUploadAvatarUrlReq) (*GenerateUploadAvatarUrlReply, error)
	// GenerateUploadImgUrl media
	GenerateUploadImgUrl(context.Context, *GenerateUploadImgUrlReq) (*GenerateUploadImgUrlReply, error)
	GetImage(context.Context, *GetImageReq) (*GetImageReply, error)
	GetImageByQueryKVsAndPageAndOrderByDESC(context.Context, *GetImageByQueryKVsAndPageAndOrderByDESCReq) (*GetImageByQueryKVsAndPageAndOrderByDESCReply, error)
	Login(context.Context, *LoginReq) (*LoginReply, error)
	Logout(context.Context, *LogoutReq) (*LogoutReply, error)
	Ping(context.Context, *PingReq) (*PingReply, error)
	PingRPC(context.Context, *PingRPCReq) (*PingRPCReply, error)
	// Registeruser
	Register(context.Context, *RegisterReq) (*RegisterReply, error)
	ReloadCategoryRedisImageTag(context.Context, *ReloadCategoryRedisImageTagReq) (*ReloadCategoryRedisImageTagReply, error)
	SearchImageTagByNameLike(context.Context, *SearchImageTagByNameLikeReq) (*SearchImageTagByNameLikeReply, error)
	VerifyAvatarUpload(context.Context, *VerifyAvatarUploadReq) (*VerifyAvatarUploadReply, error)
	VerifyImagesUpload(context.Context, *VerifyImagesUploadReq) (*VerifyImagesUploadReply, error)
}

func RegisterWaffleInterfaceHTTPServer(s *http.Server, srv WaffleInterfaceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/register", _WaffleInterface_Register0_HTTP_Handler(srv))
	r.POST("/v1/login", _WaffleInterface_Login0_HTTP_Handler(srv))
	r.POST("/v1/logout", _WaffleInterface_Logout0_HTTP_Handler(srv))
	r.POST("/v1/Ping", _WaffleInterface_Ping0_HTTP_Handler(srv))
	r.POST("/v1/PingRPC", _WaffleInterface_PingRPC0_HTTP_Handler(srv))
	r.POST("/v1/GenerateUploadImgUrl", _WaffleInterface_GenerateUploadImgUrl0_HTTP_Handler(srv))
	r.POST("/v1/GenerateUploadAvatarUrl", _WaffleInterface_GenerateUploadAvatarUrl0_HTTP_Handler(srv))
	r.POST("/v1/VerifyImagesUpload", _WaffleInterface_VerifyImagesUpload0_HTTP_Handler(srv))
	r.POST("/v1/VerifyAvatarUpload", _WaffleInterface_VerifyAvatarUpload0_HTTP_Handler(srv))
	r.GET("/v1/w/{uid}", _WaffleInterface_GetImage0_HTTP_Handler(srv))
	r.POST("/v1/AddImageTag", _WaffleInterface_AddImageTag0_HTTP_Handler(srv))
	r.POST("/v1/SearchImageTagByNameLike", _WaffleInterface_SearchImageTagByNameLike0_HTTP_Handler(srv))
	r.POST("/v1/ReloadCategoryRedisImageTag", _WaffleInterface_ReloadCategoryRedisImageTag0_HTTP_Handler(srv))
	r.POST("/v1/SortAndQueryImage", _WaffleInterface_GetImageByQueryKVsAndPageAndOrderByDESC0_HTTP_Handler(srv))
}

func _WaffleInterface_Register0_HTTP_Handler(srv WaffleInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWaffleInterfaceRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _WaffleInterface_Login0_HTTP_Handler(srv WaffleInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWaffleInterfaceLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _WaffleInterface_Logout0_HTTP_Handler(srv WaffleInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWaffleInterfaceLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutReply)
		return ctx.Result(200, reply)
	}
}

func _WaffleInterface_Ping0_HTTP_Handler(srv WaffleInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PingReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWaffleInterfacePing)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*PingReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PingReply)
		return ctx.Result(200, reply)
	}
}

func _WaffleInterface_PingRPC0_HTTP_Handler(srv WaffleInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PingRPCReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWaffleInterfacePingRPC)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PingRPC(ctx, req.(*PingRPCReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PingRPCReply)
		return ctx.Result(200, reply)
	}
}

func _WaffleInterface_GenerateUploadImgUrl0_HTTP_Handler(srv WaffleInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GenerateUploadImgUrlReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWaffleInterfaceGenerateUploadImgUrl)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GenerateUploadImgUrl(ctx, req.(*GenerateUploadImgUrlReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GenerateUploadImgUrlReply)
		return ctx.Result(200, reply)
	}
}

func _WaffleInterface_GenerateUploadAvatarUrl0_HTTP_Handler(srv WaffleInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GenerateUploadAvatarUrlReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWaffleInterfaceGenerateUploadAvatarUrl)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GenerateUploadAvatarUrl(ctx, req.(*GenerateUploadAvatarUrlReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GenerateUploadAvatarUrlReply)
		return ctx.Result(200, reply)
	}
}

func _WaffleInterface_VerifyImagesUpload0_HTTP_Handler(srv WaffleInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in VerifyImagesUploadReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWaffleInterfaceVerifyImagesUpload)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.VerifyImagesUpload(ctx, req.(*VerifyImagesUploadReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VerifyImagesUploadReply)
		return ctx.Result(200, reply)
	}
}

func _WaffleInterface_VerifyAvatarUpload0_HTTP_Handler(srv WaffleInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in VerifyAvatarUploadReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWaffleInterfaceVerifyAvatarUpload)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.VerifyAvatarUpload(ctx, req.(*VerifyAvatarUploadReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VerifyAvatarUploadReply)
		return ctx.Result(200, reply)
	}
}

func _WaffleInterface_GetImage0_HTTP_Handler(srv WaffleInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetImageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWaffleInterfaceGetImage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetImage(ctx, req.(*GetImageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetImageReply)
		return ctx.Result(200, reply)
	}
}

func _WaffleInterface_AddImageTag0_HTTP_Handler(srv WaffleInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddImageTagReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWaffleInterfaceAddImageTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddImageTag(ctx, req.(*AddImageTagReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddImageTagReply)
		return ctx.Result(200, reply)
	}
}

func _WaffleInterface_SearchImageTagByNameLike0_HTTP_Handler(srv WaffleInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SearchImageTagByNameLikeReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWaffleInterfaceSearchImageTagByNameLike)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SearchImageTagByNameLike(ctx, req.(*SearchImageTagByNameLikeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SearchImageTagByNameLikeReply)
		return ctx.Result(200, reply)
	}
}

func _WaffleInterface_ReloadCategoryRedisImageTag0_HTTP_Handler(srv WaffleInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReloadCategoryRedisImageTagReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWaffleInterfaceReloadCategoryRedisImageTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ReloadCategoryRedisImageTag(ctx, req.(*ReloadCategoryRedisImageTagReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ReloadCategoryRedisImageTagReply)
		return ctx.Result(200, reply)
	}
}

func _WaffleInterface_GetImageByQueryKVsAndPageAndOrderByDESC0_HTTP_Handler(srv WaffleInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetImageByQueryKVsAndPageAndOrderByDESCReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWaffleInterfaceGetImageByQueryKVsAndPageAndOrderByDESC)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetImageByQueryKVsAndPageAndOrderByDESC(ctx, req.(*GetImageByQueryKVsAndPageAndOrderByDESCReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetImageByQueryKVsAndPageAndOrderByDESCReply)
		return ctx.Result(200, reply)
	}
}

type WaffleInterfaceHTTPClient interface {
	AddImageTag(ctx context.Context, req *AddImageTagReq, opts ...http.CallOption) (rsp *AddImageTagReply, err error)
	GenerateUploadAvatarUrl(ctx context.Context, req *GenerateUploadAvatarUrlReq, opts ...http.CallOption) (rsp *GenerateUploadAvatarUrlReply, err error)
	GenerateUploadImgUrl(ctx context.Context, req *GenerateUploadImgUrlReq, opts ...http.CallOption) (rsp *GenerateUploadImgUrlReply, err error)
	GetImage(ctx context.Context, req *GetImageReq, opts ...http.CallOption) (rsp *GetImageReply, err error)
	GetImageByQueryKVsAndPageAndOrderByDESC(ctx context.Context, req *GetImageByQueryKVsAndPageAndOrderByDESCReq, opts ...http.CallOption) (rsp *GetImageByQueryKVsAndPageAndOrderByDESCReply, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginReply, err error)
	Logout(ctx context.Context, req *LogoutReq, opts ...http.CallOption) (rsp *LogoutReply, err error)
	Ping(ctx context.Context, req *PingReq, opts ...http.CallOption) (rsp *PingReply, err error)
	PingRPC(ctx context.Context, req *PingRPCReq, opts ...http.CallOption) (rsp *PingRPCReply, err error)
	Register(ctx context.Context, req *RegisterReq, opts ...http.CallOption) (rsp *RegisterReply, err error)
	ReloadCategoryRedisImageTag(ctx context.Context, req *ReloadCategoryRedisImageTagReq, opts ...http.CallOption) (rsp *ReloadCategoryRedisImageTagReply, err error)
	SearchImageTagByNameLike(ctx context.Context, req *SearchImageTagByNameLikeReq, opts ...http.CallOption) (rsp *SearchImageTagByNameLikeReply, err error)
	VerifyAvatarUpload(ctx context.Context, req *VerifyAvatarUploadReq, opts ...http.CallOption) (rsp *VerifyAvatarUploadReply, err error)
	VerifyImagesUpload(ctx context.Context, req *VerifyImagesUploadReq, opts ...http.CallOption) (rsp *VerifyImagesUploadReply, err error)
}

type WaffleInterfaceHTTPClientImpl struct {
	cc *http.Client
}

func NewWaffleInterfaceHTTPClient(client *http.Client) WaffleInterfaceHTTPClient {
	return &WaffleInterfaceHTTPClientImpl{client}
}

func (c *WaffleInterfaceHTTPClientImpl) AddImageTag(ctx context.Context, in *AddImageTagReq, opts ...http.CallOption) (*AddImageTagReply, error) {
	var out AddImageTagReply
	pattern := "/v1/AddImageTag"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWaffleInterfaceAddImageTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WaffleInterfaceHTTPClientImpl) GenerateUploadAvatarUrl(ctx context.Context, in *GenerateUploadAvatarUrlReq, opts ...http.CallOption) (*GenerateUploadAvatarUrlReply, error) {
	var out GenerateUploadAvatarUrlReply
	pattern := "/v1/GenerateUploadAvatarUrl"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWaffleInterfaceGenerateUploadAvatarUrl))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WaffleInterfaceHTTPClientImpl) GenerateUploadImgUrl(ctx context.Context, in *GenerateUploadImgUrlReq, opts ...http.CallOption) (*GenerateUploadImgUrlReply, error) {
	var out GenerateUploadImgUrlReply
	pattern := "/v1/GenerateUploadImgUrl"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWaffleInterfaceGenerateUploadImgUrl))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WaffleInterfaceHTTPClientImpl) GetImage(ctx context.Context, in *GetImageReq, opts ...http.CallOption) (*GetImageReply, error) {
	var out GetImageReply
	pattern := "/v1/w/{uid}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWaffleInterfaceGetImage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WaffleInterfaceHTTPClientImpl) GetImageByQueryKVsAndPageAndOrderByDESC(ctx context.Context, in *GetImageByQueryKVsAndPageAndOrderByDESCReq, opts ...http.CallOption) (*GetImageByQueryKVsAndPageAndOrderByDESCReply, error) {
	var out GetImageByQueryKVsAndPageAndOrderByDESCReply
	pattern := "/v1/SortAndQueryImage"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWaffleInterfaceGetImageByQueryKVsAndPageAndOrderByDESC))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WaffleInterfaceHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWaffleInterfaceLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WaffleInterfaceHTTPClientImpl) Logout(ctx context.Context, in *LogoutReq, opts ...http.CallOption) (*LogoutReply, error) {
	var out LogoutReply
	pattern := "/v1/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWaffleInterfaceLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WaffleInterfaceHTTPClientImpl) Ping(ctx context.Context, in *PingReq, opts ...http.CallOption) (*PingReply, error) {
	var out PingReply
	pattern := "/v1/Ping"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWaffleInterfacePing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WaffleInterfaceHTTPClientImpl) PingRPC(ctx context.Context, in *PingRPCReq, opts ...http.CallOption) (*PingRPCReply, error) {
	var out PingRPCReply
	pattern := "/v1/PingRPC"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWaffleInterfacePingRPC))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WaffleInterfaceHTTPClientImpl) Register(ctx context.Context, in *RegisterReq, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/v1/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWaffleInterfaceRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WaffleInterfaceHTTPClientImpl) ReloadCategoryRedisImageTag(ctx context.Context, in *ReloadCategoryRedisImageTagReq, opts ...http.CallOption) (*ReloadCategoryRedisImageTagReply, error) {
	var out ReloadCategoryRedisImageTagReply
	pattern := "/v1/ReloadCategoryRedisImageTag"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWaffleInterfaceReloadCategoryRedisImageTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WaffleInterfaceHTTPClientImpl) SearchImageTagByNameLike(ctx context.Context, in *SearchImageTagByNameLikeReq, opts ...http.CallOption) (*SearchImageTagByNameLikeReply, error) {
	var out SearchImageTagByNameLikeReply
	pattern := "/v1/SearchImageTagByNameLike"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWaffleInterfaceSearchImageTagByNameLike))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WaffleInterfaceHTTPClientImpl) VerifyAvatarUpload(ctx context.Context, in *VerifyAvatarUploadReq, opts ...http.CallOption) (*VerifyAvatarUploadReply, error) {
	var out VerifyAvatarUploadReply
	pattern := "/v1/VerifyAvatarUpload"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWaffleInterfaceVerifyAvatarUpload))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WaffleInterfaceHTTPClientImpl) VerifyImagesUpload(ctx context.Context, in *VerifyImagesUploadReq, opts ...http.CallOption) (*VerifyImagesUploadReply, error) {
	var out VerifyImagesUploadReply
	pattern := "/v1/VerifyImagesUpload"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWaffleInterfaceVerifyImagesUpload))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
