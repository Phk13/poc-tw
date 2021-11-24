package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/phk13/poc-tw/db"
	"github.com/phk13/poc-tw/jwt"
	"github.com/phk13/poc-tw/models"
)

/* Login tries to login using credentials and returns a JWT token if successful*/
func Login(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Email or password invalid "+err.Error(), http.StatusBadRequest)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	userDB, exists := db.TryLogin(user.Email, user.Password)
	if !exists {
		http.Error(w, "Email or password invalid", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerateJWT(userDB)
	if err != nil {
		http.Error(w, "An error happened while generating login token", http.StatusInternalServerError)
		return
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
