package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lordnr/login-register/database"
	"github.com/lordnr/login-register/models"
)

func AllUser(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "User Read Success",
		"users":   users,
	})
}

func Profile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user models.User
	database.DB.First(&user, userId)

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "User Read Success",
		"users":   user,
	})

}
