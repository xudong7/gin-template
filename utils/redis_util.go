package utils

import (
	"context"
	"gin-second-fish/global"
	"time"
)

func SetToken(key string, token string) error {
	// set token to redis
	ctx := context.Background()
	expire := 24 * time.Hour
	return global.Rdb.Set(ctx, key, token, expire).Err()
}
