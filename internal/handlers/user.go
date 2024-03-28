package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/hossokawa/go-todo-app/internal/auth"
	"github.com/hossokawa/go-todo-app/internal/db"
	"github.com/hossokawa/go-todo-app/internal/models"
	"github.com/hossokawa/go-todo-app/view"
	"github.com/hossokawa/go-todo-app/view/components"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func GetRegisterScreen(c echo.Context) error {
	component := view.Register()
	return component.Render(context.Background(), c.Response().Writer)
}

type userDTO struct {
	Username string `bson:"username"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

func RegisterUser(c echo.Context) error {
	coll := db.GetCollection("users")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}, {Key: "_id", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	hashedPwd := getHash([]byte(password))

	user := userDTO{
		Username: username,
		Email:    email,
		Password: hashedPwd,
	}

	_, err = coll.InsertOne(context.TODO(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "The provided email is already registered")
	}

	c.Response().Header().Set("HX-Redirect", "/")
	c.Response().WriteHeader(http.StatusOK)

	return nil
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
		c.Response().WriteHeader(http.StatusInternalServerError)
		component := components.ErrorMsg("Something went wrong. Try again later.")
		return component.Render(context.Background(), c.Response().Writer)
	}

	userPwd := []byte(user.Password)
	hashedPwd := []byte(dbUser.Password)
	err = bcrypt.CompareHashAndPassword(hashedPwd, userPwd)
	if err != nil {
		c.Response().WriteHeader(http.StatusUnauthorized)
		component := components.ErrorMsg("Invalid email or password")
		return component.Render(context.Background(), c.Response().Writer)
	}

	token, err := auth.GenerateJWT(dbUser.Id, dbUser.Username, dbUser.Email)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		component := components.ErrorMsg("Error generating JWT. Try again later.")
		return component.Render(context.Background(), c.Response().Writer)
	}

	expiration := time.Now().Add(30 * 24 * time.Hour)
	cookie := http.Cookie{Name: "jwt", Value: token, Expires: expiration, Path: "/", HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode}

	c.SetCookie(&cookie)

	c.Response().Header().Set("HX-Redirect", "/")
	c.Response().WriteHeader(http.StatusOK)

	return nil
}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func GetUser(id string) (model.User, error) {
	coll := db.GetCollection("users")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.User{}, echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}
	user := model.User{}

	err = coll.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return model.User{}, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return user, nil

}
