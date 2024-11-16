package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

type ParkingSpaceController struct {
	ParkingSpaceRepo repository.ParkingSpaceRepository
}

func NewParkingSpaceController(repo repository.ParkingSpaceRepository) *ParkingSpaceController {
	return &ParkingSpaceController{ParkingSpaceRepo: repo}
}

// @Summary Create a parking space
// @Description Create a new parking space
// @Tags ParkingSpace
// @Accept json
// @Produce json
// @Param parkingSpace body models.ParkingSpace true "Parking Space Details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/parkingspace [post]
func (psc *ParkingSpaceController) CreateParkingSpace(c *fiber.Ctx) error {
	var parkingSpace models.ParkingSpace
	if err := c.BodyParser(&parkingSpace); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	success, err := psc.ParkingSpaceRepo.Create(&parkingSpace)
	if err != nil || !success {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Parking space created successfully"})
}

// @Summary Get parking spaces by parking lot id
// @Description Retrieve the parking spaces by parking lot id
// @Tags ParkingSpace
// @Accept json
// @Produce json
// @Param id path int true "Parking Lot ID"
// @Success 200 {array} models.ParkingSpace
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/parkingspace/{id} [get]
func (psc *ParkingSpaceController) GetParkingSpaceByParkingLotId(c *fiber.Ctx) error {
	parkingLotId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid parking lot id"})
	}

	parkingSpaces, err := psc.ParkingSpaceRepo.GetAllStatusByLotID(parkingLotId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Space not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"spaces": parkingSpaces})
}

// @Summary Update parking space status
// @Description Update the status of a parking space
// @Tags ParkingSpace
// @Accept json
// @Produce json
// @Param lotid path int true "Parking Lot ID"
// @Param spaceid path int true "Parking Space ID"
// @Param parkingSpace body models.ParkingSpace true "Updated Parking Space Details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/parkingspace/{lotid}/{spaceid} [put]
func (psc *ParkingSpaceController) UpdateParkingSpaceStatus(c *fiber.Ctx) error {
	parkingLotID, err := strconv.Atoi(c.Params("lotid"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid parking lot id"})
	}
	parkingSpaceId, err := strconv.Atoi(c.Params("spaceid"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid parking space id"})
	}

	var parkingSpace models.ParkingSpace
	if err := c.BodyParser(&parkingSpace); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	success, err := psc.ParkingSpaceRepo.UpdateStatusBySpaceID(&parkingSpace, parkingSpaceId, parkingLotID)
	if err != nil || !success {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Parking space updated successfully"})
}

// @Summary Get parking space by license plate
// @Description Retrieve the parking space by license plate
// @Tags ParkingSpace
// @Accept json
// @Produce json
// @Param plateNumber path string true "License Plate Number"
// @Success 200 {array} models.ParkingSpace
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/parkingspace/plate/{plateNumber} [get]
func (psc *ParkingSpaceController) GetParkingSpaceByLicensePlate(c *fiber.Ctx) error {
	plateNumber := c.Params("plateNumber") //车牌号是通过 URL 参数传递的

	spaces, err := psc.ParkingSpaceRepo.FindVehicleSpaceInLotByPlateNumber(plateNumber)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"spaces": spaces})
}

// @Summary Get parking space by user id
// @Description Retrieve the parking space by user id
// @Tags ParkingSpace
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} models.ParkingSpace
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/parkingspace/user/{id} [get]
func (psc *ParkingSpaceController) GetParkingSpaceByUserID(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user id"})
	}

	spaces, err := psc.ParkingSpaceRepo.FindVehicleSpaceInLotByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"spaces": spaces})
}

// @Summary Get parking space status by id
// @Description Retrieve the status of a parking space by id
// @Tags ParkingSpace
// @Accept json
// @Produce json
// @Param lotid query int true "Parking Lot ID"
// @Param spaceid query int true "Parking Space ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/parkingspace/status [get]
func (psc *ParkingSpaceController) GetParkingSpaceStatusById(c *fiber.Ctx) error {
	parkingLotID, err := strconv.Atoi(c.Query("lotid"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid parking lot id"})
	}
	parkingSpaceId, err := strconv.Atoi(c.Query("spaceid"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid parking space id"})
	}

	status, err := psc.ParkingSpaceRepo.GetStatusByLotIDAndSpaceID(parkingLotID, parkingSpaceId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Space not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": status})
}

// @Summary Get free parking spaces
// @Description Retrieve the free parking spaces
// @Tags ParkingSpace
// @Accept json
// @Produce json
// @Success 200 {array} models.ParkingSpace
// @Failure 500 {object} map[string]string
// @Router /api/parkingspace/free [get]
func (psc *ParkingSpaceController) GetFreeParkingSpace(c *fiber.Ctx) error {
	spaces, err := psc.ParkingSpaceRepo.FindFreeSpaceInAllLots()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"spaces": spaces})
}
