package routers

import (
	"encoding/json"
	"net/http"

	"github.com/phk13/poc-tw/db"
	"github.com/phk13/poc-tw/models"
)

/* ModifyProfile modifies a user profile.*/
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Incorrect data "+err.Error(), http.StatusBadRequest)
		return
	}
	status, err := db.ModifyRegister(user, UserID)
	if err != nil {
		http.Error(w, "An error happened while modifying the register. Try again "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !status {
		http.Error(w, "Could not modify the user register "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
