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
	var violationRecords []models.ViolationRecord
	query := `
		SELECT *
		FROM violationrecord 
		WHERE RecordID IN (
			SELECT RecordID
			FROM parkingrecord 
			WHERE VehicleID IN (
				SELECT VehicleID
				FROM vehicle 
				WHERE UserID = ?
			)
		)`
	result := r.DB.Raw(query, id).Scan(&violationRecords)
	if result.Error != nil {
		return nil, result.Error
	}

	return violationRecords, nil

}

func (r *ViolationRecordRepositoryImpl) StatisticalViolationsByType(t string) (interface{}, error) {
	query := `SELECT ViolationType, Status, COUNT(*) AS TotalViolations, 
	SUM(CASE WHEN Status = 'PAID' THEN 1 ELSE 0 END) AS PaidCount,
	SUM(CASE WHEN Status = 'UNPAID' THEN 1 ELSE 0 END) AS UnpaidCount,
	SUM(CASE WHEN Status = 'DISPUTED' THEN 1 ELSE 0 END) AS DisputedCount,
	SUM(FineAmount) AS TotalFineAmount
	FROM violationrecord WHERE ViolationType = ? GROUP BY ViolationType, Status
	ORDER BY ViolationType, Status;
	`
	var report []struct {
		ViolationType   string  `gorm:"column:ViolationType"`
		Status          string  `gorm:"column:Status"`
		TotalViolations int     `gorm:"column:TotalViolations"`
		PaidCount       int     `gorm:"column:PaidCount"`
		UnpaidCount     int     `gorm:"column:UnpaidCount"`
		DisputedCount   int     `gorm:"column:DisputedCount"`
		TotalFineAmount float64 `gorm:"column:TotalFineAmount"`
	}
	result := r.DB.Raw(query, t).Scan(&report)
	if result.Error != nil {
		return nil, result.Error
	}

	for i := range report {
		report[i].TotalFineAmount = float64(int(report[i].TotalFineAmount*100)) / 100
	}

	return report, nil
}
