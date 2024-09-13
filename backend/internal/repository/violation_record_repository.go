package repository

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/models"

// TODO: where ="unpaid" 找到RecordID 从RID找VehicleID 找人的ID
// 实现发违规记录给某人

type ViolationRecordRepository interface {
	//创建一条记录
	Create(violation *models.ViolationRecord) error
	//根据记录id获取罚款金额
	GetFineAmountByRecordID(id int) ([]models.ViolationRecord, error)
	//根据记录id获取状态
	GetStatusByRecordID(id int) ([]models.ViolationRecord, error)
	//根据记录id获取违章类型
	GetViolationTypeByRecordID(id int) ([]models.ViolationRecord, error)
}
