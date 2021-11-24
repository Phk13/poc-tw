package models

/* ResponseRelationship has status true or false if relationship exists between two users.*/
type ResponseRelationship struct {
	Status bool `json:"status"`
}
