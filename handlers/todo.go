package handlers

import (
	"ToDo/database"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ToDo struct {
	Id       string             `bson:"_id,omitempty"`
	Title    string             `bson:"title,omitempty"`
	Category string             `bson:"category,omitempty"`
	Text     string             `bson:"text,omitempty"`
	Due      string             `bson:"due"`
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
	for i := 0; i < len(results); i++ {
		if len(results[i].Text) > 100 {
			results[i].Text = results[i].Text[:100] + "..."
		}
	}
	return results
}

func AddToDo(c *gin.Context) {
	todos := database.Client.Database("project").Collection("todos")
	title := c.PostForm("title")
	category := c.PostForm("category")
	text := c.PostForm("text")
	date := strings.FieldsFunc(c.PostForm("due"), split)
	year, _ := strconv.Atoi(date[0])
	month, _ := strconv.Atoi(date[1])
	day, _ := strconv.Atoi(date[2])
	hour, _ := strconv.Atoi(date[3])
	minute, _ := strconv.Atoi(date[4])
	loc := time.Now().Location()
	due := time.Date(year, getMonth(month), day, hour, minute, 0, 0, loc).Format(time.RFC3339)
	due = due[:16]
	_id := primitive.NewObjectID().Hex()
	todos.InsertOne(c, ToDo{_id, title, category, text, due, "Not complete", GetUser().Id})
	c.Redirect(http.StatusSeeOther, "/todo")
}

func EditToDo(c *gin.Context) {
	todos := database.Client.Database("project").Collection("todos")
	_id := c.Param("id")
	title := c.PostForm("title")
	category := c.PostForm("category")
	text := c.PostForm("text")
	date := strings.FieldsFunc(c.PostForm("due"), split)
	year, _ := strconv.Atoi(date[0])
	month, _ := strconv.Atoi(date[1])
	day, _ := strconv.Atoi(date[2])
	hour, _ := strconv.Atoi(date[3])
	minute, _ := strconv.Atoi(date[4])
	state := c.PostForm("state")
	loc := time.Now().Location()
	due := time.Date(year, getMonth(month), day, hour, minute, 0, 0, loc).Format(time.RFC3339)
	due = due[:16]
	filter := bson.D{{Key: "_id", Value: _id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{{Key: "title", Value: title}}},
		{Key: "$set", Value: bson.D{{Key: "category", Value: category}}},
		{Key: "$set", Value: bson.D{{Key: "text", Value: text}}},
		{Key: "$set", Value: bson.D{{Key: "due", Value: due}}},
		{Key: "$set", Value: bson.D{{Key: "state", Value: state}}},
	}
	_, err := todos.UpdateOne(c, filter, update)
	if err != nil {
		panic(err)
	}
	c.Redirect(http.StatusSeeOther, "/todo")
}

func GetOne(c *gin.Context) ToDo {
	todos := database.Client.Database("project").Collection("todos")
	_id := c.Param("id")
	filter := bson.D{{Key: "_id", Value: _id}}
	var todo ToDo
	todos.FindOne(c, filter).Decode(&todo)
	return todo
}

func DeleteToDo(c *gin.Context) {
	todos := database.Client.Database("project").Collection("todos")
	_id := c.Param("id")
	filter := bson.D{{Key: "_id", Value: _id}}
	todos.DeleteOne(c, filter)
	c.Redirect(http.StatusSeeOther, "/todo")
}

func Seacrh(c *gin.Context) []ToDo {
	keyword := c.Query("search")
	todos := database.Client.Database("project").Collection("todos")
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: "\"" + keyword + "\""}}}}}}
	cursor, err := todos.Aggregate(c, mongo.Pipeline{matchStage})
	if err != nil {
		panic(err)
	}
	var results []ToDo
	if err = cursor.All(c, &results); err != nil {
		panic(err)
	}
	return results
}

func Sort(c *gin.Context) []ToDo {
	keyword := c.Query("sort")
	todos := database.Client.Database("project").Collection("todos")
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{Key: keyword, Value: 1}})
	cursor, err := todos.Find(c, filter, opts)
	var results []ToDo
	if err = cursor.All(c, &results); err != nil {
		panic(err)
	}
	return results
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

func split(r rune) bool {
	return r == ':' || r == '-' || r == 'T'
}
