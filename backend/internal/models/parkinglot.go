package models

type ParkingLot struct {
	ParkingLotID int
	ParkingName  string
	Longitude    float64
	Latitude     float64
	Capacity     int
	Rates        float64
}
