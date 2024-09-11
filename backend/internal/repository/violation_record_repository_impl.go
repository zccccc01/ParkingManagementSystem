package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type ViolationRecordRepositoryImpl struct {
	DB *gorm.DB
}

func NewViolationRecordRepository(db *gorm.DB) *ViolationRecordRepositoryImpl {
	return &ViolationRecordRepositoryImpl{DB: db}
}

func (r *ViolationRecordRepositoryImpl) Create(violation *models.ViolationRecord) error {
	return r.DB.Create(violation).Error
}

func (r *ViolationRecordRepositoryImpl) GetFineAmountByRecordID(id int) (float64, error) {
	var violationRecord models.ViolationRecord
	result := r.DB.First(&violationRecord, "RecordId = ?", id)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return -1, gorm.ErrRecordNotFound
	}
	return violationRecord.FineAmount, nil
}

func (r *ViolationRecordRepositoryImpl) GetStatusByRecordID(id int) (string, error) {
	var violationRecord models.ViolationRecord
	result := r.DB.First(&violationRecord, "RecordId = ?", id)
	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		return "", gorm.ErrRecordNotFound
	}
	return violationRecord.Status, nil
}

func (r *ViolationRecordRepositoryImpl) GetViolationTypeByRecordID(id int) (string, error) {
	var violationRecord models.ViolationRecord
	result := r.DB.First(&violationRecord, "RecordId = ?", id)
	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		return "", gorm.ErrRecordNotFound
	}
	return violationRecord.ViolationType, nil
}
