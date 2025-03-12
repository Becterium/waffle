package util

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/metadata"
	"strconv"
)

func MetadataGetUserIdFromMetaData(ctx context.Context) (uint, error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return 0, errors.New("metadata can't get context from server")
	}
	userId := md.Get("x-md-local-userId")
	id, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("parse user_id fail, error: %s", err))
	}
	return uint(id), nil
}
