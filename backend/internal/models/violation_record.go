package models

type ViolationRecord struct {
	ViolationID   int `gorm:"primaryKey"`
	RecordID      int `gorm:"foreignKey:RecordID"`
	FineAmount    float64
	ViolationType string
	Status        string
}
