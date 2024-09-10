package repository

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/models"

type ParkingSpaceRepository interface {
	// 创建一条记录
	Create(space *models.ParkingSpace) error
	// 根据车位id获取停车场id记录,TODO:应该是从LotID连表查spaceID
	GetLotIDBySpaceID(id int) (int, error)
	// 根据车位id获取状态
	GetStatusBySpaceID(id int) (string, error)
	// 根据车位id更新状态
	UpdateStatusBySpaceID(space *models.ParkingSpace, id int) error
	// 根据车位id删除一条记录
	Delete(id int) error
}
