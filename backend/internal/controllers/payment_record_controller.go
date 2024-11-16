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

// @Summary Get payment amount by reservation ID
// @Description Retrieve the payment amount by reservation ID
// @Tags PaymentRecord
// @Accept json
// @Produce json
// @Param id path int true "Reservation ID"
// @Success 200 {object} float64
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/paymentrecord/reservation/{id} [get]
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

// @Summary Get payment status by reservation ID
// @Description Retrieve the payment status by reservation ID
// @Tags PaymentRecord
// @Accept json
// @Produce json
// @Param id path int true "Reservation ID"
// @Success 200 {object} bool
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/paymentrecord/reservation/{id}/status [get]
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

// @Summary Get payment amount by record ID
// @Description Retrieve the payment amount by record ID
// @Tags PaymentRecord
// @Accept json
// @Produce json
// @Param id path int true "Record ID"
// @Success 200 {object} float64
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/paymentrecord/record/{id} [get]
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

// @Summary Get payment status by record ID
// @Description Retrieve the payment status by record ID
// @Tags PaymentRecord
// @Accept json
// @Produce json
// @Param id path int true "Record ID"
// @Success 200 {object} bool
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/paymentrecord/record/{id}/status [get]
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

// @Summary Get payment amount by plate number
// @Description Retrieve the payment amount by plate number
// @Tags PaymentRecord
// @Accept json
// @Produce json
// @Param plate path string true "Plate Number"
// @Success 200 {object} float64
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/paymentrecord/plate/{plate} [get]
func (prc *PaymentRecordController) GetFeeByPlate(c *fiber.Ctx) error {
	plate := c.Params("plate")

	amount, err := prc.PaymentRecordRepo.GetPaymentFeeByPlateNumber(plate)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
	}

	return c.Status(fiber.StatusOK).JSON(amount)
}

// @Summary Get payment info by plate number
// @Description Retrieve the payment info by plate number
// @Tags PaymentRecord
// @Accept json
// @Produce json
// @Param plate path string true "Plate Number"
// @Success 200 {object} models.PaymentRecord
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/paymentrecord/plate/{plate}/info [get]
func (prc *PaymentRecordController) GetPaymentInfoByPlate(c *fiber.Ctx) error {
	plate := c.Params("plate")

	paymentInfo, err := prc.PaymentRecordRepo.GetPaymentInfoByPlateNumber(plate)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
	}

	return c.Status(fiber.StatusOK).JSON(paymentInfo)
}
