package repository

import (
	"time"

	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

// ParkingLotRepository 定义停车场仓库接口
type ParkingLotRepository interface {
	// 创建一条记录
	Create(lot *models.ParkingLot) error
	// 根据ID查找一条记录
	FindByID(id int) (*models.ParkingLot, error)
	// 根据名称查找一条记录
	FindByName(name string) (*models.ParkingLot, error)
	// 查找所有记录
	FindAll() ([]models.ParkingLot, error)
	// 根据ID更新一条记录
	Update(lot *models.ParkingLot, id int) error
	// 根据ID删除一条记录
	Delete(id int) error
	// 查询某个时间段停车占用情况
	QueryOccupancy(start time.Time, end time.Time) ([]models.ParkingRecord, error)
	// 查询停车场收入
	QueryRevenue(start time.Time, end time.Time) (float64, error)
	// 查询违规停车统计
	QueryViolations(start time.Time, end time.Time) ([]models.ViolationRecord, error)
	// 根据ID查找停车位
	FindSpaceByID(spaceID int) (*models.ParkingSpace, error)
	// 查找停车场的违规记录
	FindViolationsByLotID(lotID int) ([]models.ViolationRecord, error)
}
