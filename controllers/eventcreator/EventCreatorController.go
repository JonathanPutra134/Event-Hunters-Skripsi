package eventcreator

import (
	"github.com/gofiber/fiber/v2"
)

func EventCreatorDashboardController(c *fiber.Ctx) error {
	return c.Render("eventcreator-dashboard/home/index", fiber.Map{})
}
