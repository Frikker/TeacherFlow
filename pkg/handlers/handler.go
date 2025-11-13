package handlers

import (
	"chat-service/pkg/middleware/db"
	"chat-service/pkg/models"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Model *models.Model
}

func NewHandler(conn *db.Connection) *Handler {
	model := models.NewModel(conn)
	return &Handler{Model: model}
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

	r.GET("/ws", handler.handleNotifications)

	return r
}
