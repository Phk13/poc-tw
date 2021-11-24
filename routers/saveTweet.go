package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/phk13/poc-tw/db"
	"github.com/phk13/poc-tw/models"
)

/* SaveTweet allows to save a tweet in DB.*/
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Incorrect request "+err.Error(), http.StatusBadRequest)
		return
	}

	register := models.SaveTweet{
		UserID:  UserID,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.InsertTweet(register)
	if err != nil {
		http.Error(w, "An error happened while inserting the tweet. Try again "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !status {
		http.Error(w, "Could not insert the tweet", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
