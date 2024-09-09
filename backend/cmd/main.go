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
	parkingRecordRepo := repository.NewParkingRecordRepository(db)

	// record := models.ParkingRecord{
	// 	RecordID:  123,
	// 	VehicleID: 11,
	// 	SpaceID:   1,
	// 	LotID:     1,
	// 	StartTime: time.Time{},
	// 	EndTime:   time.Time{},
	// }

	// ans := parkingRecordRepo.CreateRecordEntry(&record)
	// if ans != nil {
	// 	log.Fatalf("failed to create record: %v", ans)
	// }

	// ans2 := parkingRecordRepo.UpdateRecordExitByRecordID(123, time.Now())
	// if ans2 != nil {
	// 	log.Fatalf("failed to create record: %v", ans2)
	// }
	// ans, err := parkingRecordRepo.GetFeeByRecordID(123)
	// if err != nil {
	// 	log.Fatalf("failed to get fee: %v", err)
	// }
	// fmt.Println(ans)

	ans2, err2 := parkingRecordRepo.GetFeeByVehicleID(11)
	if err2 != nil {
		log.Fatalf("failed to get fee: %v", err2)
	}
	fmt.Println(ans2)
}
