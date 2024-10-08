package controllers

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"

type ViolationRecordController struct {
	ViolationRecordRepo repository.ViolationRecordRepository
}

func NewViolationRecordController(repo repository.ViolationRecordRepository) *ViolationRecordController {
	return &ViolationRecordController{ViolationRecordRepo: repo}
}
