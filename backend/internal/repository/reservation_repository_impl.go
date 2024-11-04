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
	reservation.Status = "Doing"
	lotID := reservation.LotID
	spaceID := reservation.SpaceID
	// update parkingspace set Status = "Reserved" where ParkingLotID = ? and SpaceID = ?
	query := "UPDATE parkingspace SET Status = 'Reserved' WHERE ParkingLotID = ? AND SpaceID = ?"
	// Exec这个方法是直接执行sql语句
	r.DB.Exec(query, lotID, spaceID)
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

func (r *ReservationRepositoryImpl) UpdateByReservationID(id int, reservation *models.Reservation) error {
	var existingReservation models.Reservation
	result := r.DB.Model(&existingReservation).Where("ReservationID = ?", id).Updates(reservation)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
