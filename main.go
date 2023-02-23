package main

import (
	"ToDo/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static/")
	router.GET("/", handlers.HomePage)
	// Authentication and Authorization
	router.GET("/logout", handlers.Logout)
	router.GET("/login", handlers.LoginPage)
	router.GET("/register", handlers.RegistrationPage)
	router.POST("register", handlers.Register)
	router.POST("/login", handlers.Login)
	// ToDo operations
	router.GET("/todo", handlers.ToDoPage)
	router.GET("/add", handlers.AddToDoPage)
	router.GET("/read/:id", handlers.ReadPage)
	router.GET("/edit/:id", handlers.EditPage)
	router.GET("/delete/:id", handlers.DeleteToDo)
	router.POST("/add", handlers.AddToDo)
	router.POST("/edit/:id", handlers.EditToDo)
	// Seacrh, Group, Sort
	router.GET("/search", handlers.ToDoPage)
	router.Run("localhost:3000")
}
