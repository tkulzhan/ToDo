package handlers

import (
	"ToDo/database"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDo struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title,omitempty"`
	Category string             `bson:"category,omitempty"`
	Text     string             `bson:"text,omitempty"`
	Due      time.Time          `bson:"due"`
	State    string             `bson:"state,omitempty"`
	Author   primitive.ObjectID `bson:"author,omitempty"`
}

func GetToDoList(c *gin.Context) []ToDo {
	todos := database.Client.Database("project").Collection("todos")
	filter := bson.D{{Key: "author", Value: GetUser().Id}}
	cursor, _ := todos.Find(c, filter)
	var results []ToDo
	for cursor.Next(c) {
		var result ToDo
		err := cursor.Decode(&result)
		if err != nil {
			panic(err)
		}
		results = append(results, result)
	}
	return results
}

func AddToDo(c *gin.Context) {
	todos := database.Client.Database("project").Collection("todos")
	title := c.PostForm("title")
	category := c.PostForm("category")
	text := c.PostForm("text")
	year, _ := strconv.Atoi(c.PostForm("year"))
	month, _ := strconv.Atoi(c.PostForm("month"))
	day, _ := strconv.Atoi(c.PostForm("day"))
	hour, _ := strconv.Atoi(c.PostForm("hour"))
	minute, _ := strconv.Atoi(c.PostForm("minute"))
	loc := time.Now().Location()
	due := time.Date(year, getMonth(month), day, hour, minute, 0, 0, loc)
	_id := primitive.NewObjectID()
	todos.InsertOne(c, ToDo{_id, title, category, text, due, "Not complete", GetUser().Id})
	c.Redirect(http.StatusSeeOther, "/todo")
}

func getMonth(n int) time.Month {
	switch n {
	case 1:
		return time.January
	case 2:
		return time.February
	case 3:
		return time.March
	case 4:
		return time.April
	case 5:
		return time.May
	case 6:
		return time.June
	case 7:
		return time.July
	case 8:
		return time.August
	case 9:
		return time.September
	case 10:
		return time.October
	case 11:
		return time.November
	}
	return time.December
}
