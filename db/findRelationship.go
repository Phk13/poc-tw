package db

import (
	"log"

	"github.com/phk13/poc-tw/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* FindRelationship searches a relationship between two users.*/
func FindRelationship(rel models.Relationship) (bool, error) {
	/* ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := DBConnector.Database("twittor")
	col := db.Collection("relationship") */
	col, ctx, cancel := GetCollection("relationship")
	defer cancel()

	condition := bson.M{
		"userid":             rel.UserID,
		"userrelationshipid": rel.UserRelationshipID,
	}

	var result models.Relationship
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	return true, nil
}
