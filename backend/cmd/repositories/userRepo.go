package repositories

import (
	"context"
	"log"
	"time"

	"github.com/Whisp-0/ToDoRPG/cmd/database"
	"github.com/Whisp-0/ToDoRPG/cmd/models"
	"github.com/Whisp-0/ToDoRPG/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Создание пользователя
func CreateUser(u models.User) (models.User, error) {
	db := database.GetDB()

	// Получаем ссылку на коллекцию
	collection := db.Database("ToDoRPG").Collection("myUsers")

	u.CreatedAt = time.Now()
	u.UpdatedAt = u.CreatedAt
	u.Password, _ = utils.HashPassword(u.Password)

	user := bson.D{
		{Key: "email", Value: u.Email},
		{Key: "password", Value: u.Password},
		{Key: "name", Value: u.Name},
		{Key: "created_at", Value: u.CreatedAt},
		{Key: "updated_at", Value: u.UpdatedAt},
	}

	// Добавляем пользовтаеля в БД
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return u, err
	}

	u.Id = result.InsertedID.(primitive.ObjectID)

	return u, nil
}

// Поиск пользователя
func FindUser(filter bson.D) (bson.M, error) {
	db := database.GetDB()
	collection := db.Database("ToDoRPG").Collection("myUsers")

	var result bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	log.Println("Найден документ", result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func DeleteUser(filter bson.D) (*mongo.DeleteResult, error) {
	db := database.GetDB()
	collection := db.Database("ToDoRPG").Collection("myUsers")

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return result, nil
}
