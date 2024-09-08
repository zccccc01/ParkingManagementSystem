package models

type ParkingSpace struct {
	SpaceID      int `gorm:"primaryKey"`
	Number       int
	Status       string
	ParkingLotID int `gorm:"forignKey:ParkingLotID"`
}
