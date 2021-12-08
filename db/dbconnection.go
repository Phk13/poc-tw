package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/phk13/poc-tw/config"
)

/* DBConnector exports a Mongo connection. */
var DBConnector = connectDB()
var clientOptions = options.Client().ApplyURI(config.AppCfg.Database)

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


func GetCollection(col string) (*mongo.Collection, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	collection := DBConnector.Database("twittor").Collection(col)
	return collection, ctx, cancel
}