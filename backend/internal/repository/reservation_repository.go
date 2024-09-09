package repository

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/models"

type ReservationRepository interface {
	Create(reservation *models.Reservation) error
	UpdateStatusByReservationID(id int, status string) error
	Delete(reservationID int) error
}
