package repository

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/models"

type VehicleRepository interface {
	//创建一条记录
	Create(vehicle *models.Vehicle) error
	//根据车辆id获取一条记录
	GetAllByVehicleID(id int) (*models.Vehicle, error)
	//根据用户id获取所有记录
	GetAllByUserID(id int) ([]*models.Vehicle, error)
	//根据车辆id更新车牌号和颜色,用户id是外键,不能更改
	UpdateVehicleByVehicleID(id int, vehicle *models.Vehicle) error
	//根据车辆id删除一条记录
	DeleteByVehicleID(id int) error
}
