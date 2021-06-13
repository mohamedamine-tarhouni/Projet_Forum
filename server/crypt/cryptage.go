package cryptage

import (
	"golang.org/x/crypto/bcrypt"
)

//this function takes a string and encrypt it
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//this function takes a hashed password normal password and check if they are the same
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//this function will be used to check on the password entered by the user
func Verif(password, db_password string) bool {
	match := CheckPasswordHash(password, db_password)
	return match
}
