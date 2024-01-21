package helpers

import (
	"event-hunters/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SetUserSession(c *fiber.Ctx, user *models.User) error {
	session := c.Locals("session").(*session.Session)
	session.Set("userID", user.ID)
	if err := session.Save(); err != nil {
		// Handle the error, e.g., log it or return an internal server error
		return err
	}
	return nil
}
