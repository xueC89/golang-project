package utils

import (
	"encoding/json"
	"net/http"
)

// Response 统一API响应结构
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse 返回成功响应
func SuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := Response{
		Success: true,
		Message: message,
		Data:    data,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// ErrorResponse 返回错误响应
func ErrorResponse(w http.ResponseWriter, statusCode int, errorMsg string) {
	response := Response{
		Success: false,
		Error:   errorMsg,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
