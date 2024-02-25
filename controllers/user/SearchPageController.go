package user

import (
	"event-hunters/helpers"
	"event-hunters/repository"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SearchHandler(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	sessionID := c.Cookies("sessionID")
	if sessionID == "" {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	user, err := repository.GetUserBySessionID(sessionID)
	if err != nil {
		return c.Redirect("/loginuser?alertType=danger&alertMessage=Please Login Again", http.StatusSeeOther)
	}
	// Extract form values
	keyword := c.FormValue("Keyword")
	categories := helpers.CategoriesFormHandler(c)
	minRegDate := c.FormValue("MinRegDate")
	maxRegDate := c.FormValue("MaxRegDate")
	minEventStartDate := c.FormValue("MinEventStartDate")
	maxEventStartDate := c.FormValue("MaxEventStartDate")
	fmt.Println("EVENT DATE FORM VALUE")
	fmt.Println(minRegDate)
	fmt.Println(maxRegDate)

	ParsedSearchDate, err := helpers.ParseSearchDate(minRegDate, maxRegDate, minEventStartDate, maxEventStartDate)
	if err != nil {
		return err
	}
	err = helpers.DateValidation(ParsedSearchDate)
	if err != nil {
		return err
	}
	eventType := c.FormValue("EventType")
	events, err := repository.GetSearchedEvents(keyword, categories, ParsedSearchDate, eventType)
	if err != nil {
		return err
	}

	// for i, event := range events {
	// 	fmt.Println("Event", i)
	// 	fmt.Println(event.Title)
	// 	fmt.Println(event.PreregisterDate)
	// 	fmt.Println(event.StarteventDate)
	// }
	return c.Render("mainpage/search/index", fiber.Map{"BaseURL": baseURL, "Finished": false, "User": user, "Events": events})
}
