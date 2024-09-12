package repositories

import (
	"context"
	"fmt"

	"github.com/Whisp-0/ToDoRPG/cmd/database"
	"github.com/Whisp-0/ToDoRPG/cmd/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateXpBar(b models.Xpbar) (models.Xpbar, error) {
	db := database.GetDB()
	collection := db.Database("ToDoRPG").Collection("XpBars")
	b.CurrentXp = 0
	b.RequaredXp = 1000

	xpBar := bson.D{
		{Key: "name", Value: b.Name},
		{Key: "user_name", Value: b.UserName},
		{Key: "current_xp", Value: b.CurrentXp},
		{Key: "requared_xp", Value: b.RequaredXp},
	}

	// Добавляем пользовтаеля в БД
	result, err := collection.InsertOne(context.TODO(), xpBar)
	if err != nil {
		return b, err
	}

	b.Id = result.InsertedID.(primitive.ObjectID)

	return b, nil
}

func FindAllXpBar(filter bson.D) ([]models.Xpbar, error) {
	db := database.GetDB()
	collection := db.Database("ToDoRPG").Collection("XpBars")

	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	var results []models.Xpbar

	for cursor.Next(context.TODO()) {
		var xpBar models.Xpbar
		if err := cursor.Decode(&xpBar); err != nil {
			return nil, err
		}
		results = append(results, xpBar)
	}

	fmt.Println("Список задач:", results)

	return results, nil
}
