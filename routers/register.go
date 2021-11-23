package routers

import (
	"encoding/json"
	"net/http"

	"github.com/phk13/poc-tw/db"
	"github.com/phk13/poc-tw/models"
)

/* Register creates in DB an user register.*/
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		http.Error(w, "Error in received data " + err.Error(), http.StatusBadRequest)
	}
	if len(user.Email) == 0 {
		http.Error(w, "User email is required", http.StatusBadRequest)
		return
	}
	if len(user.Password) < 6 {
		http.Error(w, "Password should contain at least 6 characters", http.StatusBadRequest) 
		return
	}

	_, found, _ := db.CheckIfUserExists(user.Email)
	if found {
		http.Error(w, "User with provided email already exists", http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertRegister(user)
	if err != nil {
		http.Error(w, "An error happened while registering user " + err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(w, "User register was unsuccessful" , http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
