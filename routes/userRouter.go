package routes

import (
	"github.com/gin-gonic/gin"
	UserController "github.com/lordnr/login-register/controllers/user"
	"github.com/lordnr/login-register/middleware"
)

func UserRouter(incomingRoutes *gin.Engine) {
	authorized := incomingRoutes.Group("/users", middleware.JWTAuthen)
	authorized.GET("/alluser", UserController.AllUser)
	authorized.GET("/profile", UserController.Profile)
}
