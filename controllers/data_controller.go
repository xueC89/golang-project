package controllers

import (
	"net/http"
	"time"

	"frontend-backend/models"
	"frontend-backend/utils"
)

// GetDataHandler 获取系统信息和功能特性
func GetDataHandler(w http.ResponseWriter, r *http.Request) {
	// 使用系统信息模型
	systemInfo := models.SystemInfo{
		Message:   "欢迎使用Go后端API",
		Version:   "1.0.0",
		Timestamp: time.Now(),
		Features: []string{
			"用户管理",
			"数据查询",
			"文件上传",
			"认证授权",
		},
	}

	utils.SuccessResponse(w, http.StatusOK, "获取系统信息成功", systemInfo)
}
