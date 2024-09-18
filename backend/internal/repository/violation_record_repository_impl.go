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

func (r *ViolationRecordRepositoryImpl) Create(violation *models.ViolationRecord) (bool, error) {
	result := r.DB.Create(&violation)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r *ViolationRecordRepositoryImpl) GetFineAmountByRecordID(id int) ([]models.ViolationRecord, error) {
	var violationRecords []models.ViolationRecord
	result := r.DB.Where("RecordID = ?", id).Find(&violationRecords)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	var details []models.ViolationRecord
	for _, record := range violationRecords {
		details = append(details, models.ViolationRecord{
			RecordID:   record.RecordID,
			FineAmount: record.FineAmount,
		})
	}
	return details, nil
}

func (r *ViolationRecordRepositoryImpl) GetStatusByRecordID(id int) ([]models.ViolationRecord, error) {
	var violationRecords []models.ViolationRecord
	result := r.DB.Where("RecordID = ?", id).Find(&violationRecords)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	var details []models.ViolationRecord
	for _, record := range violationRecords {
		details = append(details, models.ViolationRecord{
			RecordID: record.RecordID,
			Status:   record.Status,
		})
	}
	return details, nil
}

func (r *ViolationRecordRepositoryImpl) GetViolationTypeByRecordID(id int) ([]models.ViolationRecord, error) {
	var violationRecords []models.ViolationRecord
	result := r.DB.Where("RecordID = ?", id).Find(&violationRecords)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	var details []models.ViolationRecord
	for _, record := range violationRecords {
		details = append(details, models.ViolationRecord{
			RecordID:      record.RecordID,
			ViolationType: record.ViolationType,
		})
	}
	return details, nil
}

func (r *ViolationRecordRepositoryImpl) FindViolationRecordByUserID(id int) ([]models.ViolationRecord, error) {
	//select * from violationrecord where RecordID in (
	//	select RecordID from parkingrecord where VehicleID in (
	//  	select VehicleID from vehicle where UserID = ?))
	return nil, nil
}
