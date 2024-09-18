package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) Create(user *models.User) (bool, error) {
	result := r.DB.Create(user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r *UserRepositoryImpl) UpdatePasswordByID(id int, password string) (bool, error) {
	var existingUser models.User
	result := r.DB.Model(&existingUser).Where("UserID = ?", id).Update("Password", password)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, gorm.ErrRecordNotFound
	}
	return true, nil
}

func (r *UserRepositoryImpl) UpdateTelByID(id int, tel string) (bool, error) {
	var existingUser models.User
	result := r.DB.Model(&existingUser).Where("UserID = ?", id).Update("Tel", tel)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, gorm.ErrRecordNotFound
	}
	return true, nil
}

func (r *UserRepositoryImpl) GetTelByID(id int) (string, error) {
	var user models.User
	result := r.DB.First(&user, "UserID = ?", id)
	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		return "", gorm.ErrRecordNotFound
	}
	return user.Tel, nil
}

func (r *UserRepositoryImpl) Delete(id int) (bool, error) {
	result := r.DB.Delete(&models.User{}, "UserID = ?", id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r *UserRepositoryImpl) HasUserByID(id int) (bool, error) {
	var existingUser models.User
	result := r.DB.First(&existingUser, "UserID = ?", id)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, gorm.ErrRecordNotFound
	}
	return true, nil
}

func (r *UserRepositoryImpl) HasUserByTel(tel string) (bool, error) {
	var existingUser models.User
	result := r.DB.Find(&existingUser, "Tel = ?", tel)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, gorm.ErrRecordNotFound
	}
	return true, nil
}

func (r *UserRepositoryImpl) UpdateUserName(id int, tel string, newname string) (bool, error) {
	var existingUser models.User
	result := r.DB.Model(&existingUser).Where("UserID = ? AND Tel = ?", id, tel).Update("UserName", newname)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, gorm.ErrRecordNotFound
	}
	return true, nil
}

func (r *UserRepositoryImpl) FindUserByID(id int) (*models.User, error) {
	var user models.User
	result := r.DB.First(&user, "UserID = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindUserByTel(tel string) (*models.User, error) {
	var user models.User
	result := r.DB.First(&user, "Tel = ?", tel)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

func (r *UserRepositoryImpl) UpdateUserNameByID(id int, newName string) (bool, error) {

	// 首先，从数据库中查找对应的用户
	var user models.User
	result := r.DB.First(&user, id)
	if result.Error != nil {
		// 如果没有找到用户或者发生其他错误，返回错误
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil // 可以选择返回 nil 错误表示用户不存在
		}
		return false, result.Error
	}

	// 更新用户名
	user.UserName = newName
	result = r.DB.Save(&user)
	if result.Error != nil {
		// 如果更新失败，返回错误
		return false, result.Error
	}

	// 如果一切顺利，返回 true 表示成功
	return true, nil
}
