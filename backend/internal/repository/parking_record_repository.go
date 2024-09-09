package repository

import (
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

// ParkingRecordRepository 定义停车场记录接口
type ParkingRecordRepository interface {
	// 创建进场记录
	RecordEntry(record *models.ParkingRecord) error
	// 创建出场记录
	RecordExit(record *models.ParkingRecord) error
	// 根据记录ID获取费用
	GetFeeByRecordID(recordID int) (float64, error)
	// 根据车辆ID获取费用
	GetFeeByVehicleID(vehicleID int) (float64, error)
}
