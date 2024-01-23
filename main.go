package main

import (
	"ToDo/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Println("Error loading app.env: ", err)
	}
}

func main() {
	// Setup Gin router
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
	router.GET("/sort", handlers.ToDoPage)
	// Admin
	router.GET("/admin/search", handlers.AdminPage)
	router.GET("/admin/sort", handlers.AdminPage)
	router.GET("/admin", handlers.AdminPage)
	router.GET("admin/delete/:user", handlers.DeleteUser)
	port := GetEnv("PORT", "3000")
	router.Run(":" + port)
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Println("Could not find " + key + " in env. Returning fallback")
	return fallback
}
