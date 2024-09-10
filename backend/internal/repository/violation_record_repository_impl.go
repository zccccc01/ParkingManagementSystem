package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

// ViolationRecordRepositoryImpl 实现了 ViolationRecordRepository 接口
type ViolationRecordRepositoryImpl struct {
	DB *gorm.DB
}

// NewViolationRecordRepositoryImpl 是创建 ViolationRecordRepositoryImpl 实例的构造函数
func NewViolationRecordRepositoryImpl(db *gorm.DB) *ViolationRecordRepositoryImpl {
	return &ViolationRecordRepositoryImpl{DB: db}
}

// Create 创建一条违规记录
func (r *ViolationRecordRepositoryImpl) Create(violation *models.ViolationRecord) error {
	return r.DB.Create(violation).Error
}

// GetStatusByRecordID 根据记录ID获取状态
func (r *ViolationRecordRepositoryImpl) GetStatusByRecordID(id int) (string, error) {
	var violation models.ViolationRecord
	result := r.DB.First(&violation, "record_id = ?", id)
	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		return "", gorm.ErrRecordNotFound
	}
	return violation.Status, nil
}

// SendViolationNotice 发送罚单通知
func (r *ViolationRecordRepositoryImpl) SendViolationNotice(id int) error {
	// 这里应该包含发送通知的逻辑，例如发送邮件或短信
	// 以下代码仅为示例，实际发送逻辑需要根据通知服务提供商的API进行实现
	// 假设发送成功后，更新状态为 "通知已发送"
	var violation models.ViolationRecord
	result := r.DB.First(&violation, "violation_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	violation.Status = "通知已发送"
	result = r.DB.Save(&violation)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
