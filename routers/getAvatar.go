package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/phk13/poc-tw/db"
)

/* GetAvatar sends the avatar image.*/
func GetAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send parameter id", http.StatusBadRequest)
		return
	}

	profile, err := db.ProfileSearch(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	file, err := os.Open("uploads/avatar/" + profile.Avatar)
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error copying image", http.StatusInternalServerError)
		return
	}
}
