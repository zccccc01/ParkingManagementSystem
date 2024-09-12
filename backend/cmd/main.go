package main

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
)

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql", "root:123456@/chao_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed to connect database: %v", err) // 使用日志记录错误，而不是panic
	}
	// 开db的debug模式
	db = db.Debug()
	// 实例一个接口
	paymentRepo := repository.NewPaymentRecordRepository(db)

	// TODO: where ="unpaid" 找到RecordID 从RID找VehicleID 找人的ID
	// 实现发违规记录给某人

	//timeStamp, _ := paymentRepo.GetPaymentTimeStampByPaymentID(1)

	ans, err := paymentRepo.GetPaymentStatusByPaymentTimeStamp(time.Now())
	if err != nil {
		log.Fatalf("failed to get payment status: %v", err)
	}
	log.Printf("payment status: %v", ans)

}
