package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	v1 "waffle/api/waffle/interface/v1"
	"waffle/app/waffle/interface/internal/conf"
	"waffle/app/waffle/interface/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/waffle.interface.v1.WaffleInterface/Login"] = struct{}{}
	whiteList["/waffle.interface.v1.WaffleInterface/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, waffle *service.WaffleInterface, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			ratelimit.Server(),
			logging.Server(logger),
			selector.Server(
				jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
					return []byte(ac.JwtKey), nil
				}, jwt.WithSigningMethod(jwtv5.SigningMethodHS256), jwt.WithClaims(func() jwtv5.Claims {
					return jwtv5.MapClaims{}
				})),
			).
				Match(NewWhiteListMatcher()).
				Build(),
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
