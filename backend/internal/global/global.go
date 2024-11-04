package global

import (
	"github.com/jinzhu/gorm"
	"github.com/redis/go-redis/v9"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
)
