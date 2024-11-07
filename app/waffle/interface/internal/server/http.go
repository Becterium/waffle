package server

import (
	"github.com/go-kratos/kratos/v2/middleware/logging"
	v1 "waffle/api/waffle/interface/v1"
	"waffle/app/waffle/interface/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"waffle/app/waffle/interface/internal/conf"
	//jwt5 "github.com/golang-jwt/jwt/v5"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, waffle *service.WaffleInterface, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			//jwt.Server(func(token *jwt5.Token) (interface{}, error) {
			//	return []byte(ac.ApiKey), nil
			//}, jwt.WithSigningMethod(jwt5.SigningMethodHS256)),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterWaffleInterfaceHTTPServer(srv, waffle)
	return srv
}
