package admin

import (
	"event-hunters/config"
	"log"

	"github.com/gofiber/fiber/v2"
)

func AdminPageController(c *fiber.Ctx) error {
	return c.Render("adminpage/index", fiber.Map{})
}
func UsersManagementController(c *fiber.Ctx) error {
	return c.Render("adminpage/users_management/index", fiber.Map{})
}
func EventsManagementController(c *fiber.Ctx) error {
	return c.Render("adminpage/events_management/index", fiber.Map{})
}
func CategoriesManagementController(c *fiber.Ctx) error {
	return c.Render("adminpage/categories_management/index", fiber.Map{})
}
func EventCreatorsManagementController(c *fiber.Ctx) error {
	return c.Render("adminpage/eventscreator_management/index", fiber.Map{})
}
func TicketsManagementController(c *fiber.Ctx) error {
	return c.Render("adminpage/tickets_management/index", fiber.Map{})
}
func RatingsManagementController(c *fiber.Ctx) error {
	return c.Render("adminpage/ratings_management/index", fiber.Map{})
}
func EventViewsManagementController(c *fiber.Ctx) error {
	return c.Render("adminpage/events_views_management/index", fiber.Map{})
}
func EventBookmarksManagementController(c *fiber.Ctx) error {
	return c.Render("adminpage/events_bookmarks_management/index", fiber.Map{})
}
func EventCategoriesRelationManagementController(c *fiber.Ctx) error {
	return c.Render("adminpage/events_categories_relation_management/index", fiber.Map{})
}

func DeleteAllUsers(c *fiber.Ctx) error {
	err := config.DeleteUsers()

	if err != nil {
		log.Fatal(err)
	}
	return err
}

func DeleteAllCategories(c *fiber.Ctx) error {
	err := config.DeleteCategories()

	if err != nil {
		log.Fatal(err)
	}
	return err
}

func DeleteAllEvents(c *fiber.Ctx) error {
	err := config.DeleteEvents()

	if err != nil {
		log.Fatal(err)
	}
	return err
}

func InitiateEvents(c *fiber.Ctx) error {
	err := config.SeedEvents()

	if err != nil {
		log.Fatal(err)
	}
	return err
}

func InitiateUsers(c *fiber.Ctx) error {
	err := config.SeedDataUsers()

	if err != nil {
		log.Fatal(err)
	}
	return err
}

func InitiateEventCreators(c *fiber.Ctx) error {
	err := config.SeedEventCreators()

	if err != nil {
		log.Fatal(err)
	}
	return err
}

func InitiateCategories(c *fiber.Ctx) error {
	err := config.SeedDataCategory()

	if err != nil {
		log.Fatal(err)
	}
	return err
}
