package routes

import (
	"frontend-backend/controllers"
	"frontend-backend/middleware"
)

// SetupRoutes 配置所有路由
func SetupRoutes() {
	// 创建根路由组
	root := NewRouter("")

	// 为所有路由添加CORS中间件
	root.Use(middleware.CORSMiddleware)

	// 健康检查路由
	root.GET("/health", controllers.HealthCheckHandler)

	// 配置用户管理路由
	SetupUserRoutes(root)

	// 配置数据管理路由
	SetupDataRoutes(root)
}
