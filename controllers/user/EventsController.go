package user

import (
	"event-hunters/helpers"
	"event-hunters/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

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
		return c.Render("errorpage/index", fiber.Map{"Error": err, "User": user, "BaseURL": baseURL})
	}
	viewed := repository.CheckViewExist(user.ID, event.ID)
	if !viewed {
		err = repository.InsertEventView(user.ID, eventID)
		if err != nil {
			return c.Render("errorpage/index", fiber.Map{"Error": err, "User": user, "BaseURL": baseURL})
		}
	}
	bookmarked := repository.CheckBookmarkExist(user.ID, event.ID)
	registered := repository.CheckTicketExist(user.ID, event.ID)
	return c.Render("mainpage/eventdetails/index", fiber.Map{"BaseURL": baseURL, "User": user, "Event": event, "Bookmarked": bookmarked, "Registered": registered})
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
		return c.Render("errorpage/index", fiber.Map{"Error": err, "User": user, "BaseURL": baseURL})
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
		return c.Render("errorpage/index", fiber.Map{"Error": err, "User": user, "BaseURL": baseURL})
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
		return c.Render("errorpage/index", fiber.Map{"Error": err, "User": user, "BaseURL": baseURL})
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
		return c.Render("errorpage/index", fiber.Map{"Error": err, "User": user, "BaseURL": baseURL})
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
		return c.Render("errorpage/index", fiber.Map{"Error": err, "User": user, "BaseURL": baseURL})
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
	events, err := repository.GetFilteredEventsByCategory("Sports")
	if err != nil {
		return c.Render("errorpage/index", fiber.Map{"Error": err, "User": user, "BaseURL": baseURL})
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
		return c.Render("errorpage/index", fiber.Map{"Error": err, "User": user, "BaseURL": baseURL})
	}
	return c.Render("mainpage/home/categories/expo", fiber.Map{"BaseURL": baseURL, "User": user, "Events": events, "Truncate": helpers.Truncate, "Category": "expo"})
}
