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

// 创建一条记录
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

// 根据记录id获取罚款金额
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

// 根据记录id获取状态
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

// 根据记录id获取违章类型
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

// 根据UserID查违章记录
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

func (vrc *ViolationRecordController) StatisticalViolationsByType(c *fiber.Ctx) error {
	t := c.Params("type")
	violationRecords, err := vrc.ViolationRecordRepo.StatisticalViolationsByType(t)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"violationRecords": violationRecords})
}
