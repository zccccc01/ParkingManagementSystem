package models

type ParkingLot struct {
	ParkingLotID int `gorm:"primaryKey"`
	ParkingName  string
	Longitude    float64
	Latitude     float64
	Capacity     int
	Rates        float64
}
