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
	fmt.Println("SESSIONNNNN")
	fmt.Println(session)

	session.Set("userID", int64(user.ID))
	keys := session.Keys()
	fmt.Println("keys")
	fmt.Println(keys)
	userID, found := session.Get("userID").(int64)
	fmt.Println(found)
	fmt.Println(userID)
	if err := session.Save(); err != nil {
		// Handle the error, e.g., log it or return an internal server error
		return err
	}
	return nil
}

func GetUserIDFromSession(c *fiber.Ctx) (int64, error) {
	session := c.Locals("session").(*session.Session)
	fmt.Println("SESSION 2 CUY")
	fmt.Println(session)
	keys := session.Keys()
	fmt.Println("keys")
	fmt.Println(keys)
	userIDValue, found := session.Get("userID").(int64)
	if !found {
		fmt.Println("MASUK SINIIIIIIIIIII")
		return 0, errors.New("userID is not type of Int")
	}
	return userIDValue, nil
}
