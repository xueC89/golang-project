package main

import (
	"log"

	"github.com/astaxie/beego"

	"frontend-backend/controllers"
	"frontend-backend/database"
)

func main() {
	// 初始化数据库连接
	if err := database.InitDB(); err != nil {
		log.Fatalf("数据库初始化失败: %v\n", err)
	}

	// 注册路由
	registerRoutes()

	// 启动beego服务器
	beego.Run()
}

// registerRoutes 注册所有路由
func registerRoutes() {
	// 健康检查路由
	beego.Router("/health", &controllers.HealthController{}, "get:Get")

	// 用户管理路由
	beego.Router("/api/users", &controllers.UserController{}, "get:GetUsers;post:CreateUser")

	// 数据管理路由
	beego.Router("/api/data", &controllers.DataController{}, "get:GetData")
}
