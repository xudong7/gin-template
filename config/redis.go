package config

import (
	"context"
	"fmt"
	"gin-second-fish/global"
	"log"

	"github.com/redis/go-redis/v9"
)

func initRedis() {
	addr := fmt.Sprintf("%s:%s", AppConfig.Redis.Host, AppConfig.Redis.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: AppConfig.Redis.Password,
		DB:       AppConfig.Redis.DB,
	})

	// Test the connection
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Printf("Failed to connect to Redis: %v", err)
	} else {
		log.Println("Connected to Redis successfully")
	}

	global.Rd = client
}
