package models

import (
	"fmt"
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	BirthDate    string    `json:"birth_date"`
	RegisteredAt time.Time `json:"registered_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (m *Model) NewUser(params map[string]string) (user User, err error) {
	return m.FindUser(params)
}

func (m *Model) FindUser(params map[string]string) (user User, err error) {
	paramsCount := len(params)
	query := "SELECT * FROM users WHERE"
	for key, value := range params {
		switch key {
		case "id":
			query += fmt.Sprintf(" id = %v", value)
		case "name":
			query += fmt.Sprintf(" name = '%s'", value)
		case "email":
			query += fmt.Sprintf(" email = '%s'", value)
		case "age":
			query += fmt.Sprintf(" age = %v", value)
		case "registered_at":
			query += fmt.Sprintf(" registered_at = '%v'", value)
		case "created_at":
			query += fmt.Sprintf(" created_at = '%v'", value)
		case "updated_at":
			query += fmt.Sprintf(" updated_at = '%v'", value)
		default:
			return User{}, fmt.Errorf("Unknown key: %s", key)
		}
		paramsCount--
		if len(params) > 1 && paramsCount != 0 {
			query += " AND"
		}
	}
	err = m.conn.QueryRow(query).Scan(&user.ID, &user.Username, &user.Email, &user.BirthDate, &user.RegisteredAt, &user.CreatedAt, &user.UpdatedAt)
	return
}

func (m *Model) FindUsers() (users []User, err error) {
	rows, err := m.conn.SelectAll("users")
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.BirthDate, &user.RegisteredAt, &user.CreatedAt, &user.UpdatedAt)
		users = append(users, user)
	}
	return
}

func (m *Model) CreateUser(name string, email string, age int) (err error) {
	_, err = m.conn.CreateRow("users(name, email, age)", name, email, age)
	return err
}
