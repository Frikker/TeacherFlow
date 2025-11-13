package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	id       int    `json:"id" uri:"id" binding:"required"`
	username string `json:"username" form:"username" uri:"username" binding:"required"`
	email    string `json:"email" form:"email" uri:"email" binding:"required"`
	password string `json:"password" form:"password" binding:"required"`
}

func (handler *Handler) usersList(c *gin.Context) {
	users, err := handler.Model.FindUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"users": users})
	}
}

func (handler *Handler) getUser(c *gin.Context) {
	var params User
	if err := c.BindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	json := make(map[string]string)
	json["id"] = fmt.Sprint(params.id)
	user, err := handler.Model.FindUser(json)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}
