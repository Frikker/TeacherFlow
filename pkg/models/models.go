package models

import (
	"chat-service/pkg/middleware/db"
)

type Model struct {
	conn *db.Connection
}

func NewModel(conn *db.Connection) *Model {
	return &Model{conn: conn}
}
