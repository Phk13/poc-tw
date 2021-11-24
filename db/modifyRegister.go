package db

import (
	"context"
	"time"

	"github.com/phk13/poc-tw/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* ModifyRegister modifies an already existing profile.*/
func ModifyRegister(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := DBConnector.Database("twittor")
	col := db.Collection("users")

	register := make(map[string]interface{})
	if len(u.FirstName) > 0 {
		register["firstName"] = u.FirstName
	}
	if len(u.LastName) > 0 {
		register["lastName"] = u.LastName
	}
	register["birthDate"] = u.BirthDate
	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}
	if len(u.Bio) > 0 {
		register["bio"] = u.Bio
	}
	if len(u.Location) > 0 {
		register["location"] = u.Location
	}
	if len(u.Site) > 0 {
		register["site"] = u.Site
	}

	updateString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateString)
	return err == nil, err
}
