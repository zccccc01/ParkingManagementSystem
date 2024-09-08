package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type ParkingLotRepositoryImpl struct {
	DB *gorm.DB
}

func NewParkingLotRepository(db *gorm.DB) ParkingLotRepository {
	return &ParkingLotRepositoryImpl{DB: db}
}

func (r *ParkingLotRepositoryImpl) Create(lot *models.ParkingLot) error {
	return r.DB.Create(lot).Error
}

func (r *ParkingLotRepositoryImpl) FindByID(id int) (*models.ParkingLot, error) {
	var lot models.ParkingLot
	result := r.DB.First(&lot, "ParkingLotID = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &lot, nil
}

func (r *ParkingLotRepositoryImpl) FindByName(name string) (*models.ParkingLot, error) {
	var lot models.ParkingLot
	result := r.DB.First(&lot, "ParkingName = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &lot, nil
}

func (r *ParkingLotRepositoryImpl) FindAll() ([]models.ParkingLot, error) {
	var lots []models.ParkingLot
	result := r.DB.Find(&lots)
	if result.Error != nil {
		return nil, result.Error
	}
	return lots, nil
}

func (r *ParkingLotRepositoryImpl) Update(lot *models.ParkingLot, id int) error {
	// 先查询记录
	var existingLot models.ParkingLot
	result := r.DB.First(&existingLot, "ParkingLotID = ?", id)
	if result.Error != nil {
		return result.Error
	}

	var updates = map[string]interface{}{}
	if lot.ParkingName != existingLot.ParkingName {
		updates["ParkingName"] = lot.ParkingName
	}
	if lot.Longitude != existingLot.Longitude && !lot.Longitude.Equal(decimal.RequireFromString("0")) {
		updates["Longitude"] = lot.Longitude
	}
	if lot.Latitude != existingLot.Latitude && !lot.Latitude.Equal(decimal.RequireFromString("0")) {
		updates["Latitude"] = lot.Latitude
	}
	if lot.Capacity != existingLot.Capacity {
		updates["Capacity"] = lot.Capacity
	}
	if lot.Rates != existingLot.Rates && !lot.Rates.Equal(decimal.RequireFromString("0")) {
		updates["Rates"] = lot.Rates
	}
	// 使用 Model 和 Updates 方法
	result = r.DB.Model(&existingLot).Where("ParkingLotID = ?", id).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ParkingLotRepositoryImpl) Delete(id int) error {
	return r.DB.Delete(&models.ParkingLot{}, "ParkingLotID = ?", id).Error
}
