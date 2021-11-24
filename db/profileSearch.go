package db

import (
	"context"
	"log"
	"time"

	"github.com/phk13/poc-tw/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* ProfileSearch searches for a profile in DB.*/
func ProfileSearch(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := DBConnector.Database("twittor")
	col := db.Collection("users")

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		log.Println("Register not found " + err.Error())
		return profile, err
	}
	return profile, nil
}
