package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type ParkingSpaceRepositoryImpl struct {
	DB *gorm.DB
}

func NewParkingSpaceRepository(db *gorm.DB) ParkingSpaceRepository {
	return &ParkingSpaceRepositoryImpl{DB: db}
}

func (r *ParkingSpaceRepositoryImpl) Create(space *models.ParkingSpace) (bool, error) {
	result := r.DB.Create(&space)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r *ParkingSpaceRepositoryImpl) GetAllStatusByLotID(id int) ([]models.ParkingSpace, error) {
	var spaces []models.ParkingSpace
	result := r.DB.Find(&spaces, "ParkingLotID = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return spaces, nil
}

func (r *ParkingSpaceRepositoryImpl) GetStatusBySpaceID(id int) (string, error) {
	var space models.ParkingSpace
	result := r.DB.First(&space, "SpaceID = id", id)
	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		return "", gorm.ErrRecordNotFound
	}
	return space.Status, nil
}

func (r *ParkingSpaceRepositoryImpl) UpdateStatusBySpaceID(space *models.ParkingSpace, id int) (bool, error) {
	var existingSpace models.ParkingSpace
	result := r.DB.Model(&existingSpace).Where("SpaceID = ?", id).Updates(space)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, gorm.ErrRecordNotFound
	}
	return true, nil
}

func (r *ParkingSpaceRepositoryImpl) FindVehicleSpaceInLotByPlateNumber(plateNumber string) (map[int]int, error) {
	var space []struct {
		LotID   int `gorm:"column:LotID"`
		SpaceID int `gorm:"column:SpaceID"`
	}
	// select LotID, SpaceID from parkingrecord where VehicleID in (select VehicleID from vehicle where plateNumber = ?)
	query := `
			SELECT LotID, SpaceID
			FROM parkingrecord
			WHERE VehicleID IN (
				SELECT VehicleID
				FROM vehicle
				WHERE plateNumber = ?
			)
	`
	result := r.DB.Raw(query, plateNumber).Scan(&space)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	var ans = make(map[int]int)
	for _, record := range space {
		ans[record.LotID] = record.SpaceID
	}
	return ans, nil
}

func (r *ParkingSpaceRepositoryImpl) FindVehicleSpaceInLotByUserID(id int) (map[int]int, error) {
	var spaces []struct {
		LotID   int `gorm:"column:LotID"`
		SpaceID int `gorm:"column:SpaceID"`
	}
	// select LotID, SpaceID from parkingrecord where VehicleID in (select VehicleID from vehicle where UserID = ?)
	query := `
			SELECT LotID, SpaceID
			FROM parkingrecord
			WHERE VehicleID IN (
				SELECT VehicleID
				FROM vehicle
				WHERE UserID = ?
			)
	`
	result := r.DB.Raw(query, id).Scan(&spaces)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	var ans = make(map[int]int)
	for _, record := range spaces {
		ans[record.LotID] = record.SpaceID
	}
	return ans, nil
}

func (r *ParkingSpaceRepositoryImpl) FindFreeSpaceInLot(id int) ([][]int, error) {
	// select lotID,SpaceID from parkingspace where Status = "FREE"
	return nil, nil
}
