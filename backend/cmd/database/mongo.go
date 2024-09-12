package database

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func InitMongo() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	db, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecnting to mongo db")
	}

	err = db.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Error connecnting to mongo db")
	}
}

func GetDB() *mongo.Client {
	return db
}
