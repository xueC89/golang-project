package routes

import (
	"frontend-backend/controllers"
)

// SetupDataRoutes 配置数据管理相关路由
func SetupDataRoutes(router *Router) {
	// 数据路由组，前缀为/api/data
	dataGroup := router.Group("/api/data")

	// GET /api/data 获取系统信息
	dataGroup.GET("", controllers.GetDataHandler)

	// GET /api/data/stats 获取统计数据（预留接口）
	// dataGroup.GET("/stats", controllers.GetStatsHandler)

	// GET /api/data/logs 获取系统日志（预留接口）
	// dataGroup.GET("/logs", controllers.GetLogsHandler)

	// GET /api/data/config 获取配置信息（预留接口）
	// dataGroup.GET("/config", controllers.GetConfigHandler)
}
