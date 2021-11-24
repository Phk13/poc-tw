package routers

import (
	"net/http"
	"github.com/phk13/poc-tw/db"
	"github.com/phk13/poc-tw/models"
)

/* RemoveRelationship deletes a relationship between users.*/
func RemoveRelationship(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send parameter id", http.StatusBadRequest)
		return
	}

	var rel models.Relationship
	rel.UserID = UserID
	rel.UserRelationshipID = ID

	status, err := db.DeleteRelationship(rel)
	if err != nil {
		http.Error(w, "An error happened while deleting relationship "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !status {
		http.Error(w, "Relationship could not be deleted", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}