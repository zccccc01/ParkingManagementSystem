package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/service"
)

type ParkingLotController struct {
	service *service.ParkingLotService
}

func NewParkingLotController(service *service.ParkingLotService) *ParkingLotController {
	return &ParkingLotController{service: service}
}

// CreateParkingLot 创建停车场
func (c *ParkingLotController) CreateParkingLot(ctx *fiber.Ctx) error {
	var lot models.ParkingLot
	if err := ctx.BodyParser(&lot); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	success, err := c.service.CreateParkingLot(&lot)
	if err != nil || !success {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Parking lot created successfully"})
}

// GetParkingLotByID 根据ID获取停车场
func (c *ParkingLotController) GetParkingLotByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	lot, err := c.service.GetParkingLotByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Parking lot not found"})
	}

	return ctx.Status(fiber.StatusOK).JSON(lot)
}

// GetAllParkingLots 获取所有停车场
func (c *ParkingLotController) GetAllParkingLots(ctx *fiber.Ctx) error {
	lots, err := c.service.GetAllParkingLots()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(lots)
}

// UpdateParkingLot 更新停车场信息
func (c *ParkingLotController) UpdateParkingLot(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var lot models.ParkingLot
	if err := ctx.BodyParser(&lot); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.service.UpdateParkingLot(&lot, id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Parking lot updated successfully"})
}

// DeleteParkingLot 删除停车场
func (c *ParkingLotController) DeleteParkingLot(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := c.service.DeleteParkingLot(id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Parking lot deleted successfully"})
}
