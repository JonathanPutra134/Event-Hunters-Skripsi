package user

import (
	"event-hunters/repository"
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

	latestEvents, err := repository.GetLatestEvent(5)
	if err != nil {
		return c.Render("errorpage/index", fiber.Map{"Error": err, "User": user, "BaseURL": baseURL})
	}
	return c.Render("mainpage/home/index", fiber.Map{
		"BaseURL": baseURL,
		"User":    user,
		"Events":  latestEvents,
	})
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
