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
	// 连接数据库
	db, err := gorm.Open("mysql", "root:123456@/chao_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed to connect database: %v", err) // 使用日志记录错误，而不是panic
	}

	// 实例一个接口
	vehicleRepo := repository.NewVehicleRepository(db)

	vehicle := models.Vehicle{
		//VehicleID:   11,
		//UserID:      2,
		PlateNumber: "12345",
		Color:       "yellow",
	}

	result := vehicleRepo.UpdateVehicleByVehicleID(11, &vehicle)
	if result != nil {
		fmt.Println(result)
	}
}
