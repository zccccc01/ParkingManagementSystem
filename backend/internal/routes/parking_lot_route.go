package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/controllers"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func SetupParkingLotRoutes(app *fiber.App, db *gorm.DB) {
	// 初始化 repository 和 service
	parkingLotRepo := repository.NewParkingLotRepository(db)
	parkingLotController := controllers.NewParkingLotController(parkingLotRepo)

	parkingLot := app.Group("/api/parkinglot")

	// 定义路由
	parkingLot.Post("/", parkingLotController.CreateParkingLot)
	parkingLot.Get("/:id", parkingLotController.GetParkingLotByID)
	parkingLot.Get("/", parkingLotController.GetAllParkingLots)
	parkingLot.Put("/:id", parkingLotController.UpdateParkingLot)
	// 这个有外键约束
	// parkingLot.Delete("/:id", parkingLotController.DeleteParkingLot)
}
