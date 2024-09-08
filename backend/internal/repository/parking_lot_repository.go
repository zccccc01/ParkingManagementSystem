package repository

import (
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

// ParkingLotRepository 定义停车场仓库接口
type ParkingLotRepository interface {
	Create(lot *models.ParkingLot) error
	FindByID(id int) (*models.ParkingLot, error)
	FindByName(name string) (*models.ParkingLot, error)
	FindAll() ([]models.ParkingLot, error)
	Update(lot *models.ParkingLot, id int) error
	Delete(id int) error
}
