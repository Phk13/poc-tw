package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/phk13/poc-tw/db"
)

/* GetTweets recovers tweets*/
func GetTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send parameter id", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Must send parameter page", http.StatusBadRequest)
		return
	}
	page_int, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Page value is incorrect", http.StatusBadRequest)
		return
	}
	page := int64(page_int)
	results, status := db.ReadTweets(ID, page)
	if !status {
		http.Error(w, "Error while reading tweets", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}
