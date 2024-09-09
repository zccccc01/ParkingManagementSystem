package repository

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type ParkingRecordRepositoryImpl struct {
	DB *gorm.DB
}

func NewParkingRecordRepository(db *gorm.DB) ParkingRecordRepository {
	return &ParkingRecordRepositoryImpl{DB: db}
}

func (r *ParkingRecordRepositoryImpl) CreateRecordEntry(record *models.ParkingRecord) error {
	farFutureTime := time.Now().Add(100 * 365 * 24 * time.Hour)
	tmp := models.ParkingRecord{
		RecordID:  record.RecordID,
		VehicleID: record.VehicleID,
		SpaceID:   record.SpaceID,
		LotID:     record.LotID,
		StartTime: time.Now(),
		EndTime:   farFutureTime,
	}
	result := r.DB.Save(tmp)
	return result.Error
}

func (r *ParkingRecordRepositoryImpl) UpdateRecordExitByRecordID(id int, now time.Time) error {
	var existingRecord models.ParkingRecord
	result := r.DB.First(&existingRecord, "RecordID = ?", id)
	if result.Error != nil {
		return result.Error
	}
	existingRecord.EndTime = now
	result = r.DB.Model(&existingRecord).Update("EndTime", now)
	return result.Error
}

func (r *ParkingRecordRepositoryImpl) GetFeeByRecordID(id int) (float64, error) {
	var existingRecord models.ParkingRecord
	result := r.DB.First(&existingRecord, "RecordID = ?", id)
	if result.Error != nil {
		return 0, result.Error
	}
	start := existingRecord.StartTime
	end := existingRecord.EndTime
	timeDiff := end.Sub(start)
	lotID := existingRecord.LotID
	var existingLot models.ParkingLot
	result = r.DB.First(&existingLot, "ParkingLotID = ?", lotID)
	if result.Error != nil {
		return 0, result.Error
	}
	// 这个方法将decimal类型转化为float64类型
	rate, _ := existingLot.Rates.Float64()
	return float64(timeDiff.Hours()) * rate, nil
}

func (r *ParkingRecordRepositoryImpl) GetFeeByVehicleID(id int) (float64, error) {
	var existingRecord models.ParkingRecord
	result := r.DB.First(&existingRecord, "VehicleID = ?", id)
	if result.Error != nil {
		return 0, result.Error
	}
	start := existingRecord.StartTime
	end := existingRecord.EndTime
	timeDiff := end.Sub(start)
	lotID := existingRecord.LotID
	var existingLot models.ParkingLot
	result = r.DB.First(&existingLot, "ParkingLotID = ?", lotID)
	if result.Error != nil {
		return 0, result.Error
	}
	// 这个方法将decimal类型转化为float64类型
	rate, _ := existingLot.Rates.Float64()
	return float64(timeDiff.Hours()) * rate, nil
}
