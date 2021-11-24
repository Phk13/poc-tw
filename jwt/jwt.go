package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/phk13/poc-tw/models"
)

/* GenerateJWT encrypts a user into a JWT*/
func GenerateJWT(user models.User) (string, error) {
	key := []byte("secret")

	payload := jwt.MapClaims{
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"birthdate": user.BirthDate,
		"bio":       user.Bio,
		"location":  user.Location,
		"site":      user.Site,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
