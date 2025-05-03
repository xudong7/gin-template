package global

import (
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	// mysql and redis connection
	Db  *gorm.DB
	Rdb *redis.Client

	// fields for different models
	ItemFields     = []string{"seller_id", "category_id"}
	CategoryFields = []string{"parent_id"}
	OrderFields    = []string{"buyer_id", "seller_id", "item_id", "address_id"}

	// constants
	JwtSecret     = []byte("secret")
	JwtExpireTime = time.Now().Add(time.Hour * 72).Unix()

	// redis key prefix
	RedisKeyPrefix = "second-fish:"
)
