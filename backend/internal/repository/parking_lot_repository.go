package repository

import (
	"time"

	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

// ParkingLotRepository 定义停车场仓库接口
type ParkingLotRepository interface {
	// 创建一条记录
	Create(lot *models.ParkingLot) (bool, error)
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
	// 获取停车场的总收入
	FindAllIncomeByLotID(id int) (float64, error)
	// 获取停车场的占用率
	FindOccupancyRateByLotID(id int) (float64, error)
	// 按时间获取停车场的停车情况
	FindOccupancyByLotIDAndTime(id int, start time.Time, end time.Time) ([]models.ParkingSpace, error)
	// 可视化返回的三个状态的停车位占用饼图
	GetFreeSpaceByLotID(id int) (int, error)
	GetOccupiedSpaceByLotID(id int) (int, error)
	GetReservedSpaceByLotID(id int) (int, error)
}
