package routers

import (
	"encoding/json"
	"net/http"

	"github.com/phk13/poc-tw/db"
	"github.com/phk13/poc-tw/models"
)

/* CheckRelationship checks if there is a relationship between two users.*/
func CheckRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send parameter id", http.StatusBadRequest)
		return
	}

	var rel models.Relationship
	rel.UserID = UserID
	rel.UserRelationshipID = ID

	var response models.ResponseRelationship

	status, err := db.FindRelationship(rel)
	response.Status = err == nil || status

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}
