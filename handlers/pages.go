package handlers

import (
	"ToDo/database"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func HomePage(c *gin.Context) {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(c.Writer, ""); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
}

func LoginPage(c *gin.Context) {
	tmpl, err := template.ParseFiles("./templates/login.html")
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(c.Writer, ""); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
}

func RegistrationPage(c *gin.Context) {
	tmpl, err := template.ParseFiles("./templates/register.html")
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(c.Writer, ""); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
}

func ToDoPage(c *gin.Context) {
	if !isAuth(c) {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	tmpl, err := template.ParseFiles("./templates/todo.html")
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(c.Writer, GetToDoList(c)); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
}

func AddToDoPage(c *gin.Context) {
	if !isAuth(c) {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	tmpl, err := template.ParseFiles("./templates/addtodo.html")
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(c.Writer, ""); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
}

func EditPage(c *gin.Context) {
	if !isAuth(c) {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	tmpl, err := template.ParseFiles("./templates/edittodo.html")
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
	todos := database.Client.Database("project").Collection("todos")
	_id := c.Param("id")
	filter := bson.D{{Key: "_id", Value: _id}}
	var todo ToDo
	todos.FindOne(c, filter).Decode(&todo)
	if err := tmpl.Execute(c.Writer, todo); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
}

func ReadPage(c *gin.Context) {
	if !isAuth(c) {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	tmpl, err := template.ParseFiles("./templates/readtodo.html")
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
	todos := database.Client.Database("project").Collection("todos")
	_id := c.Param("id")
	filter := bson.D{{Key: "_id", Value: _id}}
	var todo ToDo
	todos.FindOne(c, filter).Decode(&todo)
	if err := tmpl.Execute(c.Writer, todo); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
}
