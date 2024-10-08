package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type ReservationRepositoryImpl struct {
	DB *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepositoryImpl {
	return &ReservationRepositoryImpl{DB: db}
}

func (r *ReservationRepositoryImpl) Create(reservation *models.Reservation) (bool, error) {
	result := r.DB.Create(&reservation)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r *ReservationRepositoryImpl) UpdateStatusByReservationID(id int, status string) error {
	var existingReservation models.Reservation
	result := r.DB.Model(&existingReservation).Where("ReservationID = ?", id).Update("Status", status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ReservationRepositoryImpl) DeleteByReservationID(id int) error {
	return r.DB.Delete(&models.Reservation{}, "ReservationID = ?", id).Error
}
