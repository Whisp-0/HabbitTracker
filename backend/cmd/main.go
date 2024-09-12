package main

import (
	"context"
	"net/http"

	"github.com/Whisp-0/ToDoRPG/cmd/database"
	"github.com/Whisp-0/ToDoRPG/cmd/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func stopServer(c echo.Context) error {
	err := c.Echo().Shutdown(context.Background())
	if err != nil {
		if err != http.ErrServerClosed {
			c.Echo().Logger.Fatal("shutting down the server")
		}
	}
	return nil
}

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Content-Type"},
	}))

	//Загрузить файл с переменными окружения
	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal(err)
	}

	//Создать подключение к базе данных
	database.InitMongo()

	e.GET("/", handlers.Home)

	e.GET("/stop", stopServer)

	// users router
	e.GET("/users/account/:name", handlers.FindUserByName)

	e.DELETE("/users/delete/:name", handlers.DeleteUserByName)

	e.POST("/users", handlers.CreateUser)

	//xpbar router

	e.POST("/xpbar/new", handlers.CreateXpBar)

	e.GET("users/account/:name/tasks", handlers.GetAllXpBarByUsername)

	e.Logger.Fatal(e.Start(":3000"))
}
