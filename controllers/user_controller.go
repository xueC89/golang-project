package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"

	"frontend-backend/models"
	"frontend-backend/utils"
)

// UserController 用户控制器
type UserController struct {
	beego.Controller
}

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateUserRequest 创建用户请求结构
type CreateUserRequest struct {
	UserRequest
}

type LoginRequest struct {
	UserRequest
}

// CreateUser 创建新用户
func (c *UserController) CreateUser() {
	// 解析请求体
	var req CreateUserRequest
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&req); err != nil {
		c.Data["json"] = map[string]interface{}{
			"code":    0,
			"success": false,
			"message": "无效的请求数据: " + err.Error(),
		}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}

	// 验证请求数据
	if req.Username == "" || req.Password == "" {
		c.Data["json"] = map[string]interface{}{
			"code":    0,
			"success": false,
			"message": "用户名和密码不能为空",
		}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}

	// 验证用户名是否已存在
	if models.QueryUserWithUsername(req.Username) != 0 {
		c.Data["json"] = map[string]interface{}{
			"code":    0,
			"success": false,
			"message": "用户名已存在",
		}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}

	// 创建新用户
	_, err := models.InsertUser(&models.User{
		Username: req.Username,
		Password: utils.MD5(req.Password),
		Status:   1,
	})
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code":    0,
			"success": false,
			"message": "创建用户失败: " + err.Error(),
		}
		c.Ctx.Output.SetStatus(500)
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"code":    1,
		"success": true,
		"message": "用户创建成功",
		"data":    nil,
	}
	c.Ctx.Output.SetStatus(200)
	c.ServeJSON()
}

// Login 用户登录
func (c *UserController) Login() {
	// 解析请求体
	var req LoginRequest
	// json.NewDecoder(c.Ctx.Request.Body).Decode(&req)
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&req); err != nil {
		c.Data["json"] = map[string]interface{}{
			"code":    0,
			"success": false,
			"message": "无效的请求数据: " + err.Error(),
		}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	fmt.Println(req.Username)

	// 验证请求数据
	if req.Username == "" || req.Password == "" {
		c.Data["json"] = map[string]interface{}{
			"code":    0,
			"success": false,
			"message": "用户名或密码不能为空",
		}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}

	id := models.QueryUserWithParam(req.Username, utils.MD5(req.Password))
	// 验证用户是否存在
	if id == 0 {
		c.Data["json"] = map[string]interface{}{
			"code":    0,
			"success": false,
			"message": "用户名或密码错误",
		}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	} else {
		c.Data["json"] = map[string]interface{}{
			"code":    1,
			"success": true,
			"message": "登录成功",
			"data":    nil,
		}
		c.Ctx.Output.SetStatus(200)
		c.ServeJSON()
		return
	}
}
