package models

import (
	"chat-service/pkg/middleware/db"
	"context"
	"fmt"
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Age          int       `json:"age"`
	RegisteredAt time.Time `json:"registered_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewUser[T int | string | time.Time](hash map[string]T) (user User, err error) {
	return FindUser(hash)
}

func FindUser[T int | string | time.Time](hash map[string]T) (user User, err error) {
	hashCount := len(hash)
	query := "SELECT * FROM users WHERE"
	for key, value := range hash {
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
		hashCount--
		if len(hash) > 1 && hashCount != 0 {
			query += " AND"
		}
	}
	err = conn.QueryRow(query).Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.RegisteredAt, &user.CreatedAt, &user.UpdatedAt)
	return
}

func FindUsers() (users []User, err error) {
	conn, err := db.NewConn(context.Background())
	if err != nil {
		return users, err
	}
	defer conn.Close()
	rows, err := conn.SelectAll("users")
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.RegisteredAt, &user.CreatedAt, &user.UpdatedAt)
		users = append(users, user)
	}
	return
}

func CreateUser(name string, email string, age int) (err error) {
	conn, err := db.NewConn(context.Background())
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.CreateRow("users(name, email, age)", name, email, age)
	return err
}
