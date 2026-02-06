package routes

import (
	"frontend-backend/controllers"
)

// SetupUserRoutes 配置用户管理相关路由
func SetupUserRoutes(router *Router) {
	// 用户路由组，前缀为/api/users
	userGroup := router.Group("/api/users")

	// GET /api/users 获取用户列表
	userGroup.GET("", controllers.GetUsersHandler)

	// POST /api/users 创建新用户
	userGroup.POST("", controllers.CreateUserHandler)

	// GET /api/users/:id 获取单个用户信息（预留接口）
	// userGroup.GET("/:id", controllers.GetUserHandler)

	// PUT /api/users/:id 更新用户信息（预留接口）
	// userGroup.PUT("/:id", controllers.UpdateUserHandler)

	// DELETE /api/users/:id 删除用户（预留接口）
	// userGroup.DELETE("/:id", controllers.DeleteUserHandler)
}
