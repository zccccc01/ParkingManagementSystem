package models

import "time"

type Reservation struct {
	ReservationID int
	StartTime     time.Time
	EndTime       time.Time
	SpaceID       int
	VehicleID     int
	LotID         int
	Status        string
}
