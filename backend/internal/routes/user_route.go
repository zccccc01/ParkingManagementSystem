package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/controllers"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func SetupUserRoutes(app *fiber.App, db *gorm.DB) {
	// 初始化 repository 和 controller
	userRepo := repository.NewUserRepository(db)
	userController := controllers.NewUserController(userRepo)

	// 定义路由组
	user := app.Group("/api/user")

	// 定义路由
	user.Post("/register", userController.Register)
	user.Post("/login", userController.Login)
	user.Get("/", userController.AuthenticatedUser)
	user.Post("/logout", userController.Logout)
}
