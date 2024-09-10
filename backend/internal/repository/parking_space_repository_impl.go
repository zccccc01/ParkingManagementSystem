package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type ParkingSpaceRepositoryImpl struct {
	DB *gorm.DB
}

func NewParkingSpaceRepository(db *gorm.DB) ParkingRecordRepository {
	return &ParkingSpaceRepositoryImpl{DB: db}
}

func (r *ParkingSpaceRepositoryImpl) Create(space *models.ParkingSpace) error {
	return r.DB.Create(space).Error
}

func (r *ParkingSpaceRepositoryImpl) GetLotIDBySpaceID(id int) (int, error) {
	var space models.ParkingSpace
	result := r.DB.First(&space, "SpaceID = ?", id)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return -1, gorm.ErrRecordNotFound
	}
	return space.ParkingLotID, nil
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

func (r *ParkingSpaceRepositoryImpl) UpdateStatusBySpaceID(space *models.ParkingSpace, id int) error {
	return nil
}

func (r *ParkingSpaceRepositoryImpl) Delete(id int) error {
	return nil
}
