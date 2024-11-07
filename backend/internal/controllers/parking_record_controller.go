package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

type ParkingRecordController struct {
	ParkingRecordRepo repository.ParkingRecordRepository
}

func NewParkingRecordController(repo repository.ParkingRecordRepository) *ParkingRecordController {
	return &ParkingRecordController{ParkingRecordRepo: repo}
}

// 创建停车记录
func (prc *ParkingRecordController) CreateParkingRecord(c *fiber.Ctx) error {
	var parkingRecord models.ParkingRecord
	if err := c.BodyParser(&parkingRecord); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err := prc.ParkingRecordRepo.CreateRecordEntry(&parkingRecord)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Parking record created successfully"})
}

// 根据记录ID获取费用
func (prc *ParkingRecordController) GetParkingRecordFee(c *fiber.Ctx) error {
	recordID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid record ID"})
	}

	fee, err := prc.ParkingRecordRepo.GetFeeByRecordID(recordID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Record not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"fee": fee})
}

// 根据车辆ID获取费用
func (prc *ParkingRecordController) GetParkingRecordFeeByVehicleID(c *fiber.Ctx) error {
	vehicleID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid vehicle ID"})
	}

	fee, err := prc.ParkingRecordRepo.GetFeeByVehicleID(vehicleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Record not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"fee": fee})
}

// 根据UserID查历史记录
func (prc *ParkingRecordController) GetParkingRecordByUserID(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	records, err := prc.ParkingRecordRepo.FindHistoryRecordByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"records": records})
}

// 根据ID更新出场记录
func (prc *ParkingRecordController) UpdateParkingRecord(c *fiber.Ctx) error {
	recordID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid record ID"})
	}

	var parkingRecord models.ParkingRecord
	if err := c.BodyParser(&parkingRecord); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := prc.ParkingRecordRepo.UpdateRecordExitByRecordID(recordID, time.Now()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Parking record updated successfully"})
}

// 根据年月获取月度报告
func (prc *ParkingRecordController) GetMonthlyReport(c *fiber.Ctx) error {
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid year"})
	}
	month, err := strconv.Atoi(c.Query("month"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid month"})
	}

	records, err := prc.ParkingRecordRepo.GetMonthlyReport(year, month)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"records": records})
}

// 根据年获取年度报告
func (prc *ParkingRecordController) GetAnnualReport(c *fiber.Ctx) error {
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid year"})
	}

	records, err := prc.ParkingRecordRepo.GetAnnualReport(year)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"records": records})
}
