package controllers

import "github.com/gofiber/fiber/v2"

func LandingPageController(c *fiber.Ctx) error {
	return c.Render("landingpage/index", fiber.Map{})
}
