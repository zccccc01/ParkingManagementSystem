package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql", "root:123456@/chao_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed to connect database: %v", err) // 使用日志记录错误，而不是panic
	}
	// 开db的debug模式
	db = db.Debug()
	// 实例一个接口
	userRepo := repository.NewUserRepository(db)
	ans1, _ := userRepo.HasUserByID(1)
	ans2, _ := userRepo.HasUserByID(1000)
	ans3, _ := userRepo.HasUserByTel("15126192562")
	ans4, _ := userRepo.HasUserByTel("1122545")
	if ans1 {
		fmt.Println("has user1")
	}
	if !ans2 {
		fmt.Println("no has")
	}
	if ans3 {
		fmt.Println("has tel ...")
	}
	if !ans4 {
		fmt.Println("no has")
	}

}
