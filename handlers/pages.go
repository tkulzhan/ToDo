package handlers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return
	}
	if err := tmpl.Execute(c.Writer, ""); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return
	}
}

func LoginPage(c *gin.Context) {
	tmpl, err := template.ParseFiles("./templates/login.html")
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return
	}
	if err := tmpl.Execute(c.Writer, ""); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return
	}
}

func RegistrationPage(c *gin.Context) {
	tmpl, err := template.ParseFiles("./templates/register.html")
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return
	}
	if err := tmpl.Execute(c.Writer, ""); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
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
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return
	}
	var todos []ToDo
	if c.Query("search") != "" {
		todos = Seacrh(c)
	} else if c.Query("sort") != "" {
		k := c.Query("sort")
		t := c.Query("sortType")
		v := 1
		if t == "desc" {
			v = -1
		}
		todos = Sort(c, v, k)
	}else {
		todos = GetToDoList(c)
	}
	if err := tmpl.Execute(c.Writer, todos); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
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
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return
	}
	if err := tmpl.Execute(c.Writer, ""); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
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
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return
	}
	todo := GetOne(c)
	if err := tmpl.Execute(c.Writer, todo); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
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
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return
	}
	todo := GetOne(c)
	if err := tmpl.Execute(c.Writer, todo); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return
	}
}
