package models

import (
	"fmt"
	"frontend-backend/database"
	"time"
)

// User 用户模型
type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Status     int       `json:"status"` // 状态 0:禁用 1:启用
	Createtime time.Time `json:"createtime"`
}

// 插入用户
func InsertUser(user *User) (int64, error) {
	count, err := database.ModifyDB("INSERT INTO users (username, password, status) VALUES (?, ?, ?)", user.Username, user.Password, user.Status)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("SELECT id FROM users %s", con)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

// 按用户名查询用户ID
func QueryUserWithUsername(username string) int {
	return QueryUserWightCon(fmt.Sprintf("WHERE username = '%s'", username))
}

// 按用户名和密码查询用户ID
func QueryUserWithParam(username string, password string) int {
	return QueryUserWightCon(fmt.Sprintf("WHERE username = '%s' AND password = '%s'", username, password))
}
