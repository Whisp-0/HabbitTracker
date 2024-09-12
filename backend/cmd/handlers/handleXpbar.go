package handlers

import (
	"log"
	"net/http"

	"github.com/Whisp-0/ToDoRPG/cmd/models"
	"github.com/Whisp-0/ToDoRPG/cmd/repositories"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateXpBar(c echo.Context) error {
	xpBar := models.Xpbar{}
	c.Bind(&xpBar)
	newXpBar, err := repositories.CreateXpBar(xpBar)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newXpBar)
}

func GetAllXpBarByUsername(c echo.Context) error {

	username := c.Param("name")

	filter := bson.D{{Key: "user_name", Value: username}}
	log.Println("запрос на получение всех навыков для пользователя:", username)
	xpBars, err := repositories.FindAllXpBar(filter)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	//jsonData, err := json.Marshal(xpBars)
	// if err != nil {
	// 	log.Fatal("Не удалось распарсить JSON")
	// 	return c.JSON(http.StatusBadRequest, err)
	// }

	return c.JSON(http.StatusOK, xpBars)
}
