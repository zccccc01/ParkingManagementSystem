package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type VehicleRepositoryImpl struct {
	DB *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) *VehicleRepositoryImpl {
	return &VehicleRepositoryImpl{DB: db}
}

func (r *VehicleRepositoryImpl) Create(vehicle *models.Vehicle) (bool, error) {
	result := r.DB.Create(&vehicle)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r *VehicleRepositoryImpl) GetAllByVehicleID(id int) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	result := r.DB.First(&vehicle, "VehicleID = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &vehicle, nil
}

func (r *VehicleRepositoryImpl) GetAllByUserID(id int) ([]*models.Vehicle, error) {
	var vehicles []models.Vehicle
	result := r.DB.Where("UserID = ?", id).Find(&vehicles)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	// 将车辆列表转换为指向车辆的指针列表
	var vehiclePointers []*models.Vehicle
	for _, vehicle := range vehicles {
		vehiclePointers = append(vehiclePointers, &vehicle)
	}
	return vehiclePointers, nil
}

func (r *VehicleRepositoryImpl) UpdateVehicleByVehicleID(id int, vehicle *models.Vehicle) error {
	var existingVehicle models.Vehicle
	result := r.DB.First(&existingVehicle, "VehicleID = ?", id)
	if result.Error != nil {
		return result.Error
	}
	existingVehicle.PlateNumber = vehicle.PlateNumber
	existingVehicle.Color = vehicle.Color
	result = r.DB.Model(&existingVehicle).Updates(existingVehicle)
	return result.Error
}

func (r *VehicleRepositoryImpl) DeleteByVehicleID(id int) error {
	var existingVehicle models.Vehicle
	result := r.DB.Delete(&existingVehicle, "VehicleID = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
