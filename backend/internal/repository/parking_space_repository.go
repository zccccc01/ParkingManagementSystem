package repository

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/models"

type ParkingSpaceRepository interface {
	// 创建一条记录
	Create(space *models.ParkingSpace) error
	// 根据停车场id获取该停车场车位空余情况
	GetAllStatusByLotID(id int) ([]models.ParkingSpace, error)
	// 根据车位id获取状态 TODO:(前提是知道LotID),需要修改/删除
	GetStatusBySpaceID(id int) (string, error)
	// 根据车位id更新状态
	UpdateStatusBySpaceID(space *models.ParkingSpace, id int) error
}
