package db

import (
	"github.com/phk13/poc-tw/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* RedFollowerTweets reads tweets from all followers.*/
func ReadFollowerTweets(ID string, page int) ([]models.ReturnFollowerTweets, bool) {
	/* ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := DBConnector.Database("twittor")
	col := db.Collection("relationship") */
	col, ctx, cancel := GetCollection("relationship")
	defer cancel()

	skip := (page - 1) * 20
	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userrelationshipid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, _ := col.Aggregate(ctx, conditions)
	var result []models.ReturnFollowerTweets
	err := cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
