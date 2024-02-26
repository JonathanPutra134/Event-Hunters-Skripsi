package user

import (
	"event-hunters/dto"
	"event-hunters/helpers"
	"event-hunters/repository"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func LandingPageController(c *fiber.Ctx) error {
	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Render("landingpage/index", fiber.Map{})
	}
	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil && user == nil {
		return c.Render("landingpage/index", fiber.Map{})
	}
	return c.Redirect("/mainpage", http.StatusSeeOther)
}

func LoginTypePageController(c *fiber.Ctx) error {
	return c.Render("logintype/index", fiber.Map{})
}

func LoginPageController(c *fiber.Ctx) error {
	// userType := c.FormValue("user_type")
	// c.Locals("userType", userType)
	return c.Render("loginpage/index", fiber.Map{})
}

func LoginHandler(c *fiber.Ctx) error {
	email := c.FormValue("Email")
	password := c.FormValue("Password")
	// Retrieve the user based on the provided email
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return c.Render("loginpage/index", fiber.Map{
			"alertType":    "danger",
			"alertMessage": "Wrong Email / Password please check again",
		})
	}
	err = helpers.ComparePasswords(user.Password.String, password)
	if err != nil {
		// Passwords match, login successful
		return c.Render("loginpage/index", fiber.Map{
			"alertType":    "danger",
			"alertMessage": "Wrong Email / Password please check again",
		})
	}
	sessionID, err := repository.StoreSession(user)
	if err != nil {
		return c.Render("loginpage/index", fiber.Map{
			"alertType":    "danger",
			"alertMessage": "Login failed. Please try again later.",
		})
	}
	// Set the sessionID as a cookie
	c.Cookie(&fiber.Cookie{
		Name:     "sessionID",
		Value:    sessionID,
		Secure:   true,     // Set to true if serving over HTTPS
		HTTPOnly: true,     // Reduce the risk of XSS attacks
		SameSite: "Strict", // Strict is recommended for better security
	})
	return c.Redirect("/mainpage", http.StatusSeeOther)
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

	return c.Redirect("/loginuser?alertType=success&alertMessage=Registration+Successful", http.StatusSeeOther)
	// return c.Render("loginpage/index", fiber.Map{
	// 	"alertType":    "success", // Corrected the key name
	// 	"alertMessage": "Registration Successful",
	// })
}

func LogoutController(c *fiber.Ctx) error {
	// Retrieve session ID from cookie
	sessionID := c.Cookies("sessionID")
	// Check if session ID exists
	err := repository.DeleteUserSession(sessionID)
	if err != nil {
		return c.Render("errorpage/index", fiber.Map{"Error": err})
	}

	c.ClearCookie("sessionID")
	return c.Redirect("/loginuser?alertType=success&alertMessage=Logout successful", http.StatusSeeOther)
}
