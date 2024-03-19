package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/hossokawa/go-todo-app/internal/db"
	"github.com/hossokawa/go-todo-app/model"
	"github.com/hossokawa/go-todo-app/view"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func GetRegisterScreen(c echo.Context) error {
	component := view.Register()
	return component.Render(context.Background(), c.Response().Writer)
}

func RegisterUser(c echo.Context) error {
	coll := db.GetCollection("users")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	email := c.FormValue("email")
	password := c.FormValue("password")
	hashedPwd := getHash([]byte(password))

	user := model.User{
		Email:    email,
		Password: hashedPwd,
	}

	_, err = coll.InsertOne(context.TODO(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "The provided email is already registered")
	}

	component := view.Register()
	return component.Render(context.Background(), c.Response().Writer)
}

func GetLoginScreen(c echo.Context) error {
	component := view.Login()
	return component.Render(context.Background(), c.Response().Writer)
}

func LoginUser(c echo.Context) error {
	coll := db.GetCollection("users")

	user := model.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	var dbUser model.User

	err := coll.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&dbUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	userPwd := []byte(user.Password)
	hashedPwd := []byte(dbUser.Password)
	err = bcrypt.CompareHashAndPassword(hashedPwd, userPwd)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid password")
	}

	tasks, err := refreshTasks()

	component := view.Index(tasks, true)
	return component.Render(context.Background(), c.Response().Writer)
}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
