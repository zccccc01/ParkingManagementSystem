package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type ParkingLotRepositoryImpl struct {
	DB *gorm.DB
}

func NewParkingLotRepository(db *gorm.DB) ParkingLotRepository {
	return &ParkingLotRepositoryImpl{DB: db}
}

func (r *ParkingLotRepositoryImpl) Create(lot *models.ParkingLot) (bool, error) {
	result := r.DB.Create(&lot)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r *ParkingLotRepositoryImpl) FindByID(id int) (*models.ParkingLot, error) {
	var lot models.ParkingLot
	result := r.DB.First(&lot, "ParkingLotID = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &lot, nil
}

func (r *ParkingLotRepositoryImpl) FindByName(name string) (*models.ParkingLot, error) {
	var lot models.ParkingLot
	result := r.DB.First(&lot, "ParkingName = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}
	return &lot, nil
}

func (r *ParkingLotRepositoryImpl) FindAll() ([]models.ParkingLot, error) {
	var lots []models.ParkingLot
	result := r.DB.Find(&lots)
	if result.Error != nil {
		return nil, result.Error
	}
	return lots, nil
}

func (r *ParkingLotRepositoryImpl) Update(lot *models.ParkingLot, id int) error {
	var existingLot models.ParkingLot

	// 使用 Model 和 Updates 方法
	result := r.DB.Model(&existingLot).Where("ParkingLotID = ?", id).Updates(lot)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (r *ParkingLotRepositoryImpl) Delete(id int) error {
	return r.DB.Delete(&models.ParkingLot{}, "ParkingLotID = ?", id).Error
}

func (r *ParkingLotRepositoryImpl) FindAllIncomeByLotID(id int) (float64, error) {
	var totalIncome sql.NullFloat64
	// select sum(Amount) from paymentrecord where RecordID in (select RecordID from parkingrecord where LotID = ?)
	query := `
        SELECT SUM(Amount)
        FROM paymentrecord 
        WHERE RecordID IN (
            SELECT RecordID 
            FROM parkingrecord 
            WHERE LotID = ?
        )
    `
	// 写好sql语句,再调用Raw方法查询,Raw()用于执行原生SQL查询
	result := r.DB.Raw(query, id).Row().Scan(&totalIncome)
	if result != nil {
		return 0, fmt.Errorf("error executing query: %w", result)
	}
	if totalIncome.Valid {
		return totalIncome.Float64, nil
	}
	return 0, nil
}

func (r *ParkingLotRepositoryImpl) FindOccupancyRateByLotID(id int) (float64, error) {
	var occupancyRate sql.NullFloat64
	var totalCapacity int
	var totalOccupied int
	// SELECT COUNT(SpaceID) FROM parkingspace WHERE `Status` != 'Free' AND ParkingLotID = 1)
	result := r.DB.Table("parkingspace").Where("Status != ? AND ParkingLotID = ?", "Free", id).Count(&totalOccupied)
	if result.Error != nil {
		return 0, result.Error
	}
	// SELECT COUNT(SpaceID) FROM parkingspace WHERE ParkingLotID = ?
	result = r.DB.Table("parkingspace").Where("ParkingLotID = ?", id).Count(&totalCapacity)
	if result.Error != nil {
		return 0, result.Error
	}

	if totalCapacity == 0 {
		return 0, nil
	}
	occupancyRate.Float64 = float64(totalOccupied) / float64(totalCapacity)
	return occupancyRate.Float64, nil
}

// 注意时区(要选北京时区)
func (r *ParkingLotRepositoryImpl) FindOccupancyByLotIDAndTime(id int, start time.Time, end time.Time) ([]models.ParkingSpace, error) {
	var parkingSpaces []models.ParkingSpace
	// SELECT * FROM parkingspace WHERE ParkingLotID IN (SELECT LotID FROM parkingrecord WHERE StartTime > '?' AND EndTime < '?')
	query := `
		SELECT * 
        FROM parkingspace 
        WHERE ParkingLotID IN (
            SELECT LotID 
            FROM parkingrecord 
            WHERE StartTime > ? AND EndTime < ?
        )`
	result := r.DB.Raw(query, start, end).Scan(&parkingSpaces)
	if result.Error != nil {
		return nil, result.Error
	}
	return parkingSpaces, nil
}

func (r *ParkingLotRepositoryImpl) GetFreeSpaceByLotID(id int) (int, error) {
	var freeSpace int
	// SELECT COUNT(SpaceID) FROM parkingspace WHERE Status = 'Free' AND ParkingLotID = ?
	result := r.DB.Table("parkingspace").Where("Status = ? AND ParkingLotID = ?", "Free", id).Count(&freeSpace)
	if result.Error != nil {
		return 0, result.Error
	}
	return freeSpace, nil

}

func (r *ParkingLotRepositoryImpl) GetOccupiedSpaceByLotID(id int) (int, error) {
	var occupiedSpace int
	// SELECT COUNT(SpaceID) FROM parkingspace WHERE Status = 'Occupied' AND ParkingLotID = ?
	result := r.DB.Table("parkingspace").Where("Status = ? AND ParkingLotID = ?", "Occupied", id).Count(&occupiedSpace)
	if result.Error != nil {
		return 0, result.Error
	}
	return occupiedSpace, nil
}

func (r *ParkingLotRepositoryImpl) GetReservedSpaceByLotID(id int) (int, error) {
	var reservedSpace int
	// SELECT COUNT(SpaceID) FROM parkingspace WHERE Status = 'Reserved' AND ParkingLotID = ?
	result := r.DB.Table("parkingspace").Where("Status = ? AND ParkingLotID = ?", "Reserved", id).Count(&reservedSpace)
	if result.Error != nil {
		return 0, result.Error
	}
	return reservedSpace, nil
}
