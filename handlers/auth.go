package handlers

import (
	"ToDo/database"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	users := database.Client.Database("project").Collection("users")
	filter := bson.D{{Key: "username", Value: username}}
	var result User
	err := users.FindOne(c, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}
	if !CheckPasswordHash(password, result.Password) {
		return
	}
	session, _ := store.Get(c.Request, "user")
	session.Values["id"] = username
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		ErrorHandler(c.Writer, c.Request, 500)
		return
	}
	c.Redirect(http.StatusFound, "/todo")
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password, _ := HashPassword(c.PostForm("password"))
	project := database.Client.Database("project")
	users := project.Collection("users")
	users.InsertOne(c, User{username, password})
	c.Redirect(http.StatusFound, "/login")
}

func Logout(c *gin.Context) {
	session, _ := store.Get(c.Request, "user")
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
	c.Redirect(http.StatusOK, "/")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
