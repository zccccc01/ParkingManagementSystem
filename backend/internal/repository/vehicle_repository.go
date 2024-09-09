package repository

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/models"

type VehicleRepository interface {
	Create(vehicle *models.Vehicle) error
	GetAllByVehicleID(id int) (*models.Vehicle, error)
	GetAllByUserID(id int) ([]*models.Vehicle, error)
	UpdateVehicleByVehicleID(id int) error
	DeleteByVehicleID(id int) error
}
