package handlers

import (
	"context"
	"net/http"

	"github.com/hossokawa/go-todo-app/internal/db"
	"github.com/hossokawa/go-todo-app/model"
	"github.com/hossokawa/go-todo-app/view"
	"github.com/hossokawa/go-todo-app/view/components"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTasks(c echo.Context) error {
	coll := db.GetCollection("tasks")

	// find all tasks
	tasks := make([]model.Task, 0)
	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for cursor.Next(context.TODO()) {
		task := model.Task{}
		err := cursor.Decode(&task)
		if err != nil {

			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		tasks = append(tasks, task)
	}

	component := view.Index(tasks)
	return component.Render(context.Background(), c.Response().Writer)
}

func GetTask(c echo.Context) error {
	coll := db.GetCollection("tasks")

	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Id cannot be empty")
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}

	task := model.Task{}

	err = coll.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&task)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
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

	coll := db.GetCollection("tasks")
	_, err := coll.InsertOne(context.TODO(), t)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	tasks, err := refreshTasks()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error refreshing tasks")
	}

	component := components.TasksWithInput(tasks)
	return component.Render(context.Background(), c.Response().Writer)
}

func refreshTasks() ([]model.Task, error) {
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
