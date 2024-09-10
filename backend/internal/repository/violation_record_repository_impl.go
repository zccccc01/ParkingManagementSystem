package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type ViolationRecordRepositoryImpl struct {
	DB *gorm.DB
}

func (r *ViolationRecordRepositoryImpl) Create(violation *models.ViolationRecord) error {
	return r.DB.Create(violation).Error
}
func (r *ViolationRecordRepositoryImpl) GetFineAmountByRecordID(id int) (float64, error) {
	var violationRecord models.ViolationRecord
	result := r.DB.Where("record_id = ?", id).First(&violationRecord)
	if result.Error != nil {
		return 0, result.Error
	}
	return violationRecord.FineAmount, nil
}
func (r *ViolationRecordRepositoryImpl) GetStatusByRecordID(id int) (string, error) {
	var violationRecord models.ViolationRecord
	result := r.DB.Where("record_id = ?", id).First(&violationRecord)
	if result.Error != nil {
		return "", result.Error
	}
	return violationRecord.Status, nil
}
func (r *ViolationRecordRepositoryImpl) GetViolationTypeByRecordID(id int) (string, error) {
	var violationRecord models.ViolationRecord
	result := r.DB.Where("record_id = ?", id).First(&violationRecord)
	if result.Error != nil {
		return "", result.Error
	}
	return violationRecord.ViolationType, nil
}
