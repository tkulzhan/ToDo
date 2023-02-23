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
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username,omitempty"`
	Password  string             `bson:"password,omitempty"`
	TimeStamp timestamp          `bson:"timestamp,omitempty"`
	Role      string             `bson:"role,omitempty"`
}

type timestamp struct {
	Start   time.Time `bson:"start,omitempty"`
	Last    time.Time `bson:"last,omitempty"`
	VisitsN int       `bson:"visits_n,omitempty"`
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
	updateTimeStamp(c, users)
	if user.Role == "admin" {
		c.Redirect(http.StatusSeeOther, "/admin")
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
	_, err := users.InsertOne(c, User{_id, username, password, timestamp{time.Now(), time.Now(), 0}, "user"})
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

func updateTimeStamp(c *gin.Context, users *mongo.Collection) {
	filter := bson.D{{Key: "_id", Value: GetUser().Id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{{Key: "timestamp.last", Value: time.Now()}}},
		{Key: "$set", Value: bson.D{{Key: "timestamp.visits_n", Value: GetUser().TimeStamp.VisitsN + 1}}},
	}
	users.UpdateOne(c, filter, update)
}
