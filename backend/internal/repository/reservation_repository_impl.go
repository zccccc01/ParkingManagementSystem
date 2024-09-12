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

func (r *ReservationRepositoryImpl) Create(reservation *models.Reservation) error {
	return r.DB.Create(reservation).Error
}

func (r *ReservationRepositoryImpl) UpdateStatusByReservationID(id int, status string) error {
	var existingReservation models.Reservation
	result := r.DB.Model(&existingReservation).Where("ReservationID = ?", id).Update("Status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *ReservationRepositoryImpl) DeleteByReservationID(id int) error {
	return r.DB.Delete(&models.Reservation{}, "ReservationID = ?", id).Error
}
