// session.go
package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func NewSession() *session.Store {
	return session.New()
}

func GetSession(c *fiber.Ctx, store *session.Store) (*session.Session, error) {
	return store.Get(c)
}

func SetAlertState(c *fiber.Ctx, alertType string, alertMessage string) error {
	// Retrieve the session store from the context
	store, ok := c.Locals("session_store").(*session.Store)
	if !ok {
		return fmt.Errorf("session store not found in context")
	}

	sess, err := GetSession(c, store)
	if err != nil {
		return err
	}

	sess.Set("alertType", alertType)
	sess.Set("alertMessage", alertMessage)

	// Save the session
	if err := sess.Save(); err != nil {
		// Handle the error
		return err
	}

	// Set the alert values as cookies
	c.Cookie(&fiber.Cookie{
		Name:  "alertType",
		Value: alertType,
	})
	c.Cookie(&fiber.Cookie{
		Name:  "alertMessage",
		Value: alertMessage,
	})

	fmt.Println("MASUK SET ALERT STATE")
	return nil
}
