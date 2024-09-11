package main

import (
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
	violationRecordRepo := repository.NewViolationRecordRepository(db)

	// amounts, ans := violationRecordRepo.GetFineAmountByRecordID(123)
	// if ans != nil {
	// 	log.Fatalf("failed to get fine amount: %v", err) // 使用日志记录错误，而不是panic
	// }
	// log.Printf("Fine amount: %v", amounts)
	// status, ans := violationRecordRepo.GetStatusByRecordID(123)
	// if ans != nil {
	// 	log.Fatalf("failed to get fine amount: %v", err) // 使用日志记录错误，而不是panic
	// }
	// log.Printf("Fine amount: %v", status)
	type1, ans := violationRecordRepo.GetViolationTypeByRecordID(123)
	if ans != nil {
		log.Fatalf("failed to get fine amount: %v", err) // 使用日志记录错误，而不是panic
	}
	log.Printf("Fine amount: %v", type1)
}
