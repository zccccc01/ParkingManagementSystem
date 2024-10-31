package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/controllers"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func SetupPaymentRecordRoutes(app *fiber.App, db *gorm.DB) {
	// 初始化 repository 和 controller
	paymentRepo := repository.NewPaymentRecordRepository(db)
	paymentController := controllers.NewPaymentRecordController(paymentRepo)

	// 定义路由组
	paymentRecord := app.Group("/api/paymentrecord")

	// 定义路由
	paymentRecord.Get("/reservation/:id", paymentController.GetFeeByReservationID)
	paymentRecord.Get("/status/reservation/:id", paymentController.GetPaymentStatusByReservationID)
	paymentRecord.Get("/record/:id", paymentController.GetFeeByRecordID)
	paymentRecord.Get("/status/record/:id", paymentController.GetPaymentStatusByRecordID)
	paymentRecord.Get("/plate/:plate", paymentController.GetFeeByPlate)
}
