package main

import (
	"log"

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
	spaceRepo := repository.NewParkingSpaceRepository(db)

	// TODO: where ="unpaid" 找到RecordID 从RID找VehicleID 找人的ID
	// space := models.ParkingSpace{
	// 	SpaceID:      1,
	// 	Status:       "Free",
	// 	ParkingLotID: 1,
	// }
	// ans := spaceRepo.UpdateStatusBySpaceID(&space, 1)
	// if ans != nil {
	// 	log.Fatalf("failed to update status: %v", ans)
	// }

}
