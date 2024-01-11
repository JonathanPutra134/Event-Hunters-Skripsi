package helpers

import (
	"log"

	"github.com/volatiletech/null/v8"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) null.String {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("ERRORR HASHING")
		return null.StringFrom("")
	}

	return null.StringFrom(string(hashedPassword))
}
