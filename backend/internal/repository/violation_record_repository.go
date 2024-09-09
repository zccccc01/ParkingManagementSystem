package repository

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/models"

type ViolationRecordRepository interface {
	Create(violation *models.ViolationRecord) error
	GetFineAmountByRecordID(id int) (float64, error)
	GetStatusByRecordID(id int) (string, error)
	GetViolationTypeByRecordID(id int) (string, error)
}
