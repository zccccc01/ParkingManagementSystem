package repository

import (
	"github.com/jinzhu/gorm"
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

// TODO: 按ID找,其余更新,想要的效果是:
/*	创建一个 ParkingLot 实例并设置字段
	lot := &models.ParkingLot{
		ParkingLotID: 1,
		ParkingName:  "Central Parking Updated",
		// Longitude:    decimal.RequireFromString("13"),
		// Latitude:     decimal.RequireFromString("65"),
		Capacity:     150,
		Rates:        decimal.RequireFromString("90"),
	}
	这样更新,Longitude和Latitude仍然是原值而不改为0
*/
func (r *ParkingLotRepositoryImpl) Update(lot *models.ParkingLot, id int) error {
	// 先查询记录
	var existingLot models.ParkingLot
	result := r.DB.First(&existingLot, "ParkingLotID = ?", id)
	if result.Error != nil {
		return result.Error
	}
	// 只更新非零值字段
	var updates = map[string]interface{}{
		"ParkingName": lot.ParkingName,
		"Longitude":   lot.Longitude,
		"Latitude":    lot.Latitude,
		"Capacity":    lot.Capacity,
		"Rates":       lot.Rates,
	}
	// var updates = map[string]interface{}{}
	// if lot.ParkingName == "" {
	// 	updates["ParkingName"] = existingLot.ParkingName
	// }
	// updates["ParkingName"] = lot.ParkingName
	// if lot.Longitude == decimal.RequireFromString("0") {
	// 	updates["Longitude"] = existingLot.Longitude
	// }
	// updates["Longitude"] = lot.Longitude
	// if lot.Latitude == decimal.RequireFromString("0") {
	// 	updates["Latitude"] = existingLot.Latitude
	// }
	// updates["Latitude"] = lot.Latitude
	// if lot.Capacity == 0 {
	// 	updates["Capacity"] = existingLot.Capacity
	// }
	// updates["Capacity"] = lot.Capacity
	// if lot.Rates == decimal.RequireFromString("0") {
	// 	updates["Rates"] = existingLot.Rates
	// }
	// updates["Rates"] = lot.Rates
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
