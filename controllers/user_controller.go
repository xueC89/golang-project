package controllers

import (
	"encoding/json"
	"net/http"

	"frontend-backend/models"
	"frontend-backend/utils"
)

// CreateUserRequest 创建用户请求结构
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

// GetUsersHandler 获取所有用户信息
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// 创建用户仓库
	repo := models.NewUserRepository()

	// 从数据库获取所有用户
	users, err := repo.GetAllUsers()
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "获取用户列表失败: "+err.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "获取用户列表成功", map[string]interface{}{
		"users": users,
	})
}

// CreateUserHandler 创建新用户
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// 限制请求体大小为1MB
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	// 解析请求体
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "无效的请求数据: "+err.Error())
		return
	}

	// 验证请求数据
	if req.Name == "" || req.Email == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "姓名和邮箱不能为空")
		return
	}

	// 创建用户仓库
	repo := models.NewUserRepository()

	// 创建新用户
	user, err := repo.CreateUser(req.Name, req.Email)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "创建用户失败: "+err.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusCreated, "用户创建成功", map[string]interface{}{
		"user": user,
	})
}
