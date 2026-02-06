package models

import (
	"time"
)

// SystemInfo 系统信息模型
type SystemInfo struct {
	Message   string    `json:"message"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
	Features  []string  `json:"features"`
}
