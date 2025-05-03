package utils

import (
	"context"
	"gin-second-fish/global"
	"time"
)

// set token to redis
func SetToken(userId string, token string) error {
	ctx := context.Background()
	// set token to redis
	key := global.RedisKeyPrefix + "user:" + userId + ":token"
	expire := 24 * time.Hour
	return global.Rdb.Set(ctx, key, token, expire).Err()
}

// check if user is in item favorite list
func IsFavorite(userId string, itemId string) (bool, error) {
	ctx := context.Background()
	// check if userId in set
	key := global.RedisKeyPrefix + "item:" + itemId + ":favorite"
	isMember, err := global.Rdb.SIsMember(ctx, key, itemId).Result()
	if err != nil {
		return false, err
	}
	return isMember, nil
}

// toggle favorite list
// update the favorite count in redis
func SetFavorite(userId string, itemId string) error {
	ctx := context.Background()
	// check if userId in set
	userKey := global.RedisKeyPrefix + "user:" + userId + ":favorite"
	itemKey := global.RedisKeyPrefix + "item:" + itemId + ":favorite"
	isMember, err := global.Rdb.SIsMember(ctx, userKey, itemId).Result()
	if err != nil {
		return err
	}
	if isMember {
		// update favorite count in user favorite list
		global.Rdb.Decr(ctx, userKey+":count").Err()
		// update favorite count in item favorite list
		global.Rdb.Decr(ctx, itemKey+":count").Err()
		// remove item from favorite list
		return global.Rdb.SRem(ctx, userKey, itemId).Err()
	} else {
		// update favorite count in user favorite list
		global.Rdb.Incr(ctx, userKey+":count").Err()
		// update favorite count in item favorite list
		global.Rdb.Incr(ctx, itemKey+":count").Err()
		// add item to favorite list
		return global.Rdb.SAdd(ctx, userKey, itemId).Err()
	}
}
