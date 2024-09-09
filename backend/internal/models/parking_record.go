package models

import "time"

type ParkingRecord struct {
	RecordID  int       `gorm:"column:RecordID;primaryKey"`
	VehicleID int       `gorm:"column:VehicleID;foreignKey:VEID"`
	SpaceID   int       `gorm:"column:SpaceID;foreignKey:PSID"`
	LotID     int       `gorm:"column:LotID;foreignKey:PALID"`
	StartTime time.Time `gorm:"column:StartTime"`
	EndTime   time.Time `gorm:"column:EndTime"`
	Fee       float64   `gorm:"column:Fee"`
}

// 设置ParkingRecord表名为`parkingrecord`
func (p *ParkingRecord) TableName() string {
	return "parkingrecord"
}
