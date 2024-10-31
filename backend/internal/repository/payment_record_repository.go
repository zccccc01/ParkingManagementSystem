package repository

import (
	"time"

	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type PaymentRecordRepository interface {
	//创建一条记录
	Create(payment *models.PaymentRecord) (bool, error)
	//根据记录ID获取金额
	GetAmountByRecordID(id int) (float64, error)
	//根据预约ID获取金额
	GetAmountByReservationID(id int) (float64, error)
	//根据支付ID获取支付方式
	GetPaymentMethodByPaymentID(id int) (string, error)
	//根据支付ID获取支付时间戳
	GetPaymentTimeStampByPaymentID(id int) (time.Time, error)
	//根据预定ID获取支付状态
	GetPaymentStatusByReservationID(id int) (string, error)
	//根据记录ID获取支付状态
	GetPaymentStatusByRecordID(id int) (string, error)
	//根据车牌号查支付车费
	GetPaymentFeeByPlateNumber(plateNumber string) ([]float64, error)
	//根据车牌号查支付信息
	GetPaymentInfoByPlateNumber(plateNumber string) ([]models.PaymentRecord, error)
}
