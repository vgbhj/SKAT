package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// EncodeBcrypt
func EncodeBcrypt(value string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error while hashing password: %v", err)
	}
	return string(hashedPassword)
}
