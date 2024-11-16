package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

type ParkingLotController struct {
	ParkingLotRepo repository.ParkingLotRepository
}

func NewParkingLotController(repo repository.ParkingLotRepository) *ParkingLotController {
	return &ParkingLotController{ParkingLotRepo: repo}
}

// CreateParkingLot 创建停车场
// @Summary Create a parking lot
// @Description Create a new parking lot with given details
// @Tags ParkingLot
// @Accept json
// @Produce json
// @Param parkingLot body models.ParkingLot true "Parking Lot Details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /parkinglot [post]
func (plc *ParkingLotController) CreateParkingLot(c *fiber.Ctx) error {
	var lot models.ParkingLot
	if err := c.BodyParser(&lot); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	success, err := plc.ParkingLotRepo.Create(&lot)
	if err != nil || !success {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Parking lot created successfully"})
}

// GetParkingLotByID 获取停车场信息
// @Summary Get parking lot by ID
// @Description Retrieve details of a parking lot by its ID
// @Tags ParkingLot
// @Accept json
// @Produce json
// @Param id path int true "Parking Lot ID"
// @Success 200 {object} models.ParkingLot
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /parkinglot/id/{id} [get]
func (plc *ParkingLotController) GetParkingLotByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	lot, err := plc.ParkingLotRepo.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Parking lot not found"})
	}

	return c.Status(fiber.StatusOK).JSON(lot)
}

// GetAllParkingLots 获取所有停车场
// @Summary Get all parking lots
// @Description Retrieve a list of all parking lots
// @Tags ParkingLot
// @Accept json
// @Produce json
// @Success 200 {array} models.ParkingLot
// @Failure 500 {object} map[string]string
// @Router /parkinglot [get]
func (plc *ParkingLotController) GetParkingLotsByName(c *fiber.Ctx) error {
	name := c.Params("name")

	lots, err := plc.ParkingLotRepo.FindByName(name)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Parking lot not found"})
	}

	return c.Status(fiber.StatusOK).JSON(lots)
}

// @Summary Get all parking lots
// @Description Retrieve a list of all parking lots
// @Tags ParkingLot
// @Accept json
// @Produce json
// @Success 200 {array} models.ParkingLot
// @Failure 500 {object} map[string]string
// @Router /parkinglot [get]
func (plc *ParkingLotController) GetAllParkingLots(c *fiber.Ctx) error {
	lots, err := plc.ParkingLotRepo.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(lots)
}

// @Summary Get all income by lot ID
// @Description Retrieve the total income of a parking lot by its ID
// @Tags ParkingLot
// @Accept json
// @Produce json
// @Param id path int true "Parking Lot ID"
// @Success 200 {object} float64
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /parkinglot/income/{id} [get]
func (plc *ParkingLotController) GetAllIncomeByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	income, err := plc.ParkingLotRepo.FindAllIncomeByLotID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(income)
}

// @Summary Update parking lot
// @Description Update the details of a parking lot
// @Tags ParkingLot
// @Accept json
// @Produce json
// @Param id path int true "Parking Lot ID"
// @Param parkingLot body models.ParkingLot true "Updated Parking Lot Details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /parkinglot/id/{id} [put]
func (plc *ParkingLotController) UpdateParkingLot(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var lot models.ParkingLot
	if err := c.BodyParser(&lot); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := plc.ParkingLotRepo.Update(&lot, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Parking lot updated successfully"})
}

// @Summary Delete parking lot
// @Description Delete a parking lot by its ID
// @Tags ParkingLot
// @Accept json
// @Produce json
// @Param id path int true "Parking Lot ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /parkinglot/id/{id} [delete]
func (plc *ParkingLotController) DeleteParkingLot(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := plc.ParkingLotRepo.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Parking lot deleted successfully"})
}

// @Summary Get parking lot occupancy rate
// @Description Retrieve the occupancy rate of a parking lot by its ID
// @Tags ParkingLot
// @Accept json
// @Produce json
// @Param id path int true "Parking Lot ID"
// @Success 200 {object} float64
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /parkinglot/occupancy-rate/{id} [get]
func (plc *ParkingLotController) GetOccupancyRateByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	occupancyRate, err := plc.ParkingLotRepo.FindOccupancyRateByLotID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(occupancyRate)
}

// @Summary Get parking lot occupancy by time
// @Description Retrieve the occupancy of a parking lot by its ID and time range
// @Tags ParkingLot
// @Accept json
// @Produce json
// @Param id path int true "Parking Lot ID"
// @Param start query string true "Start time (YYYY-MM-DDTHH:MM:SS)"
// @Param end query string true "End time (YYYY-MM-DDTHH:MM:SS)"
// @Success 200 {object} float64
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /parkinglot/id/{id}/start/{start}/end/{end} [get]
func (plc *ParkingLotController) GetOccupancyByIDAndTime(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	startStr := c.Params("start")
	endStr := c.Params("end")

	beijingLocation, _ := time.LoadLocation("Asia/Shanghai")

	start, err := time.ParseInLocation("2006-01-02T15:04:05", startStr, beijingLocation)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start time format; expected YYYY-MM-DDTHH:MM:SS"})
	}
	end, err := time.ParseInLocation("2006-01-02T15:04:05", endStr, beijingLocation)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end time format; expected YYYY-MM-DDTHH:MM:SS"})
	}

	occupancy, err := plc.ParkingLotRepo.FindOccupancyByLotIDAndTime(id, start, end)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(occupancy)
}

// @Summary Get parking lot status by ID
// @Description Retrieve the status of a parking lot by its ID
// @Tags ParkingLot
// @Accept json
// @Produce json
// @Param id path int true "Parking Lot ID"
// @Success 200 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /parkinglot/status/lot/{id} [get]
func (plc *ParkingLotController) GetStatusByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	free, _ := plc.ParkingLotRepo.GetFreeSpaceByLotID(id)
	occupied, _ := plc.ParkingLotRepo.GetOccupiedSpaceByLotID(id)
	reserved, _ := plc.ParkingLotRepo.GetReservedSpaceByLotID(id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"free": free, "occupied": occupied, "reserved": reserved})
}

// @Summary Get all income by lot ID
// @Description Retrieve the total income of all parking lots
// @Tags ParkingLot
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /parkinglot/allincome/all [get]
func (plc *ParkingLotController) GetAllIncomeByLotID(c *fiber.Ctx) error {
	var lotsInfo []fiber.Map

	for id := 1; id <= 10; id++ {
		// 查询停车场信息
		lot, err := plc.ParkingLotRepo.FindByID(id)
		if err != nil {
			// 如果找不到停车场，跳过并继续下一个
			continue
		}

		// 查询停车场收入
		income, err := plc.ParkingLotRepo.FindAllIncomeByLotID(id)
		if err != nil {
			// 如果找不到收入信息，跳过并继续下一个
			continue
		}

		// 准备响应数据
		lotsInfo = append(lotsInfo, fiber.Map{
			"ParkingLotID": lot.ParkingLotID,
			"ParkingName":  lot.ParkingName,
			"Longitude":    lot.Longitude,
			"Latitude":     lot.Latitude,
			"income":       income,
		})
	}

	// 返回所有停车场信息和收入
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total":       len(lotsInfo),
		"parkingLots": lotsInfo,
	})
}
