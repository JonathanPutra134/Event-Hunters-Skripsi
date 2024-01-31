package user

import (
	"event-hunters/dto"
	"event-hunters/helpers"
	"event-hunters/repository"
	"fmt"
	"log"
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
			"alertMessage": "Email isn't exists",
		})
	}
	// Check if the user exists and the password matches
	if user != nil && user.Email.String != email && user.Password.Valid {
		err := helpers.ComparePasswords(user.Password.String, password)
		if err != nil {
			// Passwords match, login successful
			return c.Render("loginpage/index", fiber.Map{
				"alertType":    "danger",
				"alertMessage": "Wrong Email / Password please check again",
			})
		}
	}
	sessionID, err := repository.StoreSession(user)
	if err != nil {
		log.Fatal(err)
	}
	// Set the sessionID as a cookie
	c.Cookie(&fiber.Cookie{
		Name:  "sessionID",
		Value: sessionID,
		// You can set other cookie options, such as Secure, HTTPOnly, etc.
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

func MainPageController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	urlPath := c.Path()

	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	pathTemplates := map[string]string{
		"/mainpage/eventdetails":      "mainpage/eventdetails/index",
		"/mainpage/recommendation":    "mainpage/recommendation/index",
		"/mainpage/search":            "mainpage/search/index",
		"/mainpage/mytickets":         "mainpage/mytickets/index",
		"/mainpage/ticketinformation": "mainpage/ticketinformation/index",
		"/mainpage":                   "mainpage/home/index",
	}

	templatePath, exists := pathTemplates[urlPath]
	if !exists {
		templatePath = "mainpage/home/index"
	}

	return c.Render(templatePath, fiber.Map{"BaseURL": baseURL, "User": user})
}

func LogoutController(c *fiber.Ctx) error {
	// Retrieve session ID from cookie
	sessionID := c.Cookies("sessionID")
	// Check if session ID exists
	err := repository.DeleteUserSession(sessionID)
	if err != nil {
		log.Fatal("Error deleting user session")
	}

	c.ClearCookie("sessionID")
	return c.Redirect("/loginuser?alertType=success&alertMessage=Logout successful", http.StatusSeeOther)
}

func MainPageHomeController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}

	latestEvents, _ := repository.GetLatestEvent(5)
	fmt.Println(latestEvents[0].Title.String)
	return c.Render("mainpage/home/index", fiber.Map{
		"BaseURL": baseURL,
		"User":    user,
		"Events":  latestEvents,
	})
}

func MainPageEventDetailsController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}

	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}

	eventID := c.Query("id")

	event, err := repository.GetEventById(eventID)
	if err != nil {
		log.Fatal(err)
	}
	return c.Render("mainpage/eventdetails/index", fiber.Map{"BaseURL": baseURL, "Finished": false, "User": user, "Event": event})
}

func MainPageRecommendationController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	return c.Render("mainpage/recommendation/index", fiber.Map{"BaseURL": baseURL, "Finished": false, "User": user})
}

func SearchHandler(c *fiber.Ctx) error {
	// Extract form values
	keyword := c.FormValue("Keyword")
	category := c.FormValue("categories[]")
	minRegDate := c.FormValue("MinRegDate")
	maxRegDate := c.FormValue("MaxRegDate")
	minEventStartDate := c.FormValue("MinEventStartDate")
	maxEventStartDate := c.FormValue("MaxEventStartDate")
	eventType := c.FormValue("EventType")
	fmt.Println("MASUK SEARCH HANDLER")
	fmt.Println(keyword)
	fmt.Println(category)
	fmt.Println(minRegDate)
	fmt.Println(maxRegDate)
	fmt.Println(minEventStartDate)
	fmt.Println(maxEventStartDate)
	fmt.Println(eventType)

	return c.Redirect("/loginuser?alertType=success&alertMessage=Registration+Successful", http.StatusSeeOther)
	// return c.Render("loginpage/index", fiber.Map{
	// 	"alertType":    "success", // Corrected the key name
	// 	"alertMessage": "Registration Successful",
	// })
}

func MainPageSearchController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	return c.Render("mainpage/search/index", fiber.Map{"BaseURL": baseURL, "Finished": false, "User": user})
}

func MainPageMyTicketsController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	return c.Render("mainpage/mytickets/index", fiber.Map{"BaseURL": baseURL, "Finished": false, "User": user})
}

func MainPageTicketInformationController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	return c.Render("mainpage/ticketinformation/index", fiber.Map{"BaseURL": baseURL, "Finished": false, "User": user})
}

func MainPageEntertainmentAndPerformanceEventsController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	events, err := repository.GetFilteredEventsByCategory("Entertainment & Performance")
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	return c.Render("mainpage/home/categories/entertainment-and-performance", fiber.Map{"BaseURL": baseURL, "User": user, "Events": events, "Truncate": helpers.Truncate})
}

func MainPageArtAndCultureEventsController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	events, err := repository.GetFilteredEventsByCategory("Art & Culture")
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	return c.Render("mainpage/home/categories/art-and-culture", fiber.Map{"BaseURL": baseURL, "User": user, "Events": events, "Truncate": helpers.Truncate})
}

func MainPageCharityEventsController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	events, err := repository.GetFilteredEventsByCategory("Charity")
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	return c.Render("mainpage/home/categories/charity", fiber.Map{"BaseURL": baseURL, "User": user, "Events": events, "Truncate": helpers.Truncate})
}

func MainPageCompetitionEventsController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	events, err := repository.GetFilteredEventsByCategory("Competition")
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	return c.Render("mainpage/home/categories/competition", fiber.Map{"BaseURL": baseURL, "User": user, "Events": events, "Truncate": helpers.Truncate})
}

func MainPageEducationAndCareerEventsController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	events, err := repository.GetFilteredEventsByCategory("Education & Career")
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	return c.Render("mainpage/home/categories/education-and-career", fiber.Map{"BaseURL": baseURL, "User": user, "Events": events, "Truncate": helpers.Truncate})
}

func MainPageSportsEventsController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	events, err := repository.GetFilteredEventsByCategory("Sport")
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	return c.Render("mainpage/home/categories/sports", fiber.Map{"BaseURL": baseURL, "User": user, "Events": events, "Truncate": helpers.Truncate})
}

func MainPageExpoEventsController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	events, err := repository.GetFilteredEventsByCategory("Expo")
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	return c.Render("mainpage/home/categories/expo", fiber.Map{"BaseURL": baseURL, "User": user, "Events": events, "Truncate": helpers.Truncate, "Category": "expo"})
}
