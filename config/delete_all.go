package config

import (
	"context"
	"event-hunters/models"

	"fmt"
	"log"
)

func DeleteUsers() error {
	fmt.Println("Deleting all users...")

	// Delete all records from the users table
	_, err := models.Users().DeleteAll(context.Background(), DB)
	if err != nil {
		fmt.Println("Error deleting users:", err)
		log.Fatal(err)
	}

	fmt.Println("All users deleted successfully.")
	return nil
}

func DeleteCategories() error {
	fmt.Println("Deleting all categories..")

	// Delete all records from the users table
	_, err := models.Categories().DeleteAll(context.Background(), DB)
	if err != nil {
		fmt.Println("Error deleting categories:", err)
		log.Fatal(err)
	}

	fmt.Println("All categories deleted successfully.")
	return nil
}

func DeleteEvents() error {
	fmt.Println("Deleting all Events..")

	// Delete all records from the users table
	_, err := models.Events().DeleteAll(context.Background(), DB)
	if err != nil {
		fmt.Println("Error deleting Events:", err)
		log.Fatal(err)
	}

	fmt.Println("All Events deleted successfully.")
	return nil
}

func DeleteEventCreators() error {
	fmt.Println("Deleting all Events Creators..")

	// Delete all records from the users table
	_, err := models.EventCreators().DeleteAll(context.Background(), DB)
	if err != nil {
		fmt.Println("Error deleting Events:", err)
		log.Fatal(err)
	}

	fmt.Println("All Events deleted successfully.")
	return nil
}
func DeleteEventViewsInteraction() error {
	fmt.Println("Deleting all Events Views Interaction..")

	// Delete all records from the users table
	_, err := models.EventsViews().DeleteAll(context.Background(), DB)
	if err != nil {
		fmt.Println("Error deleting Events Views:", err)
		log.Fatal(err)
	}

	fmt.Println("All Events Views deleted successfully.")
	return nil
}
func DeleteEventBookmarksInteraction() error {
	fmt.Println("Deleting all Events Bookmarks Interaction..")

	// Delete all records from the users table
	_, err := models.EventsBookmarks().DeleteAll(context.Background(), DB)
	if err != nil {
		fmt.Println("Error deleting Events Bookmarks:", err)
		log.Fatal(err)
	}

	fmt.Println("All Events Bookmarks deleted successfully.")
	return nil
}

func DeleteEventsCategoriesRelation() error {
	fmt.Println("Deleting all Events Categories Relation..")

	// Delete all records from the users table
	_, err := models.EventsCategories().DeleteAll(context.Background(), DB)
	if err != nil {
		fmt.Println("Error deleting Categories Relation:", err)
		log.Fatal(err)
	}

	fmt.Println("All Categories Relation deleted successfully.")
	return nil
}

func DeleteTickets() error {
	fmt.Println("Deleting all Tickets..")

	// Delete all records from the users table
	_, err := models.Tickets().DeleteAll(context.Background(), DB)
	if err != nil {
		fmt.Println("Error deleting Tickets:", err)
		log.Fatal(err)
	}

	fmt.Println("All Tickets deleted successfully.")
	return nil
}

func DeleteRatings() error {
	fmt.Println("Deleting all Ratings..")

	// Delete all records from the users table
	_, err := models.Ratings().DeleteAll(context.Background(), DB)
	if err != nil {
		fmt.Println("Error deleting Ratings:", err)
		log.Fatal(err)
	}

	fmt.Println("All Ratings deleted successfully.")
	return nil
}
