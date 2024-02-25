package user

import (
	"errors"
	"event-hunters/repository"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

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
	bookmarked := repository.CheckBookmarkExist(user.ID, event.ID)
	registered := repository.CheckTicketExist(user.ID, event.ID)
	return c.Render("mainpage/eventdetails/index", fiber.Map{"BaseURL": baseURL, "Finished": false, "User": user, "Event": event, "Bookmarked": bookmarked, "Registered": registered})
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

func MainPageSavedEventsController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	userID := user.ID
	exist := repository.CheckUserExistInBookmark(userID)
	if !exist {
		fmt.Println("THERE IS NO SAVED EVENTS")
		return errors.New("There is nothing to show here")
	}
	events, err := repository.GetSavedEvents(userID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	return c.Render("mainpage/savedevents/index", fiber.Map{"BaseURL": baseURL, "Events": events, "User": user})
}
