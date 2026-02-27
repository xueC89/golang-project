package main

import (
	"log"

	"github.com/astaxie/beego"

	"project/database"
	"project/routes"
)

func main() {
	// 初始化数据库连接
	if err := database.InitDB(); err != nil {
		log.Fatalf("数据库初始化失败: %v\n", err)
	}

	// 注册路由
	routes.SetupRoutes()

	// 启动beego服务器
	beego.Run()
}
