package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lordnr/login-register/config"
	"github.com/lordnr/login-register/database"
	"github.com/lordnr/login-register/routes"
)

func init() {
	config.LoadEnvironment()
	database.ConnectDB()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(gin.Logger())

	routes.AuthRouter(r)
	routes.UserRouter(r)

	r.Run()
}