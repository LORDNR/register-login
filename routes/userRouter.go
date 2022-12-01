package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lordnr/login-register/controllers"
	"github.com/lordnr/login-register/middleware"
)

func UserRouter(UserRouter *gin.Engine) {
	authorized := UserRouter.Group("/users", middleware.JWTAuthen)
	authorized.GET("/alluser", controllers.AllUser)
	authorized.GET("/profile", controllers.Profile)
}
