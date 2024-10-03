package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/controllers"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/service"
)

// GET对应资源获取
// POST对应创造新的
// PUT对应更新资源
// DELETE对应删除

func SetupParkingLotRoutes(app *fiber.App, db *gorm.DB) {
	// 初始化 repository 和 service
	parkingLotRepo := repository.NewParkingLotRepository(db)
	parkingLotService := service.NewParkingLotService(parkingLotRepo)
	parkingLotController := controllers.NewParkingLotController(parkingLotService)

	// 定义路由
	app.Post("/api/parking_lots", parkingLotController.CreateParkingLot)
	app.Get("/api/parking_lots/:id", parkingLotController.GetParkingLotByID)
	app.Get("/api/parking_lots", parkingLotController.GetAllParkingLots)
	app.Put("/api/parking_lots/:id", parkingLotController.UpdateParkingLot)
	// 这个有外键约束
	// app.Delete("/api/parking_lots/:id", parkingLotController.DeleteParkingLot)
}
