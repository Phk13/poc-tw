package db

import "golang.org/x/crypto/bcrypt"

/* EncryptPassword encrypts a given password*/
func EncryptPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 8)
	return string(bytes), err
}
