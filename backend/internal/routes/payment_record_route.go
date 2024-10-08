package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/controllers"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func SetupPaymentRecordRoutes(app *fiber.App, db *gorm.DB) {
	// 初始化 repository 和 controller
	paymentRepo := repository.NewParkingSpaceRepository(db)
	paymentController := controllers.NewParkingSpaceController(paymentRepo)

	// 定义路由组
	paymentRecord := app.Group("/api/paymentrecord")

	// 定义路由,TODO:前三个有点怪
	paymentRecord.Get("/vehicle/:plateNumber", paymentController.GetParkingSpaceByLicensePlate)
	paymentRecord.Get("/user/:id", paymentController.GetParkingSpaceByUserID)
	paymentRecord.Get("/lot/:id", paymentController.GetParkingSpaceByParkingLotId)
	paymentRecord.Get("/status/lot/:lotid/space/:spaceid", paymentController.GetParkingSpaceStatusById) //TODO:传两个参数
	paymentRecord.Post("/", paymentController.CreateParkingSpace)
	paymentRecord.Put("/status/lot/:lotid/space/:spaceid", paymentController.UpdateParkingSpaceStatus) //TODO:传两个参数
}
