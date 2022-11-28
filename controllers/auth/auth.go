package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lordnr/login-register/database"
	"github.com/lordnr/login-register/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User
	if c.Bind(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user.Password = string(hash)
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": result.Error,
		})

		return
	}

	if result.RowsAffected > 0 {
		c.JSON(http.StatusCreated, gin.H{
			"status":  "ok",
			"message": "User Create Success",
			"userId":  user.ID,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "User Create Failed",
		})
	}
}

func Login(c *gin.Context) {
	var user models.User
	if c.Bind(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var userExists models.User
	database.DB.Where("Email = ?", user.Email).First(&userExists)
	if userExists.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "User does not Exists",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(userExists.Password),
		[]byte(user.Password),
	)

	if err == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": userExists.ID,
			"exp":    time.Now().Add(time.Hour * 60).Unix(),
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
		fmt.Println(tokenString, err)
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Login Success",
			"token":   tokenString,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Login Failed",
		})
	}

}
