package repository

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/models"

type ParkingSpaceRepository interface {
	Create(space *models.ParkingSpace) error
	GetLotIDBySpaceID(id int) (int, error)
	GetStatusBySpaceID(id int) (string, error)
	UpdateStatus(space *models.ParkingSpace, id int) error
	Delete(id int) error
}
