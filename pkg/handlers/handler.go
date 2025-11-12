package handlers

import (
	"chat-service/pkg/middleware/db"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	conn *db.Connection
}

func NewHandler(conn *db.Connection) *Handler {
	return &Handler{conn}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", handler.signUp)
		auth.POST("/sign-in", handler.signIn)
	}

	users := r.Group("/users")
	{
		users.GET("/", handler.usersList)
		users.GET("/:id", handler.getUser)
	}

	return r
}
