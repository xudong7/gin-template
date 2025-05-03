package utils

import (
	"context"
	"gin-second-fish/global"
)

// check if user is in item favorite list
func IsFavorite(userId string, itemId string) (bool, error) {
	ctx := context.Background()
	// check if userId in set
	key := global.RedisKeyPrefix + "item:" + itemId + ":favorite"
	isMember, err := global.Rdb.SIsMember(ctx, key, userId).Result()
	if err != nil {
		return false, err
	}
	return isMember, nil
}

// toggle favorite list
// update the favorite count in redis
// TODO 高并发下如何处理？
func SetFavorite(userId string, itemId string) error {
	ctx := context.Background()
	// check if userId in set
	userKey := global.RedisKeyPrefix + "user:" + userId + ":favorite"
	itemKey := global.RedisKeyPrefix + "item:" + itemId + ":favorite"
	favoriteRankingKey := global.RedisKeyPrefix + "ranking:favorite"
	isMember, err := global.Rdb.SIsMember(ctx, userKey, itemId).Result()
	if err != nil {
		return err
	}
	if isMember {
		global.Rdb.Decr(ctx, userKey+":count").Err()
		global.Rdb.Decr(ctx, itemKey+":count").Err()
		global.Rdb.ZIncrBy(ctx, favoriteRankingKey, -1, itemId).Err()
		return global.Rdb.SRem(ctx, userKey, itemId).Err()
	} else {
		global.Rdb.Incr(ctx, userKey+":count").Err()
		global.Rdb.Incr(ctx, itemKey+":count").Err()
		global.Rdb.ZIncrBy(ctx, favoriteRankingKey, 1, itemId).Err()
		return global.Rdb.SAdd(ctx, userKey, itemId).Err()
	}
}

// get top favorite items
func GetTopFavoriteItems(limit int64) ([]string, error) {
	ctx := context.Background()
	favoriteRankingKey := global.RedisKeyPrefix + "ranking:favorite"
	// get top 10 favorite items
	topItems, err := global.Rdb.ZRevRange(ctx, favoriteRankingKey, 0, limit-1).Result()
	if err != nil {
		return nil, err
	}
	return topItems, nil
}
