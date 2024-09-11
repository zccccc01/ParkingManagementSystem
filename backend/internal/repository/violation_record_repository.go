package repository

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/models"

type ViolationRecordRepository interface {
	//创建一条记录
	Create(violation *models.ViolationRecord) error
	//根据记录id获取罚款金额  TODO:相同的记录id,可能存在多条记录,return []float64
	GetFineAmountByRecordID(id int) (float64, error)
	//根据记录id获取状态  TODO:相同的记录id,可能存在多条记录,return []string
	GetStatusByRecordID(id int) (string, error)
	//根据记录id获取违章类型  TODO:相同的记录id,可能存在多条记录,return []string
	GetViolationTypeByRecordID(id int) (string, error)
}
