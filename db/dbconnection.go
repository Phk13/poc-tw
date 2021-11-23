package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* DBConnector exports a Mongo connection. */
var DBConnector = connectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://poc-tw:ujHfMC9xjqM8paxl@poc-tw.d7cbd.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

func connectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successful DB connection")
	return client
}

/* CheckConnection returns if DB is still connected. */
func CheckConnection() bool {
	err := DBConnector.Ping(context.TODO(), nil)
	return err == nil
}
