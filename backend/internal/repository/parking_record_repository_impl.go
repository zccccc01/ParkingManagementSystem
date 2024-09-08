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

// 设置进场时间
func (r *ParkingRecordRepositoryImpl) RecordEntry(record *models.ParkingRecord) error {
	record.StartTime = time.Now()
	result := r.DB.Save(record)
	return result.Error
}

// 设置出场时间
func (r *ParkingRecordRepositoryImpl) RecordExit(record *models.ParkingRecord) error {
	record.EndTime = time.Now()
	result := r.DB.Save(record)
	return result.Error
}

func (r *ParkingRecordRepositoryImpl) GetFeeByRecordID(recordID int) (float64, error) {
	return 0, nil
}
func (r *ParkingRecordRepositoryImpl) GetFeeByVehicleID(vehicleID int) (float64, error) {
	return 0, nil
}
