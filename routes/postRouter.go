package routes

import (
	"github.com/gin-gonic/gin"
	PostController "github.com/lordnr/login-register/controllers/post"
)

func PostRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/posts", PostController.NewPost)
	incomingRoutes.GET("/posts", PostController.Posts)
	incomingRoutes.GET("/posts/:id", PostController.PostById)
	incomingRoutes.PATCH("/posts/:id", PostController.PostUpdate)
}
