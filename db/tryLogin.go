package db

import (
	"github.com/phk13/poc-tw/models"
	"golang.org/x/crypto/bcrypt"
)

/* TryLogin validates a login against DB */
func TryLogin(email string, password string) (models.User, bool) {
	user, found, _ := CheckIfUserExists(email)
	if !found {
		return user, false
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, false
	}
	return user, true
}
