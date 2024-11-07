package repository

import (
	"time"

	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type ReservationRepository interface {
	//创建一条记录
	Create(reservation *models.Reservation) (bool, error)
	//根据预定id更新状态
	UpdateStatusByReservationID(id int, status string) error
	//根据预定id更新预定
	UpdateByReservationID(id int, reservation *models.Reservation) error
	//根据预定id删除记录
	DeleteByReservationID(id int) error
	//根据预定的停车场id获取费用
	GetFeeByLotIDAndTime(id int, start time.Time, end time.Time) (float64, error)
}
