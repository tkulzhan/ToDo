package main

import (
	"ToDo/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static/")
	router.GET("/", handlers.HomePage)
	router.GET("/login", handlers.LoginPage)
	router.GET("/register", handlers.RegistrationPage)
	router.GET("/todo", handlers.ToDoPage)
	router.POST("register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.POST("/logout", handlers.Logout)
	router.Run("localhost:3000")
}
