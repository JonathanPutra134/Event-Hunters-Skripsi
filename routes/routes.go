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
	app.Post("/submitlogin", user.LoginHandler)
	app.Post("/logoutuser", user.LogoutController)
	app.Get("/logineventcreator", eventcreator.LoginPageController)
	app.Get("/logintype", user.LoginTypePageController)
	app.Get("/registration", user.RegistrationPageController)
	app.Post("/register", user.RegistrationHandler)
	app.Get("/mainpage", user.MainPageHomeController)
	app.Get("/mainpage/eventdetails", user.MainPageEventDetailsController)
	app.Get("/mainpage/recommendation", user.MainPageRecommendationController)
	app.Get("/mainpage/search", user.MainPageSearchController)
	app.Get("/mainpage/mytickets", user.MainPageMyTicketsController)
	app.Get("/mainpage/ticketinformation", user.MainPageTicketInformationController)
	app.Get("/mainpage/events/entertainment-and-performance", user.MainPageEntertainmentAndPerformanceEventsController)
	app.Get("/mainpage/events/art-and-culture", user.MainPageArtAndCultureEventsController)
	app.Get("/mainpage/events/sports", user.MainPageSportsEventsController)
	app.Get("/mainpage/events/competition", user.MainPageCompetitionEventsController)
	app.Get("/mainpage/events/charity", user.MainPageCharityEventsController)
	app.Get("/mainpage/events/education-and-career", user.MainPageEducationAndCareerEventsController)
	app.Get("/mainpage/events/expo", user.MainPageExpoEventsController)

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
