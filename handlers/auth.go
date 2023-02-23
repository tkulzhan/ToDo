package handlers

import (
	"ToDo/database"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
}

var Store = sessions.NewCookieStore([]byte(os.Getenv(randomString(15))))
var user User

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	users := database.Client.Database("project").Collection("users")
	filter := bson.D{{Key: "username", Value: username}}
	var result User
	err := users.FindOne(c, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.Redirect(http.StatusSeeOther, "/login")
			fmt.Fprint(c.Writer, "err")
			return
		}
		panic(err)
	}
	if !CheckPasswordHash(password, result.Password) {
		c.Redirect(http.StatusSeeOther, "/login")
		fmt.Fprint(c.Writer, "err")
		return
	}
	user = result
	session, _ := Store.Get(c.Request, "user")
	session.Values["id"] = username
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		ErrorHandler(c.Writer, c.Request, 500)
		return
	}
	c.Redirect(http.StatusSeeOther, "/todo")
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password, _ := HashPassword(c.PostForm("password"))
	project := database.Client.Database("project")
	users := project.Collection("users")
	_id := primitive.NewObjectID()
	_, err := users.InsertOne(c, User{_id, username, password})
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/register")
		return
	}
	c.Redirect(http.StatusSeeOther, "/login")
}

func Logout(c *gin.Context) {
	session, _ := Store.Get(c.Request, "user")
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
	c.Redirect(http.StatusSeeOther, "/")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"

func randomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func GetUser() User {
	return user
}

func isAuth(c *gin.Context) bool {
	session, _ := Store.Get(c.Request, "user")
	return session.Values["id"] == GetUser().Username
}
