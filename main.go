package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hossokawa/go-todo-app/internal/db"
	"github.com/hossokawa/go-todo-app/internal/handlers"
	"github.com/hossokawa/go-todo-app/view"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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

	// tasks := []model.Task{
	// 	{
	// 		Id:        uuid.NewString(),
	// 		Name:      "Code todo app",
	// 		Completed: true,
	// 	},
	// 	{
	// 		Id:        uuid.NewString(),
	// 		Name:      "Go gym",
	// 		Completed: false,
	// 	},
	// 	{
	// 		Id:        uuid.NewString(),
	// 		Name:      "Walk the dog",
	// 		Completed: true,
	// 	},
	// }

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

	// Routes
	app.GET("/register", func(c echo.Context) error {
		component := view.Register()
		return component.Render(context.Background(), c.Response().Writer)
	})
	app.GET("/login", func(c echo.Context) error {
		component := view.Login()
		return component.Render(context.Background(), c.Response().Writer)
	})
	app.GET("/", handlers.GetTasks)
	app.GET("/tasks/:id", handlers.GetTask)
	app.POST("/tasks", handlers.CreateTask)
	app.DELETE("/tasks/:id", handlers.DeleteTask)
	app.PATCH("/tasks/:id", handlers.UpdateTask)

	app.Static("/static", "static")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Logger.Fatal(app.Start(":" + port))
}
