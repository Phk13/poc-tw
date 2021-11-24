package models

/* Relationship saves a relationship between two users*/
type Relationship struct {
	UserID             string `bson:"userid" json:"userId"`
	UserRelationshipID string `bson:"userrelationshipid" json:"userRelationshipId"`
}
