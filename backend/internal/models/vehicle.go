package models

type Vehicle struct {
	VehicleID   int `gorm:"primaryKey"`
	UserID      int `gorm:"foreignKey:UserID"`
	PlateNumber string
	Color       string
}
