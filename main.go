package main

import (
	"event-hunters/config"
	"event-hunters/middleware"
	"event-hunters/routes"
	"os"

	"github.com/gofiber/fiber/v2/middleware/helmet"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDatabase()
	config.SyncDB()
	// config.SeedDataCategory()
}
func main() {

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")
	app.Use(middleware.Session)
	app.Use(helmet.New(helmet.Config{
		XSSProtection: "1; mode=block",
	}))

	routes.Routes(app)

	app.Listen(":" + os.Getenv("PORT"))

	// //THIS IS FOR DOCKER
	// if err := app.Listen(":8080"); err != nil {
	// 	// Handle the error, if any
	// 	panic(err)
	// }
}
