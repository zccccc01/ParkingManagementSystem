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

// 新建一个车的记录
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

// 根据车辆id获取一条记录
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

// 根据用户id获取所有记录
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

// 根据车辆id更新车牌号和颜色
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

// 根据车辆id删除一条记录
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
