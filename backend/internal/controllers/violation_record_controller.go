package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

type ViolationRecordController struct {
	ViolationRecordRepo repository.ViolationRecordRepository
}

func NewViolationRecordController(repo repository.ViolationRecordRepository) *ViolationRecordController {
	return &ViolationRecordController{ViolationRecordRepo: repo}
}

// @Summary Create a new violation record
// @Description Create a new violation record
// @Tags ViolationRecord
// @Accept json
// @Produce json
// @Param violationRecord body models.ViolationRecord true "Violation record information"
// @Success 201 {object} models.ViolationRecord
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/violationrecord [post]
func (vrc *ViolationRecordController) CreateViolationRecord(c *fiber.Ctx) error {
	var violationRecord models.ViolationRecord
	if err := c.BodyParser(&violationRecord); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}

	success, err := vrc.ViolationRecordRepo.Create(&violationRecord)
	if err != nil || !success {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Violation record created successfully"})
}

// @Summary Get fine amount by record ID
// @Description Get fine amount by record ID
// @Tags ViolationRecord
// @Accept json
// @Produce json
// @Param id path int true "Record ID"
// @Success 200 {object} models.ViolationRecord
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/violationrecord/fineamount/record/{id} [get]
func (vrc *ViolationRecordController) GetFineAmountByRecordId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}

	fineAmount, err := vrc.ViolationRecordRepo.GetFineAmountByRecordID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"fineAmount": fineAmount})
}

// @Summary Get status by record ID
// @Description Get status by record ID
// @Tags ViolationRecord
// @Accept json
// @Produce json
// @Param id path int true "Record ID"
// @Success 200 {object} models.ViolationRecord
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/violationrecord/status/record/{id} [get]
func (vrc *ViolationRecordController) GetStatusByRecordId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}

	status, err := vrc.ViolationRecordRepo.GetStatusByRecordID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": status})
}

// @Summary Get violation type by record ID
// @Description Get violation type by record ID
// @Tags ViolationRecord
// @Accept json
// @Produce json
// @Param id path int true "Record ID"
// @Success 200 {object} models.ViolationRecord
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/violationrecord/type/record/{id} [get]
func (vrc *ViolationRecordController) GetViolationTypeByRecordId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}

	violationType, err := vrc.ViolationRecordRepo.GetViolationTypeByRecordID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"violationType": violationType})
}

// @Summary Get violation records by user ID
// @Description Get violation records by user ID
// @Tags ViolationRecord
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} models.ViolationRecord
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/violationrecord/user/{id} [get]
func (vrc *ViolationRecordController) GetViolationRecordsByUserID(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}

	violationRecords, err := vrc.ViolationRecordRepo.FindViolationRecordByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"violationRecords": violationRecords})
}

// @Summary Statistical violations by type
// @Description Statistical violations by type
// @Tags ViolationRecord
// @Accept json
// @Produce json
// @Param type path string true "Violation type"
// @Success 200 {array} models.ViolationRecord
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/violationrecord/violation/{type} [get]
func (vrc *ViolationRecordController) StatisticalViolationsByType(c *fiber.Ctx) error {
	t := c.Params("type")
	violationRecords, err := vrc.ViolationRecordRepo.StatisticalViolationsByType(t)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"violationRecords": violationRecords})
}
