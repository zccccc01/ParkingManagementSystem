package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shopspring/decimal"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql", "root:123456@/chao_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed to connect database: %v", err) // 使用日志记录错误，而不是panic
	}

	// 实例一个接口
	parkingLotRepo := repository.NewParkingLotRepository(db)
	// 把新的停车场数据写好
	lot := &models.ParkingLot{
		ParkingLotID: 1,
		ParkingName:  "hjh Parking",
		// Longitude:    decimal.RequireFromString("13"),
		// Latitude:     decimal.RequireFromString("65"),
		Capacity: 1,
		Rates:    decimal.RequireFromString("666"),
	}
	// 调用 Update 方法更新记录
	id := 1
	parkingLotRepo.Update(lot, id)
	res, err1 := parkingLotRepo.FindByID(id)
	if err1 != nil {
		log.Fatalf("failed to find parking lot: %v", err1)
	}
	log.Printf("Updated parking lot: %+v", res)
}
