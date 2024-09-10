package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql", "root:123456@/chao_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed to connect database: %v", err) // 使用日志记录错误，而不是panic
	}

	// 实例一个接口
	//vehicleRepo := repository.NewVehicleRepository(db)

	//TODO: violation_record_repository_impl.go 未review

}
