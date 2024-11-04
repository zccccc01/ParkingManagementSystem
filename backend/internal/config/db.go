package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/global"
)

func InitDB() {
	// 连接数据库
	db, err := gorm.Open("mysql", "root:123456@/chao_db?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	// 开debug模式
	db = db.Debug()

	global.DB = db
}
