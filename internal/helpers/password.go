package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)

}

func CheckPassword(hashedPassword string, password string) (isPasswordValid bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}
