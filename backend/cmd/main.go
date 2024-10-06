package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/routes"
)

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

	// 设置路由
	routes.SetupParkingLotRoutes(app, db)
	routes.SetupUserRoutes(app, db)

	// 启动服务器
	log.Fatal(app.Listen(":8000"))
}
