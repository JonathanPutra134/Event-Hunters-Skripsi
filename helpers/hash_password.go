package helpers

import (
	"github.com/volatiletech/null/v8"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) null.String {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return null.StringFrom("")
	}

	return null.StringFrom(string(hashedPassword))
}
func ComparePasswords(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
