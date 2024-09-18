package repository

import (
	"time"

	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type ParkingRecordRepository interface {
	// 创建进场记录
	CreateRecordEntry(record *models.ParkingRecord) error
	// 根据ID更新出场记录
	UpdateRecordExitByRecordID(id int, now time.Time) error
	// 根据记录ID获取费用
	GetFeeByRecordID(id int) (float64, error)
	// 根据车辆ID获取费用
	GetFeeByVehicleID(id int) (float64, error)
	// 根据UserID查历史记录 TODO:返回值怪怪的?函数内建结构体,再返回?
	FindHistoryRecordByUserID(id int) (userID int, vehicleID int, inTime time.Time, outTime time.Time, Amount float64, err error)
}
