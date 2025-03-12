package util

import (
	"math/rand"
	"time"
)

func RedisImageTagCreateRandTime() time.Duration {
	return time.Hour*24 +
		time.Hour*time.Duration(rand.Intn(24)) +
		time.Minute*time.Duration(rand.Intn(60)) +
		time.Second*time.Duration(rand.Intn(60))
}

func RedisCacheNullTime() time.Duration {
	return time.Minute * 5
}
