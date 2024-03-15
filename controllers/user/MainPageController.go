package user

import (
	"encoding/json"
	"errors"
	"event-hunters/dto"
	"event-hunters/repository"
	"fmt"
	"net/http"
	"strconv"

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
	userID := c.Query("id")
	if userID == "" {
		return c.Render("mainpage/recommendation/index", fiber.Map{"BaseURL": baseURL, "Finished": false, "User": user})
	}

	if userID != strconv.Itoa(user.ID) {
		return c.Render("errorpage/index", fiber.Map{"Error": errors.New("you cannot get other users recommendations"), "User": user, "BaseURL": baseURL})
	}
	url := fmt.Sprintf("https://event-hunters-recommender-production.up.railway.app/recommend/%s", userID)

	// Make HTTP request
	response, err := http.Get(url)
	if err != nil {
		return c.Render("errorpage/index", fiber.Map{"Error": err, "User": user, "BaseURL": baseURL})
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusBadRequest {
		// Handle 400 Bad Request
		return c.Render("errorpage/index", fiber.Map{"Error": errors.New("interaksi dari user tidak diketahui / Interaksimu belum cukup untuk diberikan suatu rekomendasi, dimohon untuk melakukan interaksi terhadap item event yang disediakan (minimal 5 dari bookmark, rating, view ataupun registrasi event) "), "User": user, "BaseURL": baseURL})
	} else if response.StatusCode != http.StatusOK {
		// Handle other status codes
		return c.Render("errorpage/index", fiber.Map{"Error": errors.New("unexpected error"), "User": user, "BaseURL": baseURL})
	}

	// Parse the response body into a RecommendationsResponse struct
	var recommendationsRes dto.RecommendationsResponse
	if err := json.NewDecoder(response.Body).Decode(&recommendationsRes); err != nil {
		return c.Render("errorpage/index", fiber.Map{"Error": err, "User": user, "BaseURL": baseURL})
	}
	recommendedEvents, err := repository.GetRecommendedEvents(recommendationsRes)
	if err != nil {
		return c.Render("errorpage/index", fiber.Map{"Error": err, "User": user, "BaseURL": baseURL})
	}
	return c.Render("mainpage/recommendation/index", fiber.Map{"BaseURL": baseURL, "Finished": false, "User": user, "Recommendations": recommendedEvents})
}
