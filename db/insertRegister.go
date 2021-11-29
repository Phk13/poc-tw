package db

import (
	"github.com/phk13/poc-tw/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertRegister inserts user data into DB.*/
func InsertRegister(u models.User) (string, bool, error) {
	/* ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := DBConnector.Database("twittor")
	col := db.Collection("users") */
	col, ctx, cancel := GetCollection("users")
	defer cancel()

	u.Password, _ = EncryptPassword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
