package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/phk13/poc-tw/db"
	"github.com/phk13/poc-tw/models"
)

var Email string
var UserID string

/* ProcessToken processes a token to extract its values*/
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	key := []byte("secret")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("invalid token format")
	}
	token = strings.TrimSpace(splitToken[1])
	tk, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil {
		_, found, _ := db.CheckIfUserExists(claims.Email)
		if found {
			Email = claims.Email
			UserID = claims.ID.Hex()
			return claims, found, UserID, nil
		}
		if !tk.Valid {
			return claims, false, "", errors.New("invalid token")
		}

	}
	return claims, false, "", err

}
