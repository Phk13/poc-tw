package db

import (
	"context"
	"time"

	"github.com/phk13/poc-tw/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* CheckIfUserExists receives an email and checks if it exists in DB.*/
func CheckIfUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := DBConnector.Database("twittor")
	col := db.Collection("usuarios")

	condition := bson.M{"email": email}

	var result models.User
	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
