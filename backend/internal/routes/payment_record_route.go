package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/controllers"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func SetupPaymentRecordRoutes(app *fiber.App, db *gorm.DB) {
	// 初始化 repository 和 service
	paymentRepo := repository.NewParkingSpaceRepository(db)
	parkingSpaceController := controllers.NewParkingSpaceController(paymentRepo)

	// 定义路由组
	paymentRecord := app.Group("/api/paymentrecord")

	// 定义路由
	paymentRecord.Get("/vehicle/:plateNumber", parkingSpaceController.GetParkingSpaceByLicensePlate)
	paymentRecord.Get("/user/:id", parkingSpaceController.GetParkingSpaceByUserID)
	paymentRecord.Get("/lot/:id", parkingSpaceController.GetParkingSpaceByParkingLotId)
	paymentRecord.Get("/status/lot/:lotid/space/:spaceid", parkingSpaceController.GetParkingSpaceStatusById) //TODO:传两个参数
	paymentRecord.Post("/", parkingSpaceController.CreateParkingSpace)
	paymentRecord.Put("/status/lot/:lotid/space/:spaceid", parkingSpaceController.UpdateParkingSpaceStatus) //TODO:传两个参数
}
