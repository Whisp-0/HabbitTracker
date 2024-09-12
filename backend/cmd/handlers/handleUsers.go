package handlers

import (
	"log"
	"net/http"

	"github.com/Whisp-0/ToDoRPG/cmd/models"
	"github.com/Whisp-0/ToDoRPG/cmd/repositories"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	newUser, err := repositories.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newUser)
}

func FindUserByName(c echo.Context) error {
	name := c.Param("name")

	filter := bson.D{{Key: "name", Value: name}}

	log.Println("Поиск пользователя по данным: ", name)

	user, err := repositories.FindUser(filter)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, user)
}

func DeleteUserByName(c echo.Context) error {
	name := c.Param("name")

	filter := bson.D{{Key: "name", Value: name}}

	log.Println("Удаление пользователя по данным: ", name)

	deleteResult, err := repositories.DeleteUser(filter)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, deleteResult)
}
