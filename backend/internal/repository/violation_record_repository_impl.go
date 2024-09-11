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
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	var amounts []float64
	for _, record := range violationRecords {
		amounts = append(amounts, record.FineAmount)
	}
	return amounts, nil
}

// TODO:返回值应该是ID对应某个状态(类似一个<k,v>)
func (r *ViolationRecordRepositoryImpl) GetStatusByRecordID(id int) ([]string, error) {
	var violationRecords []models.ViolationRecord
	result := r.DB.Where("RecordID = ?", id).Find(&violationRecords)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	var statuses []string
	for _, RecordID := range violationRecords {
		statuses = append(statuses, RecordID.Status)
	}
	return statuses, nil
}

// TODO:返回值应该是ID对应某个类型(类似一个<k,v>)
func (r *ViolationRecordRepositoryImpl) GetViolationTypeByRecordID(id int) ([]string, error) {
	var violationRecords []models.ViolationRecord
	result := r.DB.Where("RecordID = ?", id).Find(&violationRecords)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	var types []string
	for _, RecordID := range violationRecords {
		types = append(types, RecordID.ViolationType)
	}
	return types, nil
}
