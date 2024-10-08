package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/controllers"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func SetupReservationRoutes(app *fiber.App, db *gorm.DB) {
	// 初始化 repository 和 service
	reservationRepo := repository.NewReservationRepository(db)
	reservationController := controllers.NewReservationController(reservationRepo)

	// 定义路由组
	reservation := app.Group("/api/reservation")

	// 定义路由
	reservation.Post("/", reservationController.CreateReservation)
	reservation.Put("/id/:id", reservationController.UpdateReservationStatus)
	reservation.Delete("/id/:id", reservationController.CancelReservation)
	// reservation.Get("/id/:id", reservationController.GetReservationByID) TODO:未完成
}
