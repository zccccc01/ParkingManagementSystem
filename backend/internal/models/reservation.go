package models

import "time"

type Reservation struct {
	ReservationID int `gorm:"primaryKey"`
	StartTime     time.Time
	EndTime       time.Time
	SpaceID       int `gorm:"forignKey:SpaceID"`
	VehicleID     int `gorm:"foreignKey:VehicleID"`
	LotID         int `gorm:"foreignKey:LotID"`
	Status        string
}
