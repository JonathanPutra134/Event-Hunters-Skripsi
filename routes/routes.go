package routes

import (
	"event-hunters/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", user.LandingPageController)
	app.Get("/loginuser", user.LoginPageController)
	app.Post("/submitlogin", user.LoginHandler)
	app.Post("/logoutuser", user.LogoutController)
	app.Get("/registration", user.RegistrationPageController)
	app.Post("/register", user.RegistrationHandler)
	app.Get("/mainpage", user.MainPageHomeController)
	app.Get("/mainpage/eventdetails", user.MainPageEventDetailsController)
	app.Get("/mainpage/recommendationpage", user.MainPageRecommendationController)
	app.Get("/mainpage/search", user.SearchHandler)
	app.Post("/mainpage/search", user.SearchHandler)
	app.Get("/mainpage/mytickets", user.MainPageMyTicketsController)
	app.Get("/mainpage/ticketinformation", user.MainPageTicketInformationController)
	app.Get("/mainpage/events/entertainment-and-performance", user.MainPageEntertainmentAndPerformanceEventsController)
	app.Get("/mainpage/events/art-and-culture", user.MainPageArtAndCultureEventsController)
	app.Get("/mainpage/events/sports", user.MainPageSportsEventsController)
	app.Get("/mainpage/events/competition", user.MainPageCompetitionEventsController)
	app.Get("/mainpage/events/charity", user.MainPageCharityEventsController)
	app.Get("/mainpage/events/education-and-career", user.MainPageEducationAndCareerEventsController)
	app.Get("/mainpage/events/expo", user.MainPageExpoEventsController)
	app.Get("/mainpage/submitrating", user.MainPageSubmitRatingController)
	app.Get("/mainpage/bookmark", user.MainPageBookmarkController)
	app.Get("/mainpage/savedevents", user.MainPageSavedEventsController)
	app.Get("/mainpage/registerticket", user.MainPageRegisterTicketController)
	app.Get("/mainpage/getrecommendations", user.MainPageRecommendationController)
}
