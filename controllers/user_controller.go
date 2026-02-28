package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"frontend-backend/models"
	"frontend-backend/utils"
)

// UserRequest 用户请求基础结构
type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateUserRequest 创建用户请求结构
type CreateUserRequest struct {
	UserRequest
}

// LoginRequest 登录请求结构
type LoginRequest struct {
	UserRequest
}

// UserController 用户控制器
type UserController struct {
	beego.Controller
}

// CreateUser 创建新用户
func (c *UserController) CreateUser() {
	// 解析请求体
	var req CreateUserRequest
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&req); err != nil {
		c.jsonError(400, "无效的请求数据: "+err.Error())
		return
	}

	// 验证请求数据
	if req.Username == "" || req.Password == "" {
		c.jsonError(400, "用户名和密码不能为空")
		return
	}

	// 验证用户名是否已存在
	if models.QueryUserWithUsername(req.Username) != 0 {
		c.jsonError(400, "用户名已存在")
		return
	}

	// 创建新用户
	_, err := models.InsertUser(&models.User{
		Username: req.Username,
		Password: utils.MD5(req.Password),
		Status:   1,
	})
	if err != nil {
		c.jsonError(500, "创建用户失败: "+err.Error())
		return
	}

	c.jsonSuccess(200, "用户创建成功", nil)
}

// Login 用户登录
func (c *UserController) Login() {
	// 解析请求体
	var req LoginRequest
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&req); err != nil {
		c.jsonError(400, "无效的请求数据: "+err.Error())
		return
	}

	// 验证请求数据
	if req.Username == "" || req.Password == "" {
		c.jsonError(400, "用户名或密码不能为空")
		return
	}

	// 验证用户凭据
	id := models.QueryUserWithParam(req.Username, utils.MD5(req.Password))
	if id == 0 {
		c.jsonError(400, "用户名或密码错误")
		return
	}

	c.jsonSuccess(200, "登录成功", nil)
}

// jsonError 返回错误JSON响应
func (c *UserController) jsonError(statusCode int, message string) {
	c.Data["json"] = map[string]interface{}{
		"code":    0,
		"success": false,
		"message": message,
		"data":    nil,
	}
	c.Ctx.Output.SetStatus(statusCode)
	c.ServeJSON()
}

// jsonSuccess 返回成功JSON响应
func (c *UserController) jsonSuccess(statusCode int, message string, data interface{}) {
	c.Data["json"] = map[string]interface{}{
		"code":    1,
		"success": true,
		"message": message,
		"data":    data,
	}
	c.Ctx.Output.SetStatus(statusCode)
	c.ServeJSON()
}
