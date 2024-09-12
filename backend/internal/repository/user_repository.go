package repository

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/models"

type UserRepository interface {
	// 创建一条记录
	Create(user *models.User) error
	// 根据ID更新密码
	UpdatePasswordByID(id int, password string) error
	// 根据ID更新电话
	UpdateTelByID(id int, tel string) error
	// 根据ID获取电话
	GetTelByID(id int) (string, error)
	// 根据ID删除记录
	Delete(id int) error
}
