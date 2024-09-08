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

	// // 实例一个接口
	parkingLotRepo := repository.NewParkingLotRepository(db)
	newLot := models.ParkingLot{
		ParkingLotID: 2,
		ParkingName:  "CMU Parking",
		Longitude:    decimal.RequireFromString("121.00"),
		Latitude:     decimal.RequireFromString("5.000"),
		Capacity:     100,
		Rates:        decimal.RequireFromString("10"),
	}
	result := parkingLotRepo.Create(&newLot)
	// 用result判断是否创建成功
	if result != nil {
		log.Fatalf("failed to create record: %v", result)
	}
	// 打印创建的记录
	log.Printf("Created parking lot: %+v", newLot)

}
