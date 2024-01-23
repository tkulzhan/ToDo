package handlers

import (
	"ToDo/database"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getUsers(c *gin.Context) []User {
	users := database.Client.Database("project").Collection("users")
	filter := bson.D{}
	cursor, err := users.Find(c, filter)
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return nil
	}
	defer cursor.Close(c)
	var results []User
	for cursor.Next(c) {
		var result User
		if err := cursor.Decode(&result); err != nil {
			ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
			return nil
		}
		results = append(results, result)
	}
	return results
}

func SeacrhUsers(c *gin.Context) []User {
	keyword := c.Query("search")
	users := database.Client.Database("project").Collection("users")
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: "\"" + keyword + "\""}}}}}}
	cursor, err := users.Aggregate(c, mongo.Pipeline{matchStage})
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return nil
	}
	defer cursor.Close(c)
	var results []User
	if err = cursor.All(c, &results); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return nil
	}
	return results
}

func SortUsers(c *gin.Context, v int, keyword string) []User {
	users := database.Client.Database("project").Collection("users")
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{Key: keyword, Value: v}})
	cursor, err := users.Find(c, filter, opts)
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return nil
	}
	defer cursor.Close(c)
	var results []User
	if err := cursor.All(c, &results); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return nil
	}
	return results
}

func DeleteUser(c *gin.Context) {
	users := database.Client.Database("project").Collection("users")
	user := c.Param("user")
	filter := bson.D{{Key: "username", Value: user}}
	_, err := users.DeleteOne(c, filter)
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin")
}

func AdminPage(c *gin.Context) {
	if !isAuth(c) {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	tmpl, err := template.ParseFiles("./templates/admin.html")
	if err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return
	}
	var users []User
	if c.Query("search") != "" {
		users = SeacrhUsers(c)
	} else if c.Query("sort") != "" {
		t := c.Query("sort")
		v := 1
		if t == "desc" {
			v = -1
		}
		users = SortUsers(c, v, "username")
	} else {
		users = getUsers(c)
	}
	if err := tmpl.Execute(c.Writer, users); err != nil {
		ErrorHandler(c.Writer, c.Request, http.StatusInternalServerError, err)
		return
	}
}
