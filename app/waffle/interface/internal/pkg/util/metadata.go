package util

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

const (
	PrefixLocalMetadata  = "x-md-local-"
	PrefixGlobalMetadata = "x-md-global-"
)

// 在token中
func UnMarshalTokeToMetadata(ctx context.Context, prefix string, key string) (context.Context, error) {
	if tokenClaims, ok := jwt.FromContext(ctx); ok {
		if claims, ok := tokenClaims.(jwtv5.MapClaims); ok {
			// metadata前缀字符串拼接
			result := make([]byte, len(prefix)+len(key))
			copy(result, prefix)
			copy(result[len(prefix):], key)

			ctx = metadata.AppendToClientContext(ctx, string(result), fmt.Sprintf("%v", claims[key]))

			return ctx, nil
		}
		return nil, errors.New("the claims in context can't reflect to jwt.MapClaims")
	}
	return nil, errors.New("jwt can't get info from context")
}
