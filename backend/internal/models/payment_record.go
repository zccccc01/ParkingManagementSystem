package models

import "time"

type PaymentRecord struct {
	PaymentID        int       `gorm:"column:PaymentID;primaryKey"`
	RecordID         int       `gorm:"column:RecordID;foreignKey:RECID"`
	ReservationID    int       `gorm:"column:ReservationID;foreignKey:RID"`
	Amount           float64   `gorm:"column:Amount"`
	PaymentTimestamp time.Time `gorm:"column:PaymentTimestamp"`
	PaymentMethod    string    `gorm:"column:PaymentMethod"`
}
