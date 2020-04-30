package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnect ....
var MongoConnect = ConnectDb()
var clientOptions = options.Client().ApplyURI("mongodb+srv://kevin:y7FkYbtCiPzhExJu@eevee-mpfyc.mongodb.net/test?retryWrites=true&w=majority")

// ConnectDb ...
func ConnectDb() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Database online")
	return client
}

// CheckConnection ...
func CheckConnection() bool {
	err := MongoConnect.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
