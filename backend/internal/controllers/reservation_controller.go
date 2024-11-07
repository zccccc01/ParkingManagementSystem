package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

type ReservationController struct {
	ReservationRepo repository.ReservationRepository
}

func NewReservationController(repo repository.ReservationRepository) *ReservationController {
	return &ReservationController{ReservationRepo: repo}
}

// 创建预定记录
func (rc *ReservationController) CreateReservation(c *fiber.Ctx) error {
	var reservation models.Reservation
	if err := c.BodyParser(&reservation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	success, err := rc.ReservationRepo.Create(&reservation)
	if err != nil || !success {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Reservation created successfully"})
}

// 删除预定
func (rc *ReservationController) CancelReservation(c *fiber.Ctx) error {
	reservationId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid reservation ID"})
	}

	if err := rc.ReservationRepo.DeleteByReservationID(reservationId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Reservation cancelled successfully"})
}

// 更新预定状态
func (rc *ReservationController) UpdateReservationStatus(c *fiber.Ctx) error {
	reservationId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid reservation ID"})
	}

	var reservation models.Reservation
	if err := c.BodyParser(&reservation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := rc.ReservationRepo.UpdateStatusByReservationID(reservationId, reservation.Status); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Reservation status updated successfully"})
}

// 更新预定记录
func (rc *ReservationController) UpdateReservation(c *fiber.Ctx) error {
	reservationId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid reservation ID"})
	}

	var reservation models.Reservation
	if err := c.BodyParser(&reservation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := rc.ReservationRepo.UpdateByReservationID(reservationId, &reservation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Reservation updated successfully"})
}

func (rc *ReservationController) GetFeeByLotID(c *fiber.Ctx) error {
	lotId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid lot ID"})
	}

	startStr := c.Query("start")
	endStr := c.Query("end")
	if startStr == "" || endStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Start and end times are required"})
	}

	startTime, err := time.Parse(time.RFC3339, startStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start time format"})
	}
	endTime, err := time.Parse(time.RFC3339, endStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end time format"})
	}

	fee, err := rc.ReservationRepo.GetFeeByLotIDAndTime(lotId, startTime, endTime)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"fee": fee})
}
