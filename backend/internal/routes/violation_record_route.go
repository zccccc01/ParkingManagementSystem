package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/controllers"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

// SetupViolationRecordRoutes 设置违规记录相关的路由
// @Description Violation Record API routes
func SetupViolationRecordRoutes(app *fiber.App, db *gorm.DB) {
	// 初始化 repository 和 controller
	violationRecordRepo := repository.NewViolationRecordRepository(db)
	violationRecordController := controllers.NewViolationRecordController(violationRecordRepo)

	// 定义路由组
	violationRecord := app.Group("/api/violationrecord")

	// 定义路由
	violationRecord.Post("/", violationRecordController.CreateViolationRecord)
	violationRecord.Get("/fineamount/record/:id", violationRecordController.GetFineAmountByRecordId)
	violationRecord.Get("/status/record/:id", violationRecordController.GetStatusByRecordId)
	violationRecord.Get("/type/record/:id", violationRecordController.GetViolationTypeByRecordId)
	violationRecord.Get("/user/:id", violationRecordController.GetViolationRecordsByUserID)
	violationRecord.Get("/violation/:type", violationRecordController.StatisticalViolationsByType)
}
