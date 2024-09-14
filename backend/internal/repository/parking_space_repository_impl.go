package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type ParkingSpaceRepositoryImpl struct {
	DB *gorm.DB
}

func NewParkingSpaceRepository(db *gorm.DB) ParkingSpaceRepository {
	return &ParkingSpaceRepositoryImpl{DB: db}
}

func (r *ParkingSpaceRepositoryImpl) Create(space *models.ParkingSpace) (bool, error) {
	result := r.DB.Create(&space)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r *ParkingSpaceRepositoryImpl) GetAllStatusByLotID(id int) ([]models.ParkingSpace, error) {
	var spaces []models.ParkingSpace
	result := r.DB.Find(&spaces, "ParkingLotID = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return spaces, nil
}

func (r *ParkingSpaceRepositoryImpl) GetStatusBySpaceID(id int) (string, error) {
	var space models.ParkingSpace
	result := r.DB.First(&space, "SpaceID = id", id)
	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		return "", gorm.ErrRecordNotFound
	}
	return space.Status, nil
}

func (r *ParkingSpaceRepositoryImpl) UpdateStatusBySpaceID(space *models.ParkingSpace, id int) (bool, error) {
	var existingSpace models.ParkingSpace
	result := r.DB.Model(&existingSpace).Where("SpaceID = ?", id).Updates(space)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, gorm.ErrRecordNotFound
	}
	return true, nil
}
