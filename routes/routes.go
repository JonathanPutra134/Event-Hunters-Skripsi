package routes

import (
	"event-hunters/controllers/eventcreator"
	"event-hunters/controllers/user"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	fmt.Println("WOI")
	app.Get("/", user.LandingPageController)
	app.Post("/login", user.LoginPageController)
	app.Get("/logintype", user.LoginTypePageController)
	app.Get("/registration", user.RegistrationPageController)
	app.Get("/mainpage", user.MainPageController)
	app.Get("/mainpage/eventdetails", user.MainPageController)
	app.Get("/mainpage/recommendation", user.MainPageController)
	app.Get("/mainpage/search", user.MainPageController)
	app.Get("/mainpage/mytickets", user.MainPageController)
	app.Get("/mainpage/ticketinformation", user.MainPageController)

	app.Get("/eventcreatordashboard", eventcreator.EventCreatorDashboardController)
}
