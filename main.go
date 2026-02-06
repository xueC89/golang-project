package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"frontend-backend/database"
	"frontend-backend/routes"
)

func main() {
	// 初始化数据库连接
	if err := database.InitDB(); err != nil {
		log.Fatalf("数据库初始化失败: %v\n", err)
	}
	// 确保在程序退出时关闭数据库连接
	defer func() {
		if err := database.CloseDB(); err != nil {
			log.Printf("关闭数据库连接失败: %v\n", err)
		}
	}()

	// 设置服务器端口
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 配置路由
	routes.SetupRoutes()

	// 创建HTTP服务器
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      nil,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 在goroutine中启动服务器
	go func() {
		log.Printf("服务器正在监听端口 %s...\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务器启动失败: %v\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在关闭服务器...")

	// 设置5秒的超时时间来关闭服务器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("服务器关闭失败: %v\n", err)
	}

	log.Println("服务器已成功关闭")
}
