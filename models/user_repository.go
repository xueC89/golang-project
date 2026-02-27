package models

import (
	"database/sql"

	"project/database"
)

// UserRepository 用户数据仓库
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository 创建新的用户仓库
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.GetDB(),
	}
}

// GetAllUsers 获取所有用户
func (r *UserRepository) GetAllUsers() ([]User, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM users ORDER BY id ASC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID 根据ID获取用户
func (r *UserRepository) GetUserByID(id int) (*User, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// CreateUser 创建新用户
func (r *UserRepository) CreateUser(name, email string) (*User, error) {
	query := `INSERT INTO users (name, email) VALUES (?, ?)`
	result, err := r.db.Exec(query, name, email)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// 返回创建的用户
	return r.GetUserByID(int(id))
}

// UpdateUser 更新用户信息
func (r *UserRepository) UpdateUser(id int, name, email string) (*User, error) {
	query := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	_, err := r.db.Exec(query, name, email, id)
	if err != nil {
		return nil, err
	}

	// 返回更新后的用户
	return r.GetUserByID(id)
}

// DeleteUser 删除用户
func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
