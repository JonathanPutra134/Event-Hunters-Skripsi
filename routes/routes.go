package routes

import (
	"event-hunters/controllers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	fmt.Println("WOI")
	app.Get("/", controllers.LandingPageController)
	app.Get("/login", controllers.LoginPageController)
	app.Get("/registration", controllers.RegistrationPageController)
	app.Get("/mainpage", controllers.MainPageController)
}
