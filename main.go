package main

import (
	"event-hunters/config"
	"event-hunters/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDatabase()
	config.SyncDB()
}
func main() {

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	routes.Routes(app)
	app.Listen(":" + os.Getenv("PORT"))
}
