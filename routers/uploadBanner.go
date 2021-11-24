package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/phk13/poc-tw/db"
	"github.com/phk13/poc-tw/models"
)

func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, "Banner form invalid "+err.Error(), http.StatusBadRequest)
		return
	}
	ext := strings.Split(handler.Filename, ".")[1]
	filePath := "uploads/banners/" + UserID + "." + ext

	fileWriter, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error while uploading image "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		http.Error(w, "Error while copying image "+err.Error(), http.StatusInternalServerError)
		return
	}

	var user models.User

	user.Banner = UserID + "." + ext
	status, err := db.ModifyRegister(user, UserID)
	if err != nil || !status {
		http.Error(w, "Error while saving image in DB "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
