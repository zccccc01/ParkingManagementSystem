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

func (r *ViolationRecordRepositoryImpl) GetFineAmountByRecordID(id int) ([]float64, error) {
	var violationRecords []models.ViolationRecord
	result := r.DB.Where("RecordID = ?", id).Find(&violationRecords)
	if result.Error != nil {
		return nil, result.Error
	}
	var amounts []float64
	for _, record := range violationRecords {
		amounts = append(amounts, record.FineAmount)
	}
	return amounts, nil
}

func (r *ViolationRecordRepositoryImpl) GetStatusByRecordID(id int) ([]string, error) {
	var violationRecords []models.ViolationRecord
	result := r.DB.Where("RecordID = ?", id).Find(&violationRecords)
	if result.Error != nil {
		return nil, result.Error
	}
	var statuses []string
	for _, RecordID := range violationRecords {
		statuses = append(statuses, RecordID.Status)
	}
	return statuses, nil
}

func (r *ViolationRecordRepositoryImpl) GetViolationTypeByRecordID(id int) ([]string, error) {
	var violationRecords []models.ViolationRecord
	result := r.DB.Where("RecordID = ?", id).Find(&violationRecords)
	if result.Error != nil {
		return nil, result.Error
	}
	var types []string
	for _, RecordID := range violationRecords {
		types = append(types, RecordID.ViolationType)
	}
	return types, nil
}
