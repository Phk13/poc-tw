package db

import (
	"context"
	"time"

	"github.com/phk13/poc-tw/models"
)

/* DeleteRelationship deletes a relationship in DB.*/
func DeleteRelationship(rel models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := DBConnector.Database("twittor")
	col := db.Collection("relationship")

	_, err := col.DeleteOne(ctx, rel)
	if err != nil {
		return false, err
	}
	return true, nil
}
