package models

type ViolationRecord struct {
	ViolationID   int     `gorm:"column:ViolationID;primaryKey"`
	RecordID      int     `gorm:"column:RecordID;foreignKey:REID"`
	FineAmount    float64 `gorm:"column:FineAmount"`
	ViolationType string  `gorm:"column:ViolationType"`
	Status        string  `gorm:"column:Status"`
}
