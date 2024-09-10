package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type VehicleRepositoryImpl struct {
	DB *gorm.DB
}

// NewVehicleRepository is a helper function to create a new VehicleRepository
func NewVehicleRepository(db *gorm.DB) *VehicleRepositoryImpl {
	return &VehicleRepositoryImpl{DB: db}
}
func (r *VehicleRepositoryImpl) Create(vehicle *models.Vehicle) error {
	result := r.DB.Create(vehicle)
	return result.Error
}

func (r *VehicleRepositoryImpl) GetAllByVehicleID(id int) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	result := r.DB.Where("VehicleID = ?", id).First(&vehicle)
	if result.Error != nil {
		return nil, result.Error
	}
	return &vehicle, nil
}

func (r *VehicleRepositoryImpl) GetAllByUserID(id int) ([]*models.Vehicle, error) {
	var vehicles []models.Vehicle
	result := r.DB.Where("UserID = ?", id).Find(&vehicles)
	if result.Error != nil {
		return nil, result.Error
	}
	// 将车辆列表转换为指向车辆的指针列表
	var vehiclePointers []*models.Vehicle
	for _, vehicle := range vehicles {
		vehiclePointers = append(vehiclePointers, &vehicle)
	}
	return vehiclePointers, nil

}

func (r *VehicleRepositoryImpl) UpdateVehicleByVehicleID(id int, vehicle *models.Vehicle) error {
	result := r.DB.Model(&models.Vehicle{}).Where("VehicleID = ?", id).Updates(models.Vehicle{})

	return result.Error
}

func (r *VehicleRepositoryImpl) DeleteByVehicleID(id int) error {
	result := r.DB.Delete(&models.Vehicle{}, "VehicleID = ?", id)
	return result.Error
}
