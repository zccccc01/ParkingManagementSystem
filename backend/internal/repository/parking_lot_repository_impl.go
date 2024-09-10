package repository

import (
	"time"

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

// QueryOccupancy 查询某个时间段停车占用情况
func (r *ParkingLotRepositoryImpl) QueryOccupancy(start time.Time, end time.Time) ([]models.ParkingRecord, error) {
	var records []models.ParkingRecord
	result := r.DB.Where("start_time >= ? AND end_time <= ?", start, end).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}

// QueryRevenue 查询停车场收入
func (r *ParkingLotRepositoryImpl) QueryRevenue(start time.Time, end time.Time) (float64, error) {
	var totalRevenue float64
	result := r.DB.Model(&models.ParkingRecord{}).
		Where("end_time >= ? AND end_time <= ?", start, end).
		Select("SUM(fee) AS total_fee").
		Scan(&totalRevenue)
	if result.Error != nil {
		return 0, result.Error
	}
	return totalRevenue, nil
}

// QueryViolations 查询违规停车统计
func (r *ParkingLotRepositoryImpl) QueryViolations(start time.Time, end time.Time) ([]models.ViolationRecord, error) {
	var violations []models.ViolationRecord
	result := r.DB.Where("occurrence_time >= ? AND occurrence_time <= ?", start, end).Find(&violations)
	if result.Error != nil {
		return nil, result.Error
	}
	return violations, nil
}

// FindSpaceByID 根据ID查找停车位
func (r *ParkingLotRepositoryImpl) FindSpaceByID(spaceID int) (*models.ParkingSpace, error) {
	var space models.ParkingSpace
	result := r.DB.First(&space, "id = ?", spaceID)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &space, nil
}

// FindViolationsByLotID 查找停车场的违规记录
func (r *ParkingLotRepositoryImpl) FindViolationsByLotID(lotID int) ([]models.ViolationRecord, error) {
	var violations []models.ViolationRecord
	result := r.DB.Where("lot_id = ?", lotID).Find(&violations)
	if result.Error != nil {
		return nil, result.Error
	}
	return violations, nil
}
