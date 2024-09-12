package models

type Vehicle struct {
	VehicleID   int    `gorm:"column:VehicleID;primaryKey"`
	UserID      int    `gorm:"column:UserID;foreignKey:UID"`
	PlateNumber string `gorm:"column:PlateNumber;size:255"`
	Color       string `gorm:"column:Color;size:255"`
}

func (v *Vehicle) TableName() string {
	return "vehicle"
}
