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
