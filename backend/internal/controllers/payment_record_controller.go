package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

type PaymentRecordController struct {
	PaymentRecordRepo repository.PaymentRecordRepository
}

func NewPaymentRecordController(repo repository.PaymentRecordRepository) *PaymentRecordController {
	return &PaymentRecordController{PaymentRecordRepo: repo}
}

// 获取费用通过预定ID
func (prc *PaymentRecordController) GetFeeByReservationID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	amount, err := prc.PaymentRecordRepo.GetAmountByReservationID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
	}

	return c.Status(fiber.StatusOK).JSON(amount)
}

// 获取支付状态通过预定ID
func (prc *PaymentRecordController) GetPaymentStatusByReservationID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	status, err := prc.PaymentRecordRepo.GetPaymentStatusByReservationID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Reservation not found"})
	}

	return c.Status(fiber.StatusOK).JSON(status)
}

// 获取费用通过记录ID
func (prc *PaymentRecordController) GetFeeByRecordID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	amount, err := prc.PaymentRecordRepo.GetAmountByRecordID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
	}

	return c.Status(fiber.StatusOK).JSON(amount)
}

// 获取支付状态通过记录ID
func (prc *PaymentRecordController) GetPaymentStatusByRecordID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	status, err := prc.PaymentRecordRepo.GetPaymentStatusByRecordID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
	}

	return c.Status(fiber.StatusOK).JSON(status)
}

// 获取费用通过车牌号
func (prc *PaymentRecordController) GetFeeByPlate(c *fiber.Ctx) error {
	plate := c.Params("plate")

	amount, err := prc.PaymentRecordRepo.GetPaymentFeeByPlateNumber(plate)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
	}

	return c.Status(fiber.StatusOK).JSON(amount)
}

// 获取支付信息通过车牌号
func (prc *PaymentRecordController) GetPaymentInfoByPlate(c *fiber.Ctx) error {
	plate := c.Params("plate")

	paymentInfo, err := prc.PaymentRecordRepo.GetPaymentInfoByPlateNumber(plate)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
	}

	return c.Status(fiber.StatusOK).JSON(paymentInfo)
}
