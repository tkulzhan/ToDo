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
	router.GET("/add", handlers.AddToDoPage)
	router.GET("/read/:id", handlers.ReadPage)
	router.GET("/edit/:id", handlers.EditPage)
	router.POST("/edit/:id", handlers.EditToDo)
	router.GET("/delete/:id", handlers.DeleteToDo)
	router.POST("register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.GET("/logout", handlers.Logout)
	router.POST("/add", handlers.AddToDo)
	router.Run("localhost:3000")
}
