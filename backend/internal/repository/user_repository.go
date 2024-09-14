package repository

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/models"

type UserRepository interface {
	// 创建一条记录
	Create(user *models.User) (bool, error)
	// 根据ID更新密码
	UpdatePasswordByID(id int, password string) (bool, error)
	// 根据ID更新电话
	UpdateTelByID(id int, tel string) (bool, error)
	// 根据ID获取电话
	GetTelByID(id int) (string, error)
	// 根据ID删除记录
	Delete(id int) (bool, error)
	// 根据ID查有无此人,用于查询是否被注册
	HasUserByID(id int) (bool, error)
	// 根据Tel查有无此人,用于查询是否被注册
	HasUserByTel(tel string) (bool, error)
	// 根据Tel和ID查记录,更改userName
	UpdateUserName(id int, tel string, newname string) (bool, error)
	// 根据ID查记录,用于登录
	FindUserByID(id int) (*models.User, error)
	// 根据Tel查记录,用于登录
	FindUserByTel(tel string) (*models.User, error)
}
