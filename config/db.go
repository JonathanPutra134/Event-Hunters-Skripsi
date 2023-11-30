package config

import (
	"context"
	"database/sql"
	"event-hunters/models"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var DB *sql.DB

func ConnectToDatabase() {
	fmt.Println("MASUK DATABASE NIHHH")
	var err error
	dsn := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		fmt.Println("MASUK ERROR PAS CONNECT DATABASE NIHH")
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	DB = db
}

func SyncDB() {
	// err := DB.AutoMigrate(&models.User{})
	// if err != nil {
	// 	fmt.Printf("%#v\n", &models.User{})
	// 	fmt.Println("MASUK DI ERROR AUTO MIGRATE")
	// 	log.Fatal(err)
	// }
	// SeedDataUsers()
	SeedDataCategory()
	// DeleteUsers()
	// DeleteCategories()
}

func SeedDataUsers() {
	count, err := models.Users().Count(context.Background(), DB)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		usersToInsert := []models.User{
			{
				Name:           null.StringFrom("John Doe"),
				BirthDate:      null.TimeFrom(time.Date(1990, time.January, 15, 0, 0, 0, 0, time.UTC)),
				Longitude:      null.StringFrom("-73.985428"),
				Latitude:       null.StringFrom("40.748817"),
				Email:          null.StringFrom("john.doe@example.com"),
				Password:       null.StringFrom("securepassword1"),
				ProfilePicture: null.StringFrom("john_doe.jpg"),
				PhoneNumber:    null.StringFrom("+1234567890"),
			},
			{
				Name:           null.StringFrom("Jonggun"),
				BirthDate:      null.TimeFrom(time.Date(2002, time.January, 15, 0, 0, 0, 0, time.UTC)),
				Longitude:      null.StringFrom("-73.985428"),
				Latitude:       null.StringFrom("40.748817"),
				Email:          null.StringFrom("john.doe@example.com"),
				Password:       null.StringFrom("securepassword1"),
				ProfilePicture: null.StringFrom("john_doe.jpg"),
				PhoneNumber:    null.StringFrom("+1234567890"),
			},
			{
				Name:           null.StringFrom("usagi_coser"),
				BirthDate:      null.TimeFrom(time.Date(2000, time.January, 15, 0, 0, 0, 0, time.UTC)),
				Longitude:      null.StringFrom("-73.985428"),
				Latitude:       null.StringFrom("40.748817"),
				Email:          null.StringFrom("john.doe@example.com"),
				Password:       null.StringFrom("securepassword1"),
				ProfilePicture: null.StringFrom("john_doe.jpg"),
				PhoneNumber:    null.StringFrom("+1234567890"),
			},
			// Add more users as needed
		}
		for _, user := range usersToInsert {
			err := user.Insert(context.Background(), DB, boil.Infer())
			if err != nil {
				fmt.Println("Error creating user:", err)
				log.Fatal(err)
			}
			fmt.Printf("User %s inserted successfully.\n", user.Name.String)
		}

	} else {
		fmt.Println("Initial users already exist, seeding process will not be executed")
	}

}

func SeedDataCategory() {
	count, err := models.Categories().Count(context.Background(), DB)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		categoriesToInsert := []models.Category{
			{
				Name: null.StringFrom("Education and Career"),
			},
			{
				Name: null.StringFrom("Entertainment & Performance"),
			},
			{
				Name: null.StringFrom("Travel & Outdoor"),
			},
			{
				Name: null.StringFrom("Charity"),
			},
			{
				Name: null.StringFrom("Relaxation"),
			},
			{
				Name: null.StringFrom("Sport"),
			},
			{
				Name: null.StringFrom("Art & Culture"),
			},
			// Add more users as needed
		}
		for _, category := range categoriesToInsert {
			err := category.Insert(context.Background(), DB, boil.Infer())
			if err != nil {
				fmt.Println("Error creating category:", err)
				log.Fatal(err)
			}
			fmt.Printf("Category %s inserted successfully.\n", category.Name.String)
		}

	} else {
		fmt.Println("Initial categories already exist, seeding process will not be executed")
	}

}

func DeleteUsers() {
	fmt.Println("Deleting all users...")

	// Delete all records from the users table
	_, err := models.Users().DeleteAll(context.Background(), DB)
	if err != nil {
		fmt.Println("Error deleting users:", err)
		log.Fatal(err)
	}

	fmt.Println("All users deleted successfully.")
}

func DeleteCategories() {
	fmt.Println("Deleting all categories..")

	// Delete all records from the users table
	_, err := models.Categories().DeleteAll(context.Background(), DB)
	if err != nil {
		fmt.Println("Error deleting categories:", err)
		log.Fatal(err)
	}

	fmt.Println("All categories deleted successfully.")
}
