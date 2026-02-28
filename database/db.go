package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var (
	// DB 全局数据库连接池
	DB *sql.DB
)

// InitDB 初始化数据库连接
func InitDB() error {
	// 从beego配置获取数据库配置
	user := beego.AppConfig.String("dbuser")
	if user == "" {
		user = "root"
	}
	password := beego.AppConfig.String("dbpass")
	if password == "" {
		password = "root"
	}
	host := beego.AppConfig.String("dbhost")
	if host == "" {
		host = "localhost"
	}
	port := beego.AppConfig.String("dbport")
	if port == "" {
		port = "3306"
	}
	dbname := beego.AppConfig.String("dbname")
	if dbname == "" {
		dbname = "myblog"
	}

	// 构建DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
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
		id INT(4) AUTO_INCREMENT PRIMARY KEY NOT NULL,
		username VARCHAR(64) NOT NULL,
		password VARCHAR(64) NOT NULL,
		status INT(4),
		createtime TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`

	_, err := ModifyDB(userTableSQL)
	if err != nil {
		return fmt.Errorf("创建用户表失败: %v", err)
	}

	log.Println("数据库表结构初始化完成")
	return nil
}

func ModifyDB(sql string, args ...interface{}) (int64, error) {
	res, err := DB.Exec(sql, args...)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func QueryDB(sql string, args ...interface{}) (*sql.Rows, error) {
	rows, err := DB.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func QueryRowDB(sql string, args ...interface{}) *sql.Row {
	row := DB.QueryRow(sql, args...)
	return row
}
