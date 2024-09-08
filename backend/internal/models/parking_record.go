package models

import "time"

type ParkingRecord struct {
	RecordID  int
	VehicleID int
	SpaceID   int
	LotID     int
	StartTime time.Time
	EndTime   time.Time
	Fee       float64
}
