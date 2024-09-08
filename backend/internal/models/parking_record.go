package models

import "time"

type ParkingRecord struct {
	RecordID  int `gorm:"primaryKey"`
	VehicleID int `gorm:"foreignKey:VehicleID"`
	SpaceID   int `gorm:"foreignKey:SpaceID"`
	LotID     int `gorm:"foreignKey:LotID"`
	StartTime time.Time
	EndTime   time.Time
	Fee       float64
}
