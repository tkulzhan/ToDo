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
	if !isAuth(c) {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	tmpl, err := template.ParseFiles("./templates/todo.html")
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
	var todos []ToDo
	if c.Query("search") == "" {
		todos = GetToDoList(c)
	} else {
		todos = Seacrh(c)
	}
	if err := tmpl.Execute(c.Writer, todos); err != nil {
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
	todo := GetOne(c)
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
	todo := GetOne(c)
	if err := tmpl.Execute(c.Writer, todo); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError)
		return
	}
}
