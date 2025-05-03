package utils

import (
	"context"
	"gin-second-fish/global"
)

// toggle favorite lua script
var toggleFavoriteScript = `
-- 原子化处理收藏和取消收藏操作
-- KEYS[1]: 用户收藏键 user:{userId}:favorite
-- KEYS[2]: 物品收藏键 item:{itemId}:favorite
-- KEYS[3]: 收藏排行榜键 ranking:favorite
-- ARGV[1]: 物品ID itemId
-- ARGV[2]: 用户ID userId
-- 返回: 1 表示添加收藏, 0 表示取消收藏
local userKey = KEYS[1]
local itemKey = KEYS[2]
local rankingKey = KEYS[3]
local itemId = ARGV[1]
local userId = ARGV[2]

-- 检查用户是否已收藏该物品
local isMember = redis.call("SISMEMBER", userKey, itemId)

if isMember == 1 then
    -- 取消收藏
    redis.call("DECR", userKey..":count")
    redis.call("DECR", itemKey..":count")
    redis.call("ZINCRBY", rankingKey, -1, itemId)
    redis.call("SREM", userKey, itemId)
    return 0
else
    -- 添加收藏
    redis.call("INCR", userKey..":count")
    redis.call("INCR", itemKey..":count")
    redis.call("ZINCRBY", rankingKey, 1, itemId)
    redis.call("SADD", userKey, itemId)
    return 1
end
`

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
// use lua to deal with high concurrency
func SetFavorite(userId string, itemId string) error {
	ctx := context.Background()
	// check if userId in set
	userKey := global.RedisKeyPrefix + "user:" + userId + ":favorite"
	itemKey := global.RedisKeyPrefix + "item:" + itemId + ":favorite"
	favoriteRankingKey := global.RedisKeyPrefix + "ranking:favorite"

	// 1. use basic redis command to toggle favorite
	// isMember, err := global.Rdb.SIsMember(ctx, userKey, itemId).Result()
	// if err != nil {
	// 	return err
	// }
	// if isMember {
	// 	global.Rdb.Decr(ctx, userKey+":count").Err()
	// 	global.Rdb.Decr(ctx, itemKey+":count").Err()
	// 	global.Rdb.ZIncrBy(ctx, favoriteRankingKey, -1, itemId).Err()
	// 	return global.Rdb.SRem(ctx, userKey, itemId).Err()
	// } else {
	// 	global.Rdb.Incr(ctx, userKey+":count").Err()
	// 	global.Rdb.Incr(ctx, itemKey+":count").Err()
	// 	global.Rdb.ZIncrBy(ctx, favoriteRankingKey, 1, itemId).Err()
	// 	return global.Rdb.SAdd(ctx, userKey, itemId).Err()
	// }

	// 2. use lua script
	_, err := global.Rdb.Eval(ctx, toggleFavoriteScript, []string{
		userKey,
		itemKey,
		favoriteRankingKey,
	}, itemId, userId).Result()

	return err
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
