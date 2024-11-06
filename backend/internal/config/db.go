package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	// 连接数据库
	mdb, err := gorm.Open("mysql", "root:123456@/chao_db?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	// 开debug模式
	db = mdb.Debug()
}

func GetDBInstance() *gorm.DB {
	return db
}
