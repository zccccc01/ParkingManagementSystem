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

	// 实例一个接口
	violationRepo := repository.NewViolationRecordRepository(db)

	//TODO: violation_record_repository_impl.go 未review
	// violation := models.ViolationRecord{
	// 	ViolationID:   1,
	// 	RecordID:      123,
	// 	FineAmount:    30,
	// 	ViolationType: "NOPAY",
	// 	Status:        "UNPAID",
	// }
	// result := violationRepo.Create(&violation)
	// if result != nil {
	// 	fmt.Println(result)
	// }

	fine, res := violationRepo.GetViolationTypeByRecordID(123)
	if res != nil {
		fmt.Println(res)
	}
	fmt.Println(fine)

	//user
	//userRepo := repository.NewUserRepository(db)
	// newUser := &models.User{
	// 	Username: "john_doe",
	// 	Password: "securepassword123",
	// 	Tel:      "1234567890",
	// }
	// err = userRepo.Create(newUser)
	// if err != nil {
	// 	log.Fatal("failed to create user:", err)
	// }

	// // 更新密码
	// err = userRepo.UpdatePasswordByID(newUser.UserID, "newsecurepassword456")
	// if err != nil {
	// 	log.Fatal("failed to update user password:", err)
	// }

	// // 获取电话
	// tel, err := userRepo.GetTelByID(newUser.UserID)
	// if err != nil {
	// 	log.Fatal("failed to get user tel:", err)
	// }
	// log.Println("User Tel:", tel)

	// // 删除记录
	// err = userRepo.Delete(newUser.UserID)
	// if err != nil {
	// 	log.Fatal("failed to delete user:", err)
	// }
}
