package util

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	IdempotencyKey = "Idempotency-key"
)

func UnMarshalIdempotentKey(ctx context.Context) (string, error) {
	if request, ok := http.RequestFromServerContext(ctx); ok {
		key := request.Header.Get(IdempotencyKey)
		return key, nil
	}
	return "", errors.New(fmt.Sprintf("Idempotency UnMarshal Key from http.Header fail"))
}

func VerifyIdempotencyByRedis(ctx context.Context, rc *redis.Client, key string) (bool, error) {
	result, err := rc.SetNX(ctx, key, true, 24*time.Hour).Result()
	if result == true {
		return true, nil
	} else if err == nil {
		return false, errors.New(fmt.Sprintf("VerifyIdempotencyByRedis fail error: key has exist"))
	} else {
		return false, errors.New(fmt.Sprintf("VerifyIdempotencyByRedis fail error: %s", err.Error()))
	}
}
