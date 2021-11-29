package db

import (
	"log"

	"github.com/phk13/poc-tw/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadAllUsers(ID string, page int64, search string, type_reg string) ([]*models.User, bool) {
	/* ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := DBConnector.Database("twittor")
	col := db.Collection("users") */
	col, ctx, cancel := GetCollection("users")
	defer cancel()

	var results []*models.User

	opts := options.Find()
	opts.SetSkip((page - 1) * 20)
	opts.SetLimit(20)

	query := bson.M{
		"firstName": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, opts)
	if err != nil {
		log.Println(err.Error())
		return results, false
	}

	var found, include bool
	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Println(err.Error())
			return results, false
		}

		var rel models.Relationship
		rel.UserID = ID
		rel.UserRelationshipID = user.ID.Hex()

		include = false

		found, _ = FindRelationship(rel)
		if type_reg == "new" && !found {
			include = true
		}
		if type_reg == "follow" && found {
			include = true
		}
		if rel.UserRelationshipID == ID {
			include = false
		}

		if include {
			user.Password = ""
			user.Bio = ""
			user.Site = ""
			user.Location = ""
			user.Banner = ""
			user.Email = ""

			results = append(results, &user)
		}
	}

	err = cursor.Err()
	if err != nil {
		log.Println(err.Error())
		return results, false
	}
	cursor.Close(ctx)
	return results, true
}
