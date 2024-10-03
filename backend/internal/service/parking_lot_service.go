package service

import (
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

type ParkingLotService struct {
	repo repository.ParkingLotRepository
}

func NewParkingLotService(repo repository.ParkingLotRepository) *ParkingLotService {
	return &ParkingLotService{repo: repo}
}

func (s *ParkingLotService) CreateParkingLot(lot *models.ParkingLot) (bool, error) {
	return s.repo.Create(lot)
}

func (s *ParkingLotService) GetParkingLotByID(id int) (*models.ParkingLot, error) {
	return s.repo.FindByID(id)
}

func (s *ParkingLotService) GetAllParkingLots() ([]models.ParkingLot, error) {
	return s.repo.FindAll()
}

func (s *ParkingLotService) UpdateParkingLot(lot *models.ParkingLot, id int) error {
	return s.repo.Update(lot, id)
}

func (s *ParkingLotService) DeleteParkingLot(id int) error {
	return s.repo.Delete(id)
}
