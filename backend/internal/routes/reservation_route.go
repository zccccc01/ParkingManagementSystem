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
	reservation.Put("/id/:id", reservationController.UpdateReservationStatus) //TODO:有问题 "error": "Error 1265: Data truncated for column 'Status' at row 1"
	reservation.Delete("/id/:id", reservationController.CancelReservation)    //TODO:应该是改status,post方法才对
	// reservation.Get("/id/:id", reservationController.GetReservationByID) TODO:未完成
}
