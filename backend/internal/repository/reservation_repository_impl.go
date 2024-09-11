package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type ReservationRepositoryImpl struct {
	DB *gorm.DB
}

// NewReservationRepository is a helper function to create a new ReservationRepository
func NewReservationRepository(db *gorm.DB) *ReservationRepositoryImpl {
	return &ReservationRepositoryImpl{DB: db}
}
func (r *ReservationRepositoryImpl) Create(reservation *models.Reservation) error {
	result := r.DB.Create(reservation)
	return result.Error
}

func (r *ReservationRepositoryImpl) UpdateStatusByReservationID(id int, status string) error {
	result := r.DB.Model(&models.Reservation{}).Where("ReservationID = ?", id).Update("Status", status)
	return result.Error
}

func (r *ReservationRepositoryImpl) DeleteByReservationID(id int) error {
	result := r.DB.Delete(&models.Reservation{}, id)
	return result.Error
}
