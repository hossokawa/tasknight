package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hossokawa/go-todo-app/internal/db"
	"github.com/hossokawa/go-todo-app/internal/handlers"
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
	app.Static("/static", "static")

	// Routes
	app.GET("/", handlers.Home)
	app.GET("/register", handlers.GetRegisterScreen)
	app.GET("/login", handlers.GetLoginScreen)
	app.GET("/tasks/:id", handlers.GetTask)
	app.POST("/", handlers.CreateTask)
	app.POST("/register", handlers.RegisterUser)
	app.POST("/login", handlers.LoginUser)
	app.DELETE("/tasks/:id", handlers.DeleteTask)
	app.PATCH("/tasks/:id", handlers.UpdateTask)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Logger.Fatal(app.Start(":" + port))
}
