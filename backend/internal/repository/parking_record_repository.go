package repository

import (
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

// ParkingRecordRepository 定义停车场记录接口
type ParkingRecordRepository interface {
	RecordEntry(record *models.ParkingRecord) error
	RecordExit(record *models.ParkingRecord) error
	GetFeeByRecordID(recordID int) (float64, error)
	GetFeeByVehicleID(vehicleID int) (float64, error)
}
