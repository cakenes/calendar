package main

import (
	"database/sql"
	"main/controllers"
	"main/models"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func main() {
	r := gin.Default()

	models.Connect()
	models.Setup()

	defer models.DB.Close()

	r.POST("/register", controllers.CreateUser)
	r.POST("/login", controllers.LoginUser)
	r.GET("/validate", controllers.CheckValidation)

	r.Run()
}
