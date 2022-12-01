package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lordnr/login-register/controllers"
)

func UploadRouter(uploadRouter *gin.Engine) {
	uploadRouter.POST("/upload", controllers.Upload)
}
