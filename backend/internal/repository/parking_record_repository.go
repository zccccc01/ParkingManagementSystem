package repository

import (
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type ParkingRecordRepository interface {
	// 创建进场记录
	CreateRecordEntry(record *models.ParkingRecord) error
	// 创建出场记录
	CreateRecordExit(record *models.ParkingRecord) error
	// 根据记录ID获取费用
	GetFeeByRecordID(id int) (float64, error)
	// 根据车辆ID获取费用
	GetFeeByVehicleID(id int) (float64, error)
}
