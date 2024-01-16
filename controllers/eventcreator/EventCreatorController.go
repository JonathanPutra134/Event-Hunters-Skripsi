package eventcreator

import (
	"github.com/gofiber/fiber/v2"
)

func EventCreatorDashboardController(c *fiber.Ctx) error {
	return c.Render("eventcreator-dashboard/home/index", fiber.Map{})
}
func LoginPageController(c *fiber.Ctx) error {
	// userType := c.FormValue("user_type")
	// c.Locals("userType", userType)
	return c.Render("eventcreator-loginpage/index", fiber.Map{})
}
