package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"waffle/api/media/service/v1"
	"waffle/app/media/service/internal/conf"
	"waffle/app/media/service/internal/service"
	//jwt5 "github.com/golang-jwt/jwt/v5"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, ac *conf.Auth, logger log.Logger, s *service.MediaService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			ratelimit.Server(),
			recovery.Recovery(),
			logging.Server(logger),
			metadata.Server(),
			//jwt.Server(func(token *jwt5.Token) (interface{}, error) {
			//	return []byte(ac.Key), nil
			//}, jwt.WithSigningMethod(jwt5.SigningMethodHS256)),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterMediaServer(srv, s)
	return srv
}
