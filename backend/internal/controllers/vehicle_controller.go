package controllers

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"

type VehicleController struct {
	VehicleRepo repository.VehicleRepository
}

func NewVehicleController(repo repository.VehicleRepository) *VehicleController {
	return &VehicleController{VehicleRepo: repo}
}
