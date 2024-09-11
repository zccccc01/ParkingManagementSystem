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

func (r *UserRepositoryImpl) Create(user *models.User) error {
	return r.DB.Create(user).Error
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
