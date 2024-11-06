package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/config"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/routes"
)

func main() {
	// 创建 Fiber 实例
	app := fiber.New()

	db := config.GetDBInstance()
	rdb := config.GetRDBInstance()

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
	routes.SetupCountRoutes(app, rdb)
	// 启动服务器
	log.Fatal(app.Listen(":8000"))
}
