package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/controllers"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func SetupParkingSpaceRoutes(app *fiber.App, db *gorm.DB) {
	// 初始化 repository 和 controller
	parkingSpaceRepo := repository.NewParkingSpaceRepository(db)
	parkingSpaceController := controllers.NewParkingSpaceController(parkingSpaceRepo)

	// 定义路由组
	parkingSpace := app.Group("/api/parkingspace")

	// 定义路由
	parkingSpace.Get("/vehicle/:plateNumber", parkingSpaceController.GetParkingSpaceByLicensePlate)
	parkingSpace.Get("/user/:id", parkingSpaceController.GetParkingSpaceByUserID)
	parkingSpace.Get("/lot/:id", parkingSpaceController.GetParkingSpaceByParkingLotId)
	parkingSpace.Get("/status/lot/:lotid/space/:spaceid", parkingSpaceController.GetParkingSpaceStatusById) //TODO:传两个参数
	parkingSpace.Get("/status/free", parkingSpaceController.GetFreeParkingSpace)
	parkingSpace.Post("/", parkingSpaceController.CreateParkingSpace)
	parkingSpace.Put("/status/lot/:lotid/space/:spaceid", parkingSpaceController.UpdateParkingSpaceStatus)
}
