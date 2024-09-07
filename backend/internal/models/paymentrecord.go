package models

import "time"

type PaymentRecord struct {
	PaymentID        int
	ReservationID    int
	Amount           float64
	PaymentTimeStamp time.Time
	PaymentMethod    string
}
