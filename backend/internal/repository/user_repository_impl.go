package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

// NewUserRepository is a helper function to create a new UserRepository
func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}
func (r *UserRepositoryImpl) Create(user *models.User) error {
	result := r.DB.Create(user)
	return result.Error
}

func (r *UserRepositoryImpl) UpdatePasswordByID(id int, password string) error {
	result := r.DB.Model(&models.User{}).Where("UserID = ?", id).Update("Password", password)
	return result.Error
}

func (r *UserRepositoryImpl) UpdateTelByID(id int, tel string) error {
	result := r.DB.Model(&models.User{}).Where("UserID = ?", id).Update("Tel", tel)
	return result.Error
}

func (r *UserRepositoryImpl) GetTelByID(id int) (string, error) {
	var user models.User
	result := r.DB.First(&user, "UserID = ?", id)
	if result.Error != nil {
		return "", result.Error
	}
	return user.Tel, nil
}

func (r *UserRepositoryImpl) Delete(id int) error {
	result := r.DB.Delete(&models.User{}, id)
	return result.Error
}
