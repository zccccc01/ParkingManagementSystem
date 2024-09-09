package repository

import "github.com/zccccc01/ParkingManagementSystem/backend/internal/models"

type ReservationRepository interface {
	//创建一条预定记录
	Create(reservation *models.Reservation) error
	//根据预定id更新状态
	UpdateStatusByReservationID(id int, status string) error
	//根据预定id删除记录
	DeleteByReservationID(id int) error
	// 取消预定
	CancelReservation(id int) error
	// 处理支付
	ProcessPayment(reservation *models.Reservation) error
	// 查找预定记录
	FindByReservationID(id int) (*models.Reservation, error)
}
