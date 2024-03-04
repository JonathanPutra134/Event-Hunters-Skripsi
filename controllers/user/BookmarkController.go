package user

import (
	"event-hunters/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func MainPageBookmarkController(c *fiber.Ctx) error {
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
	userID := user.ID
	err = repository.InsertEventBookmark(userID, eventID)
	if err != nil {
		return c.Render("errorpage/index", fiber.Map{"Error": err, "User": user, "BaseURL": baseURL})
	}

	event, err := repository.GetEventById(eventID)
	if err != nil {
		return c.Render("errorpage/index", fiber.Map{"Error": err, "User": user, "BaseURL": baseURL})
	}
	return c.Render("mainpage/bookmarkevent/index", fiber.Map{"BaseURL": baseURL, "Event": event, "User": user})
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
	events, err := repository.GetSavedEvents(userID)
	if err != nil {
		return c.Render("errorpage/index", fiber.Map{"Error": err, "User": user, "BaseURL": baseURL})
	}
	return c.Render("mainpage/savedevents/index", fiber.Map{"BaseURL": baseURL, "Events": events, "User": user})
}
