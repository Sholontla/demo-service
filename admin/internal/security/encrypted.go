package security

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func SetPassword(password string) (string, error) {
	cost := 8
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Panic("Error ")
	}
	return string(hashedPassword), nil
}

func ComparePassword(s string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(s), []byte(password))
}
