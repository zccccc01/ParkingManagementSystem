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
	violationRecord.Get("/fineamount/recordid/:id", violationRecordController.GetFineAmountByRecordId)
	violationRecord.Get("/status/recordid/:id", violationRecordController.GetStatusByRecordId)
	violationRecord.Get("/type/recordid/:id", violationRecordController.GetViolationTypeByRecordId)
	violationRecord.Get("/record/userid/:id", violationRecordController.GetViolationRecordsByUserID)
}
