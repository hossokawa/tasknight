package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hossokawa/go-todo-app/internal/auth"
	"github.com/hossokawa/go-todo-app/internal/db"
	"github.com/hossokawa/go-todo-app/internal/models"
	"github.com/hossokawa/go-todo-app/view"
	"github.com/hossokawa/go-todo-app/view/components"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Home(c echo.Context) error {
	tasks, err := fetchTasks()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	userLoggedIn, err := checkLoginStatus(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if userLoggedIn {
		component := view.Index(tasks, true)
		return component.Render(context.Background(), c.Response().Writer)
	}

	component := view.Index(tasks, false)
	return component.Render(context.Background(), c.Response().Writer)
}

func EditTask(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Id cannot be empty")
	}

	task, err := fetchTask(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve task")
	}

	component := view.TaskEdit(&task)
	return component.Render(context.Background(), c.Response().Writer)
}

type taskDTO struct {
	Name      string `bson:"name" form:"name"`
	Completed bool   `bson:"completed"`
}

func CreateTask(c echo.Context) error {
	if c.FormValue("name") == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Task name cannot be empty")
	}

	t := taskDTO{
		Name:      c.FormValue("name"),
		Completed: false,
	}

	err := createTask(t)
	if err != nil {
		return err
	}

	tasks, err := fetchTasks()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error refreshing tasks")
	}

	component := components.TasksWithInput(tasks)
	return component.Render(context.Background(), c.Response().Writer)
}

func UpdateTask(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Id cannot be empty")
	}

	updatedName := c.FormValue("name")
	if updatedName == "" {
		return errors.New("Task name cannot be empty")
	}
	var updatedStatus bool
	if c.FormValue("completed") == "on" {
		updatedStatus = true
	} else {
		updatedStatus = false
	}

	task, err := updateTask(id, updatedName, updatedStatus)
	if err != nil {
		return err
	}

	component := view.TaskEdit(&task)
	return component.Render(context.Background(), c.Response().Writer)
}

func DeleteTask(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Id cannot be empty")
	}

	err := deleteTask(id)
	if err != nil {
		return err
	}

	tasks, err := fetchTasks()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching tasks")
	}

	component := components.TaskList(tasks)
	return component.Render(context.Background(), c.Response().Writer)
}

func fetchTasks() ([]model.Task, error) {
	coll := db.GetCollection("tasks")

	tasks := make([]model.Task, 0)
	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for cursor.Next(context.TODO()) {
		task := model.Task{}
		err := cursor.Decode(&task)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func fetchTask(id string) (model.Task, error) {
	coll := db.GetCollection("tasks")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Task{}, echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}
	task := model.Task{}

	err = coll.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&task)
	if err != nil {
		return model.Task{}, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return task, nil
}

func createTask(t taskDTO) error {
	coll := db.GetCollection("tasks")
	_, err := coll.InsertOne(context.TODO(), t)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil
}

func updateTask(id string, name string, completed bool) (model.Task, error) {
	coll := db.GetCollection("tasks")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Task{}, echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}

	_ = coll.FindOneAndUpdate(context.TODO(), bson.D{{Key: "_id", Value: objectId}}, bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: name}}}, {Key: "$set", Value: bson.D{{Key: "completed", Value: completed}}}})

	task := model.Task{}
	err = coll.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&task)
	if err != nil {
		return model.Task{}, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return task, nil
}

func deleteTask(id string) error {
	coll := db.GetCollection("tasks")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}

	_ = coll.FindOneAndDelete(context.TODO(), bson.M{"_id": objectId})

	return nil
}

func checkLoginStatus(c echo.Context) (bool, error) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return false, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if cookie == nil {
		return false, nil
	}
	token, err := auth.ValidateJWT(cookie.Value)
	if err != nil {
		return false, echo.NewHTTPError(http.StatusUnauthorized, "failed to authenticate token")
	}
	if !token.Valid {
		return false, echo.NewHTTPError(http.StatusUnauthorized, "failed to authenticate token")
	}
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)

	_, err = GetUser(userId)
	if err != nil {
		return false, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return true, nil
}
