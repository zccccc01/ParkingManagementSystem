package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	// 从环境变量中获取数据库连接信息
	dbHost := os.Getenv("DB_HOST")         // 从环境变量读取数据库主机地址
	dbPort := os.Getenv("DB_PORT")         // 从环境变量读取数据库端口
	dbUser := os.Getenv("DB_USER")         // 从环境变量读取数据库用户名
	dbPassword := os.Getenv("DB_PASSWORD") // 从环境变量读取数据库密码
	dbName := os.Getenv("DB_NAME")         // 从环境变量读取数据库名称
	dbRole := os.Getenv("DB_ROLE")         // 从环境变量读取数据库角色 (主/从)

	// 设置数据库连接字符串，支持通过环境变量配置
	if dbHost == "" {
		dbHost = "localhost" // 默认值
	}
	if dbPort == "" {
		dbPort = "3306" // 默认值
	}
	if dbUser == "" {
		dbUser = "root" // 默认值
	}
	if dbPassword == "" {
		dbPassword = "123456" // 默认值
	}
	if dbName == "" {
		dbName = "mydb" // 默认值
	}
	if dbRole == "" {
		dbRole = "master" // 默认连接到主库
	}

	var dbHostToUse string
	if dbRole == "slave" {
		dbHostToUse = os.Getenv("DB_SLAVE_HOST")
		if dbHostToUse == "" {
			dbHostToUse = "db_slave1"
		}
	} else {
		dbHostToUse = dbHost // 使用主库地址
	}

	// 设置数据库连接字符串
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHostToUse + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接数据库
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
