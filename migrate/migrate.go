package main

import (
	"github.com/lordnr/login-register/config"
	"github.com/lordnr/login-register/database"
	"github.com/lordnr/login-register/models"
)

func init() {
	config.LoadEnvironment()
	database.ConnectDB()
}

func main() {
	database.DB.AutoMigrate(&models.User{})
}
