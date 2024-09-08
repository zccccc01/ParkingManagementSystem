package models

type ViolationRecord struct {
	ViolationID   int
	RecordID      int
	FineAmount    float64
	ViolationType string
	Status        string
}
