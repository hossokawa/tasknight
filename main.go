package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/hossokawa/go-todo-app/internal/db"
	"github.com/hossokawa/go-todo-app/internal/handlers"
	"github.com/hossokawa/go-todo-app/model"
	"github.com/hossokawa/go-todo-app/view"
	"github.com/hossokawa/go-todo-app/view/components"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func filterById(tasks []model.Task, id string) (out []model.Task) {
	for _, task := range tasks {
		if task.Id == id {
			continue
		}
		out = append(out, task)
	}
	return out
}

func findById(tasks []model.Task, id string) (task *model.Task) {
	for _, task := range tasks {
		if task.Id == id {
			return &task
		}
	}
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env")
	}

	err = db.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tasks := []model.Task{
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
	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("%v | %v | %v\n", v.Status, v.Method, v.URI)
			return nil
		},
	}))
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(middleware.Static("/static"))

	// Routes
	app.GET("/", handlers.GetTasks)

	app.GET("/register", func(c echo.Context) error {
		component := view.Register()
		return component.Render(context.Background(), c.Response().Writer)
	})

	app.GET("/login", func(c echo.Context) error {
		component := view.Login()
		return component.Render(context.Background(), c.Response().Writer)
	})

	// app.GET("/tasks/:id", func(c echo.Context) error {
	// 	id := c.Param("id")
	// 	if id == "" {
	// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	// 	}
	// 	task := findById(tasks, id)
	// 	if task == nil {
	// 		return echo.NewHTTPError(http.StatusNotFound, "Could not find task with the provided id")
	// 	}
	// 	component := view.TaskEdit(task)
	// 	return component.Render(context.Background(), c.Response().Writer)
	// })

	app.GET("/tasks/:id", handlers.GetTask)

	// app.POST("/tasks", func(c echo.Context) error {
	// 	name := c.FormValue("name")
	// 	if name == "" {
	// 		return echo.NewHTTPError(http.StatusBadRequest, "Task name cannot be empty")
	// 	}
	// 	tasks = append(tasks, model.Task{
	// 		Id:        uuid.NewString(),
	// 		Name:      name,
	// 		Completed: false,
	// 	})
	// 	component := components.TasksWithInput(tasks)
	// 	return component.Render(context.Background(), c.Response().Writer)
	// })

	app.POST("/tasks", handlers.CreateTask)

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
		component := view.TaskEdit(task)
		return component.Render(context.Background(), c.Response().Writer)
	})

	app.Static("/static", "static")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Logger.Fatal(app.Start(":" + port))
}
