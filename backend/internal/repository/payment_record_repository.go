package repository

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/models"

type PaymentRecordRepository interface {
	Create(payment *models.PaymentRecord) error
	GetAmountByRecordID(id int) (float64, error)
}
