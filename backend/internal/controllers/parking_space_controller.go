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

// 创建停车位
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

// 根据停车场id获取该停车场车位空余情况
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

// 根据车位id更新状态
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

// 根据车牌号查看停车位置
func (psc *ParkingSpaceController) GetParkingSpaceByLicensePlate(c *fiber.Ctx) error {
	plateNumber := c.Params("plateNumber") //车牌号是通过 URL 参数传递的

	spaces, err := psc.ParkingSpaceRepo.FindVehicleSpaceInLotByPlateNumber(plateNumber)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"spaces": spaces})
}

// 根据UserID查看停车位置
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

// 根据车位id获取状态
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

// 查看空闲车位
func (psc *ParkingSpaceController) GetFreeParkingSpace(c *fiber.Ctx) error {
	spaces, err := psc.ParkingSpaceRepo.FindFreeSpaceInAllLots()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"spaces": spaces})
}
