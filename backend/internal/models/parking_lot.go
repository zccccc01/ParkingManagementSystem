package models

import (
	"github.com/shopspring/decimal"
)

// GORM 会自动将结构体内的成员改名为Parking_Lot_ID(就是给你加下划线)
// 所以需要打标签告诉gorm列名是什么 `gorm:"column:ParkingLotID;primaryKey"`
// 写一个方法告诉他表名是什么
type ParkingLot struct {
	ParkingLotID int             `gorm:"column:ParkingLotID;primaryKey"`
	ParkingName  string          `gorm:"column:ParkingName"`
	Longitude    decimal.Decimal `gorm:"type:decimal(9,6)"`
	Latitude     decimal.Decimal `gorm:"type:decimal(9,6)"`
	Capacity     int             `gorm:"default:null"`
	Rates        decimal.Decimal `gorm:"type:decimal(10,2)"`
}

func (p *ParkingLot) TableName() string {
	return "parkinglot"
}
