package models

import "time"

type ParkingRecord struct {
	RecordID  int       `gorm:"column:Record;primaryKey"`
	VehicleID int       `gorm:"column:VehicleID;foreignKey:VehicleID"`
	SpaceID   int       `gorm:"column:SpaceID;foreignKey:SpaceID"`
	LotID     int       `gorm:"column:LotID;foreignKey:LotID"`
	StartTime time.Time `gorm:"column:StartTime"`
	EndTime   time.Time `gorm:"column:EndTime"`
	Fee       float64   `gorm:"column:Fee"`
}
