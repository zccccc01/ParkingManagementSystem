package repository

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/models"

type ViolationRecordRepository interface {
	//创建一条记录
	Create(violation *models.ViolationRecord) error
	//根据记录id获取罚款金额
	GetFineAmountByRecordID(id int) (float64, error)
	//根据记录id获取状态
	GetStatusByRecordID(id int) (string, error)
	//根据记录id获取违章类型
	GetViolationTypeByRecordID(id int) (string, error)
	// 发送罚单通知
	SendViolationNotice(id int) error
}
