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
	// ans, err := userRepo.UpdateUserName(1, "123456789", "CLzz")
	// fmt.Println(ans, err)
	ans1, err1 := userRepo.FindUserByTel("123456789")
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(ans1)
	ans2, err2 := userRepo.FindUserByTel("1")
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(ans2)
}
