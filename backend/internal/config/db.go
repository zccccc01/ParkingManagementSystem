package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	// 连接数据库
	// 请确保将以下信息替换为您的实际数据库用户、密码和数据库名
	dsn := "root:123456@tcp(10.1.0.20:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	mdb, err := gorm.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	// 开启 debug 模式
	db = mdb.Debug()
}

func GetDBInstance() *gorm.DB {
	return db
}
