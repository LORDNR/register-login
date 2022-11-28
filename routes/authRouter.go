package routes

import (
	"github.com/gin-gonic/gin"
	AuthController "github.com/lordnr/login-register/controllers/auth"
)

func AuthRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/register", AuthController.Register)
	incomingRoutes.POST("/login", AuthController.Login)
}
