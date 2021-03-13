package utils

import "golang.org/x/crypto/bcrypt"

//EncryptPassword function
func EncryptPassword(pass string) (string, error) {
	//Cost for encrypt
	cost := 8
	//Encrypt the password
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)

	return string(bytes), err
}
