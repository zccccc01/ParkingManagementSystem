package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/controllers"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func SetupViolationRecordRoutes(app *fiber.App, db *gorm.DB) {
	// 初始化 repository 和 service
	violationRecordRepo := repository.NewViolationRecordRepository(db)
	violationRecordController := controllers.NewViolationRecordController(violationRecordRepo)

	// 定义路由组
	violationRecord := app.Group("/api/violationrecord")

	// 定义路由
	violationRecord.Post("/", violationRecordController.CreateViolationRecord)
	violationRecord.Get("/fineamount/record/:id", violationRecordController.GetFineAmountByRecordId)
	violationRecord.Get("/status/record/:id", violationRecordController.GetStatusByRecordId)
	violationRecord.Get("/type/record/:id", violationRecordController.GetViolationTypeByRecordId)
	violationRecord.Get("/user/:id", violationRecordController.GetViolationRecordsByUserID) // TODO:未完成
}
