package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// DB 全局数据库连接池
	DB *sql.DB
)

// InitDB 初始化数据库连接
func InitDB() error {
	// 从环境变量获取数据库配置
	user := getEnv("DB_USER", "root")
	password := getEnv("DB_PASSWORD", "root")
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "3306")
	dbname := getEnv("DB_NAME", "frontend_backend")

	// 构建DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("打开数据库连接失败: %v", err)
	}

	// 设置连接池参数
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	// 测试连接
	if err := db.Ping(); err != nil {
		return fmt.Errorf("数据库连接测试失败: %v", err)
	}

	// 设置全局连接池
	DB = db

	log.Println("数据库连接成功")

	// 初始化数据库表结构
	if err := initTables(); err != nil {
		return fmt.Errorf("初始化表结构失败: %v", err)
	}

	return nil
}

// GetDB 获取数据库连接
func GetDB() *sql.DB {
	return DB
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	if DB != nil {
		log.Println("正在关闭数据库连接...")
		return DB.Close()
	}
	return nil
}

// initTables 初始化数据库表结构
func initTables() error {
	// 创建用户表
	userTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`

	_, err := DB.Exec(userTableSQL)
	if err != nil {
		return fmt.Errorf("创建用户表失败: %v", err)
	}

	log.Println("数据库表结构初始化完成")
	return nil
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
