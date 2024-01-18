package main

import (
	"event-hunters/config"
	"event-hunters/middleware"
	"event-hunters/routes"
	"os"

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
	// Create a new session store
	// store := middleware.NewSession()
	app.Static("/", "./public")
	// app.Use(func(c *fiber.Ctx) error {
	// 	c.Locals("session_store", store)
	// 	return c.Next()
	// })
	app.Use(middleware.XSSMiddleware)
	routes.Routes(app)

	app.Listen(":" + os.Getenv("PORT"))

	// //THIS IS FOR DOCKER
	// if err := app.Listen(":8080"); err != nil {
	// 	// Handle the error, if any
	// 	panic(err)
	// }
}
