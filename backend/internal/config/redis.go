package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/global"
)

func InitRedis() {
	// 连接Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.66.130:6379",
		Password: "123456",
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()

	if err != nil {
		log.Fatalf("Failed to connect to Redis, got error: %v", err)
	}

	global.RDB = rdb
}
