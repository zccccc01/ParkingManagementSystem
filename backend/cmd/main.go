package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shopspring/decimal"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@/chao_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed to connect database: %v", err) // 使用日志记录错误，而不是panic
	}

	// 创建记录
	newLot := models.ParkingLot{
		ParkingLotID: 1,
		ParkingName:  "Central Parking",
		Longitude:    decimal.RequireFromString("115.000000"),
		Latitude:     decimal.RequireFromString("39.000000"),
		Capacity:     100,
		Rates:        decimal.RequireFromString("2.50"),
	}
	result := db.Create(&newLot)
	if result.Error != nil {
		log.Fatalf("failed to create record: %v", result.Error)
	}

	// 打印创建的记录
	log.Printf("Created parking lot: %+v", newLot)

}
