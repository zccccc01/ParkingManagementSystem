package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/controllers"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func SetupVehicleRoutes(app *fiber.App, db *gorm.DB) {
	// 初始化 repository 和 service
	vehicleRepo := repository.NewVehicleRepository(db)
	vehicleController := controllers.NewVehicleController(vehicleRepo)

	// 定义路由组
	vehicle := app.Group("/api/vehicle")

	// 定义路由
	vehicle.Post("/", vehicleController.CreateVehicle)
	vehicle.Get("/vehicleid/:id", vehicleController.GetByVehicleID)
	vehicle.Get("/userid/:id", vehicleController.GetByUserID)
	vehicle.Put("/vehicleid/:id", vehicleController.UpdateVehicle)
	vehicle.Delete("/vehicleid/:id", vehicleController.DeleteVehicle)
}
