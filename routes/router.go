package routes

import (
	"github.com/astaxie/beego"

	"frontend-backend/controllers"
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
	// API路由组
	apiGroup := NewRouterGroup("/api")

	// 用户管理路由
	// apiGroup.Register("/users", &controllers.UserController{}, "get:GetUsers")

	// 用户注册
	apiGroup.Register("/register", &controllers.UserController{}, "post:CreateUser")

	// 用户登录
	apiGroup.Register("/login", &controllers.UserController{}, "post:Login")

	// 用户登出
	apiGroup.Register("/logout", &controllers.UserController{}, "post:Logout")

	// 获取用户信息
	apiGroup.Register("/user/info", &controllers.UserController{}, "get:GetUserInfo")

	// 文章管理路由
	apiGroup.Register("/article/add", &controllers.ArticleController{}, "post:CreateArticle")

	// 获取文章列表
	apiGroup.Register("/article/list", &controllers.ArticleController{}, "get:GetArticleList")

	// 预留其他路由组
	// adminGroup := NewRouterGroup("/admin")
	// adminGroup.Register("/users", &controllers.AdminController{}, "get:List;post:Create")
}
