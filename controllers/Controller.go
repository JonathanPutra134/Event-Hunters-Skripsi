package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func LandingPageController(c *fiber.Ctx) error {
	return c.Render("landingpage/index", fiber.Map{})
}

func LoginPageController(c *fiber.Ctx) error {
	return c.Render("loginpage/index", fiber.Map{})
}
func RegistrationPageController(c *fiber.Ctx) error {
	return c.Render("registrationpage/index", fiber.Map{})
}

func MainPageController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	urlPath := c.Path()
	if urlPath == "/mainpage/eventdetails" {
		return c.Render("mainpage/eventdetails/index", fiber.Map{"BaseURL": baseURL})
	}
	if urlPath == "/mainpage/recommendation" {
		return c.Render("mainpage/recommendation/index", fiber.Map{"BaseURL": baseURL})
	}
	if urlPath == "/mainpage/search" {
		return c.Render("mainpage/search/index", fiber.Map{"BaseURL": baseURL})
	}
	fmt.Println("MASUK SINI NIH BOS")
	return c.Render("mainpage/home/index", fiber.Map{
		"BaseURL": baseURL,
	})
}

func MainPageEventDetailsController(c *fiber.Ctx) error {
	return c.Render("mainpage/home/index", fiber.Map{})
}
