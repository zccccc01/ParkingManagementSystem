package repository

import (
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
	result := r.DB.First(&existingUser, "Tel = ?", tel)

	// 检查是否发生了数据库错误，但忽略记录未找到的情况
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}
	// 如果没有找到任何记录，result.RowsAffected 会是 0
	// TODO:全部要修改这个RowsAffected
	// 如果没有找到任何记录，则返回 false，没有错误
	if result.RowsAffected == 0 {
		return false, nil
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
	var existingUser models.User
	result := r.DB.Model(&existingUser).Where("UserID = ?", id).Update("UserName", newName)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, gorm.ErrRecordNotFound
	}
	return true, nil
}
