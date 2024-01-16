package routes

import (
	"event-hunters/controllers/admin"
	"event-hunters/controllers/eventcreator"
	"event-hunters/controllers/user"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	fmt.Println("WOI")
	app.Get("/", user.LandingPageController)
	app.Get("/loginuser", user.LoginPageController)
	app.Get("/logineventcreator", eventcreator.LoginPageController)
	app.Get("/logintype", user.LoginTypePageController)
	app.Get("/registration", user.RegistrationPageController)
	app.Post("/register", user.RegistrationHandler)
	app.Get("/mainpage", user.MainPageController)
	app.Get("/mainpage/eventdetails", user.MainPageController)
	app.Get("/mainpage/recommendation", user.MainPageController)
	app.Get("/mainpage/search", user.MainPageController)
	app.Get("/mainpage/mytickets", user.MainPageController)
	app.Get("/mainpage/ticketinformation", user.MainPageController)

	app.Get("/eventcreatordashboard", eventcreator.EventCreatorDashboardController)

	app.Get("/adminpage", admin.AdminPageController)
	app.Get("/events", admin.EventsManagementController)
	app.Get("/users", admin.UsersManagementController)
	app.Get("/categories", admin.CategoriesManagementController)
	app.Get("/eventcreators", admin.EventCreatorsManagementController)
	app.Get("/tickets", admin.TicketsManagementController)
	app.Get("/ratings", admin.RatingsManagementController)
	app.Get("/eventviews", admin.EventViewsManagementController)
	app.Get("/eventbookmarks", admin.EventBookmarksManagementController)
	app.Get("/eventcategoriesrelation", admin.EventCategoriesRelationManagementController)

	app.Get("/seedsusers", admin.InitiateUsers)
	app.Get("/seedsevents", admin.InitiateEvents)
	app.Get("/seedseventcreators", admin.InitiateEventCreators)
	app.Get("/seedscategories", admin.InitiateCategories)
	app.Get("/seedseventviews", admin.InitiateEventViews)

	app.Get("/deleteallusers", admin.DeleteAllUsers)
	app.Get("/deleteallcategories", admin.DeleteAllCategories)
	app.Get("/deleteallevents", admin.DeleteAllEvents)
	app.Get("/deleteeventviews", admin.DeleteAllEventViews)

	app.Get("/regenerateevents", admin.RegenerateEvents)
	app.Get("/regenerateeventviews", admin.RegenerateEventViews)
}
