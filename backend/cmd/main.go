package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/routes"
)

// TODO:有些函数的返回值,*model,[]*,要不要改原生model
// TODO:检查数据库的实现算是code review,rowsaffected那个要改
// TODO:create函数下都两个返回值吧(bool, error)
// TODO:update函数都检查传入的元素是否改变吧,不改变保留原始值

func main() {
	// 创建 Fiber 实例
	app := fiber.New()

	// 连接数据库
	db, err := gorm.Open("mysql", "root:123456@/chao_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed to connect database: %v", err) // 使用日志记录错误，而不是panic
	}
	// 开db的debug模式
	db = db.Debug()

	defer db.Close()

	// MVC
	// Get取资源,Post新建资源,Put更新资源,Delete删除资源

	// repo := repository.NewParkingLotRepository(db)
	// fee, err := repo.FindByID(101)
	// if err != nil {
	// 	log.Fatalf("failed to connect database: %v", err)
	// }
	// log.Println(fee)

	// 设置 CORS 中间件 允许前端3000端口
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",                       // 允许来自前端的请求
		AllowMethods:     "GET,POST,PUT,DELETE",                         // 允许的 HTTP 方法
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization", // 允许的头部
		AllowCredentials: true,                                          // 允许发送凭证
	}))

	// 设置路由
	routes.SetupParkingLotRoutes(app, db)
	routes.SetupUserRoutes(app, db)
	routes.SetupParkingRecordRoutes(app, db)
	routes.SetupParkingSpaceRoutes(app, db)
	routes.SetupPaymentRecordRoutes(app, db)
	routes.SetupReservationRoutes(app, db)
	routes.SetupVehicleRoutes(app, db)
	routes.SetupViolationRecordRoutes(app, db)
	// 启动服务器
	log.Fatal(app.Listen(":8000"))
}
