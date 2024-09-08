package models

import "time"

type PaymentRecord struct {
	PaymentID        int       `gorm:"primaryKey"`
	ReservationID    int       `gorm:"foreignKey:ReservationID"`
	Amount           float64   `gorm:"column:Amount"`
	PaymentTimeStamp time.Time `gorm:"column:PaymentTimeStamp"`
	PaymentMethod    string
}
