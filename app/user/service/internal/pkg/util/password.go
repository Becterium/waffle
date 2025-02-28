package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(hash, password string) bool {
	fmt.Println()
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
