package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lordnr/login-register/controllers"
)

func AuthRouter(authRouter *gin.Engine) {
	authRouter.POST("/register", controllers.Register)
	authRouter.POST("/login", controllers.Login)
}
