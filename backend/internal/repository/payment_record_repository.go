package repository

import (
	"time"

	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type PaymentRecordRepository interface {
	//创建一条记录
	Create(payment *models.PaymentRecord) error
	//根据记录ID获取金额
	GetAmountByRecordID(id int) (float64, error)
	//根据预约ID获取金额
	GetAmountByReservationID(id int) (float64, error)
	//根据支付ID获取支付方式
	GetPaymentMethodByPaymentID(id int) (string, error)
	//根据支付ID获取支付时间戳
	GetPaymentTimeStampByPaymentID(id int) (time.Time, error)
	//根据支付时间戳获取支付状态
	//TODO:感觉有点怪,我这里的逻辑是查到传入的时间戳对应信息就return pay, 没有查到就nopay
	GetPaymentStatusByPaymentTimeStamp(timestamp time.Time) (string, error)
}
