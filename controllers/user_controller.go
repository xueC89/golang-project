package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"frontend-backend/models"
)

// UserController 用户控制器
type UserController struct {
	beego.Controller
}

// CreateUserRequest 创建用户请求结构
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

// GetUsers 获取所有用户信息
func (c *UserController) GetUsers() {
	// 创建用户仓库
	repo := models.NewUserRepository()

	// 从数据库获取所有用户
	users, err := repo.GetAllUsers()
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"error":   "获取用户列表失败: " + err.Error(),
		}
		c.Ctx.Output.SetStatus(500)
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"success": true,
		"message": "获取用户列表成功",
		"data": map[string]interface{}{
			"users": users,
		},
	}
	c.Ctx.Output.SetStatus(200)
	c.ServeJSON()
}

// CreateUser 创建新用户
func (c *UserController) CreateUser() {
	// 解析请求体
	var req CreateUserRequest
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&req); err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"error":   "无效的请求数据: " + err.Error(),
		}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}

	// 验证请求数据
	if req.Name == "" || req.Email == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"error":   "姓名和邮箱不能为空",
		}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}

	// 创建用户仓库
	repo := models.NewUserRepository()

	// 创建新用户
	user, err := repo.CreateUser(req.Name, req.Email)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"error":   "创建用户失败: " + err.Error(),
		}
		c.Ctx.Output.SetStatus(500)
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"success": true,
		"message": "用户创建成功",
		"data": map[string]interface{}{
			"user": user,
		},
	}
	c.Ctx.Output.SetStatus(201)
	c.ServeJSON()
}
