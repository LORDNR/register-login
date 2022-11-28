package routes

import (
	"github.com/gin-gonic/gin"
	UploadController "github.com/lordnr/login-register/controllers/upload"
)

func UploadRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/upload", UploadController.Upload)
}
