package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	// 连接Redis
	db := redis.NewClient(&redis.Options{
		Addr:     "10.1.0.30:6379",
		Password: "",
		DB:       0,
	})

	_, err := db.Ping(context.Background()).Result()

	if err != nil {
		log.Fatalf("Failed to connect to Redis, got error: %v", err)
	}

	rdb = db
}

func GetRDBInstance() *redis.Client {
	return rdb
}
