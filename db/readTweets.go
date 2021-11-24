package db

import (
	"context"
	"log"
	"time"

	"github.com/phk13/poc-tw/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* ReadTweets reads tweets from a profile*/
func ReadTweets(ID string, page int64) ([]*models.ReturnTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := DBConnector.Database("twittor")
	col := db.Collection("tweet")

	var results []*models.ReturnTweets

	condition := bson.M{
		"userid": ID,
	}

	opt := options.Find()
	opt.SetLimit(20)
	opt.SetSort(bson.D{{Key: "date", Value: -1}})
	opt.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, opt)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}
	for cursor.Next(context.TODO()) {
		var register models.ReturnTweets
		err := cursor.Decode(&register)
		if err != nil {
			return results, false
		}
		results = append(results, &register)
	}
	return results, true
}
