package user

import (
	"context"
	"event-hunters/config"
	"event-hunters/helpers"
	"event-hunters/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func LandingPageController(c *fiber.Ctx) error {
	return c.Render("landingpage/index", fiber.Map{})
}

func LoginTypePageController(c *fiber.Ctx) error {
	return c.Render("logintype/index", fiber.Map{})
}

func LoginPageController(c *fiber.Ctx) error {
	userType := c.FormValue("user_type")
	c.Locals("userType", userType)

	return c.Render("loginpage/index", fiber.Map{
		"userType": userType,
	})
}

func RegistrationPageController(c *fiber.Ctx) error {
	return c.Render("registrationpage/index", fiber.Map{})
}

func RegistrationHandler(c *fiber.Ctx) error {
	// Extract form values
	fullName := c.FormValue("FullName")
	email := c.FormValue("Email")
	password := c.FormValue("Password")
	phoneNumber := c.FormValue("PhoneNumber")
	address := c.FormValue("Address")

	// Perform any additional validation on the form data as needed

	// Create a new user with the extracted data
	newUser := models.User{
		Name:        null.StringFrom(fullName),
		Email:       null.StringFrom(email),
		Password:    helpers.HashPassword(password),
		PhoneNumber: null.StringFrom(phoneNumber),
		Address:     null.StringFrom(address),
	}

	err := newUser.Insert(context.Background(), config.DB, boil.Infer())
	if err != nil {
		fmt.Println("Error registering user:", err)
		log.Fatal(err)
	}
	fmt.Printf("User %s inserted successfully.\n", newUser.Name.String)

	// Redirect to a success page or send a success response
	return c.Redirect("/login")
}

func MainPageController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	urlPath := c.Path()
	if urlPath == "/mainpage/eventdetails" {
		return c.Render("mainpage/eventdetails/index", fiber.Map{"BaseURL": baseURL, "Finished": false})
	}
	if urlPath == "/mainpage/recommendation" {
		return c.Render("mainpage/recommendation/index", fiber.Map{"BaseURL": baseURL})
	}
	if urlPath == "/mainpage/search" {
		return c.Render("mainpage/search/index", fiber.Map{"BaseURL": baseURL})
	}
	if urlPath == "/mainpage/mytickets" {
		return c.Render("mainpage/mytickets/index", fiber.Map{"BaseURL": baseURL})
	}
	if urlPath == "/mainpage/ticketinformation" {
		return c.Render("mainpage/ticketinformation/index", fiber.Map{"BaseURL": baseURL, "Finished": true})
	}
	fmt.Println("MASUK SINI NIH BOS")
	return c.Render("mainpage/home/index", fiber.Map{
		"BaseURL": baseURL,
	})
}

func MainPageEventDetailsController(c *fiber.Ctx) error {
	return c.Render("mainpage/home/index", fiber.Map{})
}
