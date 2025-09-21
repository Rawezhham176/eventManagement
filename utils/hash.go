package utils

import "golang.org/x/crypto/bcrypt"

func ToHashPassword(password string) (string, error) {
	r, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(r), err
}

func CompareCredential(receivedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(receivedPassword), []byte(password))
	return err == nil
}
