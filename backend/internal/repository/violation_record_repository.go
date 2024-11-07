package repository

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/models"

type ViolationRecordRepository interface {
	//创建一条记录
	Create(violation *models.ViolationRecord) (bool, error)
	//根据记录id获取罚款金额
	GetFineAmountByRecordID(id int) ([]models.ViolationRecord, error)
	//根据记录id获取状态
	GetStatusByRecordID(id int) ([]models.ViolationRecord, error)
	//根据记录id获取违章类型
	GetViolationTypeByRecordID(id int) ([]models.ViolationRecord, error)
	//根据UserID查违章记录
	FindViolationRecordByUserID(id int) ([]models.ViolationRecord, error)
	//根据违章类型统计
	StatisticalViolationsByType(t string) (interface{}, error)
}
