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

// @Summary Create a parking record
// @Description Create a new parking record
// @Tags ParkingRecord
// @Accept json
// @Produce json
// @Param parkingRecord body models.ParkingRecord true "Parking Record Details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/parkingrecord [post]
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

// @Summary Get parking record fee by record ID
// @Description Retrieve the fee of a parking record by its ID
// @Tags ParkingRecord
// @Accept json
// @Produce json
// @Param id path int true "Parking Record ID"
// @Success 200 {object} float64
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/parkingrecord/{id} [get]
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

// @Summary Get parking record fee by vehicle ID
// @Description Retrieve the fee of a parking record by its vehicle ID
// @Tags ParkingRecord
// @Accept json
// @Produce json
// @Param id path int true "Vehicle ID"
// @Success 200 {object} float64
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/parkingrecord/vehicle/{id} [get]
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

// @Summary Get parking record history by user ID
// @Description Retrieve the history records of a user
// @Tags ParkingRecord
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} models.ParkingRecord
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/parkingrecord/user/{id} [get]
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

// @Summary Update parking record exit time
// @Description Update the exit time of a parking record
// @Tags ParkingRecord
// @Accept json
// @Produce json
// @Param id path int true "Parking Record ID"
// @Param parkingRecord body models.ParkingRecord true "Updated Parking Record Details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/parkingrecord/{id} [put]
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

// @Summary Get monthly report
// @Description Retrieve the monthly report of parking records
// @Tags ParkingRecord
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Success 200 {array} models.ParkingRecord
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/parkingrecord/month [get]
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

// @Summary Get annual report
// @Description Retrieve the annual report of parking records
// @Tags ParkingRecord
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {array} models.ParkingRecord
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/parkingrecord/year [get]
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
