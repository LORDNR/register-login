package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lordnr/login-register/controllers"
)

func PostRouter(postRouter *gin.Engine) {
	postRouter.POST("/posts", controllers.NewPost)
	postRouter.GET("/posts", controllers.Posts)
	postRouter.GET("/posts/:id", controllers.PostById)
	postRouter.PATCH("/posts/:id", controllers.PostUpdate)
	postRouter.DELETE("/posts/:id", controllers.PostDelete)
}
