package repository

import (
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

// ParkingRecordRepository 定义停车场仓库接口
type ParkingRecordRepository interface {
	Create(record *models.ParkingRecord) error
	FindByID(id int) (*models.ParkingRecord, error)
	FindByVehicleID(vehicleID int) ([]models.ParkingRecord, error)
	FindAll() ([]models.ParkingRecord, error)
	Update(record *models.ParkingRecord, id int) error
	Delete(id int) error
}
