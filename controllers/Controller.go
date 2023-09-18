package controllers

import "github.com/gofiber/fiber/v2"

func LandingPageController(c *fiber.Ctx) error {
	return c.Render("landingpage/index", fiber.Map{})
}

func LoginPageController(c *fiber.Ctx) error {

	return c.Render("loginpage/index", fiber.Map{})
}
