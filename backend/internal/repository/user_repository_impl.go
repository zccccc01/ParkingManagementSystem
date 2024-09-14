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

func (r *UserRepositoryImpl) UpdatePasswordByID(id int, password string) error {
	var existingUser models.User
	result := r.DB.Model(&existingUser).Where("UserID = ?", id).Update("Password", password)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *UserRepositoryImpl) UpdateTelByID(id int, tel string) error {
	var existingUser models.User
	result := r.DB.Model(&existingUser).Where("UserID = ?", id).Update("Tel", tel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
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

func (r *UserRepositoryImpl) Delete(id int) error {
	return r.DB.Delete(&models.User{}, "UserID = ?", id).Error
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
