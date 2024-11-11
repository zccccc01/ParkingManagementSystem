package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/config"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/routes"
)

func main() {
	// 创建 Fiber 实例
	app := fiber.New()

	db := config.GetDBInstance()
	rdb := config.GetRDBInstance()

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
	log.Fatal(app.Listen("0.0.0.0:8000"))
}
