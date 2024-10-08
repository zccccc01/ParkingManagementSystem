package controllers

import (
	"strconv"

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
	status := c.Params("status")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid reservation ID"})
	}

	if err := rc.ReservationRepo.UpdateStatusByReservationID(reservationId, status); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Reservation status updated successfully"})
}
