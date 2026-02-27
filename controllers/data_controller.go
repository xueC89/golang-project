package controllers

import (
	"time"

	"github.com/astaxie/beego"

	"project/models"
)

// DataController 数据控制器
type DataController struct {
	beego.Controller
}

// GetData 获取系统信息和功能特性
func (c *DataController) GetData() {
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

	c.Data["json"] = map[string]interface{}{
		"success": true,
		"message": "获取系统信息成功",
		"data":    systemInfo,
	}
	c.Ctx.Output.SetStatus(200)
	c.ServeJSON()
}
