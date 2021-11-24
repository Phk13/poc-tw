package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/phk13/poc-tw/db"
)

/* GetFollowerTweets reads tweets from all our followers.*/
func GetFollowerTweets(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Must send parameter page", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Page value is incorrect", http.StatusBadRequest)
		return
	}

	response, status := db.ReadFollowerTweets(UserID, page)
	if !status {
		http.Error(w, "Error while reading tweets", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}
