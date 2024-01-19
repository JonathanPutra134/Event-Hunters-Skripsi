package user

import (
	"event-hunters/dto"
	"event-hunters/helpers"
	"event-hunters/repository"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func LandingPageController(c *fiber.Ctx) error {
	return c.Render("landingpage/index", fiber.Map{})
}

func LoginTypePageController(c *fiber.Ctx) error {
	return c.Render("logintype/index", fiber.Map{})
}

func LoginPageController(c *fiber.Ctx) error {
	// userType := c.FormValue("user_type")
	// c.Locals("userType", userType)

	user, err := repository.GetUserById(5)
	if err != nil {
		return err
	}

	return c.Render("loginpage/index", fiber.Map{
		"user": user,
	})
}

func LoginHandler(c *fiber.Ctx) error {
	email := c.FormValue("Email")
	password := c.FormValue("Password")
	// Retrieve the user based on the provided email
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return c.Render("registrationpage/index", fiber.Map{
			"alertType":    "danger",
			"alertMessage": "Email isn't exists",
		})
	}
	// Check if the user exists and the password matches
	if user != nil && user.Password.Valid {
		err := helpers.ComparePasswords(user.Password.String, password)
		if err != nil {
			// Passwords match, login successful
			return c.Render("registrationpage/index", fiber.Map{
				"alertType":    "danger",
				"alertMessage": "Wrong Email / Password please check again",
			})
		}
	}
	return c.Render("mainpage/index", fiber.Map{
		"user": user,
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
	latitude := c.FormValue("Latitude")
	longitude := c.FormValue("Longitude")
	exists := repository.CheckEmailExists(email)
	if exists {
		return c.Render("registrationpage/index", fiber.Map{
			"alertType":    "danger",
			"alertMessage": "Email already Registered by other user, please use another email",
		})
	}

	usrRegistrationReq := dto.UserRegistrationRequest{
		FullName:    fullName,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
		Address:     address,
		Latitude:    latitude,
		Longitude:   longitude,
	}
	errorMessages := helpers.InputNullValidation(usrRegistrationReq)
	latFloat, err := helpers.ValidateLatitude(latitude)
	//if _, err := mail.ParseAddress(email); err != nil {
	//	errorMessages = append(errorMessages, "Invalid Email")
	//}
	if err != nil {
		errorMessages = append(errorMessages, "Invalid Latitude")
	}
	lonFloat, err := helpers.ValidateLongitude(longitude)
	if err != nil {
		errorMessages = append(errorMessages, "Invalid Longitude")
	}

	if len(errorMessages) > 0 {
		errorMessage := strings.Join(errorMessages, ", ")
		fmt.Println(errorMessage)
		return c.Render("registrationpage/index", fiber.Map{
			"alertType":    "danger",
			"alertMessage": errorMessage,
		})
	}

	helpers.SanitizeInput(&usrRegistrationReq)

	err = repository.RegisterUser(usrRegistrationReq, latFloat, lonFloat)
	if err != nil {
		return c.Render("registrationpage/index", fiber.Map{
			"alertType":    "danger",
			"alertMessage": "REGISTRATION PROCESS FAILED",
		})

	}

	return c.Render("loginpage/index", fiber.Map{
		"alertType":    "success", // Corrected the key name
		"alertMessage": "Registration Successful",
	})
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
