package controllers

import (
	"time"

	"github.com/astaxie/beego"
)

// HealthController 健康检查控制器
type HealthController struct {
	beego.Controller
}

// Get 健康检查处理函数
func (c *HealthController) Get() {
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"message": "ok",
		"data": map[string]interface{}{
			"timestamp": time.Now().Format(time.RFC3339),
		},
	}
	c.Ctx.Output.SetStatus(200)
	c.ServeJSON()
}
