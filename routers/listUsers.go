package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/phk13/poc-tw/db"
)

/* ListUsers reads a list of all users matching serach expression and relationship type*/
func ListUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page_str := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page_str)
	if err != nil {
		http.Error(w, "Must send parameter page > 0", http.StatusBadRequest)
		return
	}

	page := int64(pageTemp)
	result, status := db.ReadAllUsers(UserID, page, search, typeUser)
	if !status {
		http.Error(w, "Error while reading users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)
}
