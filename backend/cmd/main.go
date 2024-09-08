package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@/chao_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed to connect database: %v", err) // 使用日志记录错误，而不是panic
	}

	// 自动迁移
	db.AutoMigrate(&models.ParkingLot{}, &models.ParkingSpace{})

	// 创建仓库实例
	lotRepo := repository.NewParkingLotRepository(db)

	newLot := models.ParkingLot{ParkingName: "Main Parking Lot"}
	if err := lotRepo.Create(&newLot); err != nil {
		log.Printf("failed to create parking lot: %v", err)
	}
	fmt.Println("Parking lot and space created successfully.")
}
