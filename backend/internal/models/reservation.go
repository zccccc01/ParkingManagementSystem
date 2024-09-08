package models

import "time"

type Reservation struct {
	ReservationID int       `gorm:"column:ReservationID;primaryKey"`
	StartTime     time.Time `gorm:"column:StartTime"`
	EndTime       time.Time `gorm:"column:EndTime"`
	SpaceID       int       `gorm:"column:SpaceID;foreignKey:SID"`
	VehicleID     int       `gorm:"column:VehicleID;foreignKey:VID"`
	LotID         int       `gorm:"column:LotID;foreignKey:LID"`
	Status        string    `gorm:"column:Status"`
}
