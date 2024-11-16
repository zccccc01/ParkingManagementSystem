package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

type VehicleController struct {
	VehicleRepo repository.VehicleRepository
}

func NewVehicleController(repo repository.VehicleRepository) *VehicleController {
	return &VehicleController{VehicleRepo: repo}
}

// @Summary Create a new vehicle
// @Description Create a new vehicle
// @Tags Vehicle
// @Accept json
// @Produce json
// @Param vehicle body models.Vehicle true "Vehicle information"
// @Success 201 {object} models.Vehicle
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/vehicle [post]
func (vc *VehicleController) CreateVehicle(c *fiber.Ctx) error {
	var vehicle models.Vehicle
	if err := c.BodyParser(&vehicle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	success, err := vc.VehicleRepo.Create(&vehicle)
	if err != nil || !success {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(vehicle)
}

// @Summary Get vehicle by ID
// @Description Get vehicle by ID
// @Tags Vehicle
// @Accept json
// @Produce json
// @Param id path int true "Vehicle ID"
// @Success 200 {object} models.Vehicle
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/vehicle/{id} [get]
func (vc *VehicleController) GetByVehicleID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	vehicle, err := vc.VehicleRepo.GetAllByVehicleID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Vehicle not found"})
	}

	return c.Status(fiber.StatusOK).JSON(vehicle)
}

// @Summary Get all vehicles by user ID
// @Description Get all vehicles by user ID
// @Tags Vehicle
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} models.Vehicle
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/vehicle/user/{id} [get]
func (vc *VehicleController) GetByUserID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	vehicle, err := vc.VehicleRepo.GetAllByUserID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Vehicle not found"})
	}

	return c.Status(fiber.StatusOK).JSON(vehicle)
}

// @Summary Update vehicle by ID
// @Description Update vehicle by ID
// @Tags Vehicle
// @Accept json
// @Produce json
// @Param id path int true "Vehicle ID"
// @Param vehicle body models.Vehicle true "Vehicle information"
// @Success 200 {object} models.Vehicle
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/vehicle/{id} [put]
func (vc *VehicleController) UpdateVehicle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var vehicle models.Vehicle
	if err := c.BodyParser(&vehicle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := vc.VehicleRepo.UpdateVehicleByVehicleID(id, &vehicle); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Vehicle updated successfully"})
}

// @Summary Delete vehicle by ID
// @Description Delete vehicle by ID
// @Tags Vehicle
// @Accept json
// @Produce json
// @Param id path int true "Vehicle ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/vehicle/{id} [delete]
func (vc *VehicleController) DeleteVehicle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := vc.VehicleRepo.DeleteByVehicleID(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Vehicle deleted successfully"})
}
