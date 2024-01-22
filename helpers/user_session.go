package helpers

import (
	"event-hunters/models"
	"fmt"
	"github.com/friendsofgo/errors"

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

func GetUserIDFromSession(c *fiber.Ctx) (int, error) {
	session := c.Locals("session").(*session.Session)
	userIDValue := session.Get("userID")
	if userIDValue == nil {
		fmt.Println("userID is not of type int")
		// Handle the case where the "userID" key is not present in the session
		return 0, errors.New("userID is not of type int")
	}

	userID, ok := userIDValue.(int)
	if !ok {
		fmt.Println("userID is not of type int 2")
		// Handle the case where the value stored for "userID" is not of type int
		return 0, errors.New("userID is not of type int")
	}

	return userID, nil
}
