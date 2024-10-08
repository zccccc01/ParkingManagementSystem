package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/controllers"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func SetupVehicleRoutes(app *fiber.App, db *gorm.DB) {
	// 初始化 repository 和 controller
	vehicleRepo := repository.NewVehicleRepository(db)
	vehicleController := controllers.NewVehicleController(vehicleRepo)

	// 定义路由组
	vehicle := app.Group("/api/vehicle")

	// 定义路由
	vehicle.Post("/", vehicleController.CreateVehicle)
	vehicle.Get("/:id", vehicleController.GetByVehicleID)
	vehicle.Get("/user/:id", vehicleController.GetByUserID)
	vehicle.Put("/:id", vehicleController.UpdateVehicle) //TODO:数据库update有错
	vehicle.Delete("/:id", vehicleController.DeleteVehicle)
}
