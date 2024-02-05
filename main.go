package main

import (
	"event-hunters/config"
	"event-hunters/helpers"
	"event-hunters/middleware"
	"event-hunters/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDatabase()
	config.SyncDB()

}
func main() {

	engine := html.New("./views", ".html")
	engine.AddFunc("Truncate", helpers.Truncate)
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	fh, err := os.Create("debug.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	// Enable debug mode
	boil.DebugMode = true

	// Set the debug writer
	boil.DebugWriter = fh
	app.Static("/", "./public")
	app.Use(middleware.Session)
	// app.Use(helmet.New(helmet.Config{
	// 	XSSProtection: "1; mode=block",
	// }))

	routes.Routes(app)

	app.Listen(":" + os.Getenv("PORT"))

	// //THIS IS FOR DOCKER
	// if err := app.Listen(":8080"); err != nil {
	// 	// Handle the error, if any
	// 	panic(err)
	// }
}
