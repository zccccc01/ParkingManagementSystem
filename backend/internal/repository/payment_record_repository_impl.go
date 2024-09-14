package repository

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type PaymentRecordRepositoryImpl struct {
	DB *gorm.DB
}

func NewPaymentRecordRepository(db *gorm.DB) PaymentRecordRepository {
	return &PaymentRecordRepositoryImpl{DB: db}
}

func (r *PaymentRecordRepositoryImpl) Create(payment *models.PaymentRecord) (bool, error) {
	result := r.DB.Create(&payment)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r *PaymentRecordRepositoryImpl) GetAmountByRecordID(id int) (float64, error) {
	var payment models.PaymentRecord
	result := r.DB.First(&payment, "RecordID=?", id)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, gorm.ErrRecordNotFound
	}
	return payment.Amount, nil
}

func (r *PaymentRecordRepositoryImpl) GetAmountByReservationID(id int) (float64, error) {
	var payment models.PaymentRecord
	result := r.DB.Find(&payment, "ReservationID=?", id)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, gorm.ErrRecordNotFound
	}
	return payment.Amount, nil
}

func (r *PaymentRecordRepositoryImpl) GetPaymentMethodByPaymentID(id int) (string, error) {
	var payment models.PaymentRecord
	result := r.DB.First(&payment, "PaymentID=?", id)
	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		return "", gorm.ErrRecordNotFound
	}
	return payment.PaymentMethod, nil
}

func (r *PaymentRecordRepositoryImpl) GetPaymentTimeStampByPaymentID(id int) (time.Time, error) {
	var payment models.PaymentRecord
	result := r.DB.First(&payment, "PaymentID=?", id)
	if result.Error != nil {
		return time.Time{}, result.Error
	}
	if result.RowsAffected == 0 {
		return time.Time{}, gorm.ErrRecordNotFound
	}
	return payment.PaymentTimestamp, nil
}

func (r *PaymentRecordRepositoryImpl) GetPaymentStatusByPaymentTimeStamp(timestamp time.Time) (string, error) {
	var payment models.PaymentRecord
	result := r.DB.Find(&payment, "PaymentTimestamp=?", timestamp)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return "NOPAY", nil
		}
		return "", result.Error
	}
	return "PAY", nil
}
