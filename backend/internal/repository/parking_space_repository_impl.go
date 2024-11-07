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

func (r *ParkingSpaceRepositoryImpl) GetStatusByLotIDAndSpaceID(lotID int, spaceID int) (string, error) {
	var space models.ParkingSpace
	result := r.DB.First(&space, "ParkingLotID = ? and SpaceID = ?", lotID, spaceID)
	if result.Error != nil {
		return "", result.Error
	}
	return space.Status, nil
}

func (r *ParkingSpaceRepositoryImpl) UpdateStatusBySpaceID(space *models.ParkingSpace, lotID int, spaceID int) (bool, error) {
	var existingSpace models.ParkingSpace
	result := r.DB.Model(&existingSpace).Where("SpaceID = ? and ParkingLotID = ?", lotID, spaceID).Updates(space)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r *ParkingSpaceRepositoryImpl) FindVehicleSpaceInLotByPlateNumber(plateNumber string) (interface{}, error) {
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

	return space, nil
}

func (r *ParkingSpaceRepositoryImpl) FindVehicleSpaceInLotByUserID(id int) (interface{}, error) {
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

	return spaces, nil
}

func (r *ParkingSpaceRepositoryImpl) FindFreeSpaceInAllLots() ([]models.ParkingSpace, error) {
	// select lotID,SpaceID from parkingspace where Status = "FREE"
	var spaces []models.ParkingSpace
	result := r.DB.Where("Status = ?", "FREE").Find(&spaces)
	if result.Error != nil {
		return nil, result.Error
	}
	return spaces, nil
}
