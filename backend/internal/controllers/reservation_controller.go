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

// @Summary Create a reservation
// @Description Create a new reservation
// @Tags Reservation
// @Accept json
// @Produce json
// @Param reservation body models.Reservation true "Reservation Details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/reservation [post]
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

// @Summary Cancel a reservation
// @Description Cancel an existing reservation
// @Tags Reservation
// @Accept json
// @Produce json
// @Param id path int true "Reservation ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/reservation/{id} [delete]
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

// @Summary Update reservation status
// @Description Update the status of an existing reservation
// @Tags Reservation
// @Accept json
// @Produce json
// @Param id path int true "Reservation ID"
// @Param reservation body models.Reservation true "Updated Reservation Details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/reservation/{id}/status [put]
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

// @Summary Update reservation details
// @Description Update the details of an existing reservation
// @Tags Reservation
// @Accept json
// @Produce json
// @Param id path int true "Reservation ID"
// @Param reservation body models.Reservation true "Updated Reservation Details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/reservation/{id} [put]
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

// @Summary Get fee for a reservation
// @Description Get the fee for a reservation based on lot ID and time range
// @Tags Reservation
// @Accept json
// @Produce json
// @Param id path int true "Lot ID"
// @Param start query string true "Start Time"
// @Param end query string true "End Time"
// @Success 200 {object} float64
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/reservation/fee/{id} [get]
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
