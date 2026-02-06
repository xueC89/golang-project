package controllers

import (
	"net/http"
	"time"

	"frontend-backend/utils"
)

// HealthCheckHandler 健康检查处理函数
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	utils.SuccessResponse(w, http.StatusOK, "ok", map[string]interface{}{
		"timestamp": time.Now().Format(time.RFC3339),
	})
}
