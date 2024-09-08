package models

type ParkingSpace struct {
	SpaceID      int    `gorm:"column:SpaceID;primaryKey"`
	Status       string `gorm:"column:Status"`
	ParkingLotID int    `gorm:"column:ParkingLotID;forignKey:ParkingLotID"`
}
