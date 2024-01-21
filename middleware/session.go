// session.go
package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Session(c *fiber.Ctx) error {
	store := session.New()
	session, err := store.Get(c)
	if err != nil {
		log.Fatal("error getting session")
	}
	c.Locals("session", session)
	return c.Next()
}
