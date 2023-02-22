package handlers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
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
	tmpl, err := template.ParseFiles("./templates/todo.html")
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(c.Writer, ""); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
}
