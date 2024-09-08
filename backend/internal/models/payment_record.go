package models

import "time"

type PaymentRecord struct {
	PaymentID        int `gorm:"primaryKey"`
	ReservationID    int `gorm:"foreignKey:ReservationID"`
	Amount           float64
	PaymentTimeStamp time.Time
	PaymentMethod    string
}
