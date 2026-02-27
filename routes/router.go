package routes

import (
	"github.com/astaxie/beego"

	"project/controllers"
)

// RouterGroup 路由组
type RouterGroup struct {
	prefix string
}

// NewRouterGroup 创建新的路由组
func NewRouterGroup(prefix string) *RouterGroup {
	return &RouterGroup{
		prefix: prefix,
	}
}

// Register 注册路由
func (g *RouterGroup) Register(path string, controller interface{}, methods string) {
	fullPath := g.prefix + path
	if c, ok := controller.(beego.ControllerInterface); ok {
		beego.Router(fullPath, c, methods)
	} else {
		panic("controller does not implement beego.ControllerInterface")
	}
}

// SetupRoutes 设置所有路由
func SetupRoutes() {
	// 健康检查路由
	beego.Router("/health", &controllers.HealthController{}, "get:Get")

	// API路由组
	apiGroup := NewRouterGroup("/api")

	// 用户管理路由
	apiGroup.Register("/users", &controllers.UserController{}, "get:GetUsers;post:CreateUser")

	// 数据管理路由
	apiGroup.Register("/data", &controllers.DataController{}, "get:GetData")

	// 预留其他路由组
	// adminGroup := NewRouterGroup("/admin")
	// adminGroup.Register("/users", &controllers.AdminController{}, "get:List;post:Create")
}
