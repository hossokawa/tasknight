package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/hossokawa/go-todo-app/model"
	"github.com/hossokawa/go-todo-app/view"
	"github.com/hossokawa/go-todo-app/view/components"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func filterById(tasks []*model.Task, id string) (out []*model.Task) {
	for _, task := range tasks {
		if task.Id == id {
			continue
		}
		out = append(out, task)
	}
	return out
}

func findById(tasks []*model.Task, id string) (task *model.Task) {
	for _, task := range tasks {
		if task.Id == id {
			return task
		}
	}
	return nil
}

func main() {
	tasks := []*model.Task{
		{
			Id:        uuid.NewString(),
			Name:      "Code todo app",
			Completed: true,
		},
		{
			Id:        uuid.NewString(),
			Name:      "Go gym",
			Completed: false,
		},
		{
			Id:        uuid.NewString(),
			Name:      "Walk the dog",
			Completed: true,
		},
	}

	app := echo.New()

	//Middleware
	app.Use(middleware.Static("/static"))

	// Routes
	app.GET("/", func(c echo.Context) error {
		component := view.Index(tasks)
		return component.Render(context.Background(), c.Response().Writer)
	})

	app.GET("/register", func(c echo.Context) error {
		component := view.Register()
		return component.Render(context.Background(), c.Response().Writer)
	})

	app.GET("/login", func(c echo.Context) error {
		component := view.Login()
		return component.Render(context.Background(), c.Response().Writer)
	})

	app.GET("/tasks/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
		}
		task := findById(tasks, id)
		if task == nil {
			return echo.NewHTTPError(http.StatusNotFound, "Could not find task with the provided id")
		}
		component := view.TaskEdit(task)
		return component.Render(context.Background(), c.Response().Writer)
	})

	app.POST("/tasks", func(c echo.Context) error {
		name := c.FormValue("name")
		if name == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Task name cannot be empty")
		}
		tasks = append(tasks, &model.Task{
			Id:        uuid.NewString(),
			Name:      name,
			Completed: false,
		})
		component := components.TasksWithInput(tasks)
		return component.Render(context.Background(), c.Response().Writer)
	})

	app.DELETE("/tasks/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
		}
		tasks = filterById(tasks, id)
		component := components.TaskList(tasks)
		return component.Render(context.Background(), c.Response().Writer)
	})

	app.PATCH("/tasks/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
		}
		task := findById(tasks, id)
		task.Name = c.FormValue("name")
		if c.FormValue("completed") == "on" {
			task.Completed = true
		} else {
			task.Completed = false
		}
		fmt.Println(task)
		component := view.TaskEdit(task)
		return component.Render(context.Background(), c.Response().Writer)
	})

	app.Static("/static", "static")

	app.Logger.Fatal(app.Start(":8080"))
}
