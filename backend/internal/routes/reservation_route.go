package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/controllers"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func SetupReservationRoutes(app *fiber.App, db *gorm.DB) {
	// 初始化 repository 和 controller
	reservationRepo := repository.NewReservationRepository(db)
	reservationController := controllers.NewReservationController(reservationRepo)

	// 定义路由组
	reservation := app.Group("/api/reservation")

	// 定义路由
	reservation.Get("/lot/:id", reservationController.GetFeeByLotID)
	reservation.Post("/", reservationController.CreateReservation)
	reservation.Put("/id/:id", reservationController.UpdateReservation)
	reservation.Delete("/id/:id", reservationController.CancelReservation)
}
