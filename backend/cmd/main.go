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
	db = db.Debug()
	// 实例一个接口
	parkingLotRepo := repository.NewParkingLotRepository(db)

	// TODO: where ="unpaid" 找到RecordID 从RID找VehicleID 找人的ID
	res, err := parkingLotRepo.FindAllIncomeByLotID(1)
	if err != nil {
		log.Fatalf("failed to find all income by lot id: %v", err)

	}
	log.Println(res)
}
