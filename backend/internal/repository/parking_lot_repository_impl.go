package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type ParkingLotRepositoryImpl struct {
	DB *gorm.DB
}

func NewParkingLotRepository(db *gorm.DB) ParkingLotRepository {
	return &ParkingLotRepositoryImpl{DB: db}
}

func (r *ParkingLotRepositoryImpl) Create(lot *models.ParkingLot) error {
	return r.DB.Create(lot).Error
}

func (r *ParkingLotRepositoryImpl) FindByID(id int) (*models.ParkingLot, error) {
	var lot models.ParkingLot
	result := r.DB.First(&lot, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &lot, nil
}

func (r *ParkingLotRepositoryImpl) FindByName(name string) (*models.ParkingLot, error) {
	var lot models.ParkingLot
	result := r.DB.First(&lot, "name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &lot, nil
}

func (r *ParkingLotRepositoryImpl) FindAll() ([]models.ParkingLot, error) {
	var lots []models.ParkingLot
	result := r.DB.Find(&lots)
	if result.Error != nil {
		return nil, result.Error
	}
	return lots, nil
}

func (r *ParkingLotRepositoryImpl) Update(lot *models.ParkingLot) error {
	return r.DB.Save(lot).Error
}

func (r *ParkingLotRepositoryImpl) Delete(id int) error {
	return r.DB.Delete(&models.ParkingLot{}, id).Error
}
