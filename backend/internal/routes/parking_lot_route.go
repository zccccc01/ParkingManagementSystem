package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/controllers"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

// SetupParkingLotRoutes 设置停车场相关路由
// @Description Parking Lot API routes
func SetupParkingLotRoutes(app *fiber.App, db *gorm.DB) {
	// 初始化 repository 和 controller
	parkingLotRepo := repository.NewParkingLotRepository(db)
	parkingLotController := controllers.NewParkingLotController(parkingLotRepo)

	// 定义路由组
	parkingLot := app.Group("/api/parkinglot")

	// 定义路由
	parkingLot.Post("/", parkingLotController.CreateParkingLot)
	parkingLot.Get("/id/:id", parkingLotController.GetParkingLotByID)
	parkingLot.Get("/name/:name", parkingLotController.GetParkingLotsByName)
	parkingLot.Get("/", parkingLotController.GetAllParkingLots)
	parkingLot.Get("/income/:id", parkingLotController.GetAllIncomeByID)
	parkingLot.Get("/occupancy-rate/:id", parkingLotController.GetOccupancyRateByID)
	parkingLot.Get("/id/:id/start/:start/end/:end", parkingLotController.GetOccupancyByIDAndTime)
	parkingLot.Put("/id/:id", parkingLotController.UpdateParkingLot)
	parkingLot.Get("/status/lot/:id", parkingLotController.GetStatusByID)
	parkingLot.Get("/allincome/all", parkingLotController.GetAllIncomeByLotID)
	// 这个有外键约束
	// parkingLot.Delete("/:id", parkingLotController.DeleteParkingLot)
}
