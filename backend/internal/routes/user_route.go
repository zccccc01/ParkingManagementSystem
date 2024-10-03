package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/controllers"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func SetupUserRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userController := controllers.NewUserController(userRepo)

	// 当你使用app.Group("/api/user")时,所有路由都会以/api/user开头
	user := app.Group("/api/user")
	user.Post("/register", userController.Register)
	user.Post("/login", userController.Login)
	user.Get("/user", userController.AuthenticatedUser)
	user.Post("/logout", userController.Logout)
}
