package routers

import (
	"net/http"

	"github.com/phk13/poc-tw/db"
)

/* RemoveTweet allows to remove a given tweet.*/
func RemoveTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send parameter id", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, UserID)
	if err != nil {
		http.Error(w, "An error happened while deleting the tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
