package routers

import (
	"encoding/json"
	"net/http"

	"github.com/phk13/poc-tw/db"
)

/* ViewProfile retrieves a profile values*/
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must use parameter ID", http.StatusBadRequest)
		return
	}

	profile, err := db.ProfileSearch(ID)
	if err != nil {
		http.Error(w, "An error happened while searching the profile "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
