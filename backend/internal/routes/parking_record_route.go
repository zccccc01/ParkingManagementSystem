package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/controllers"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func SetupParkingRecordRoutes(app *fiber.App, db *gorm.DB) {
	// 初始化 repository 和 controller
	parkingRecordRepo := repository.NewParkingRecordRepository(db)
	parkingRecordController := controllers.NewParkingRecordController(parkingRecordRepo)

	// 定义路由组
	parkingRecord := app.Group("/api/parkingrecord")

	// 定义路由
	parkingRecord.Post("/", parkingRecordController.CreateParkingRecord)
	parkingRecord.Get("/user/:id", parkingRecordController.GetParkingRecordByUserID)
	parkingRecord.Get("/vehicle/:id", parkingRecordController.GetParkingRecordFeeByVehicleID)
	parkingRecord.Get("/:id", parkingRecordController.GetParkingRecordFee)
	parkingRecord.Put("/:id", parkingRecordController.UpdateParkingRecord)

}
