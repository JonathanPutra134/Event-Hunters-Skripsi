package config

import (
	"context"

	"fmt"
	"log"
	"time"

	"event-hunters/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func SeedDataUsers() error {
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
		return nil
	}
	return nil
}

func SeedDataCategory() error {
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
		return nil
	}
	return nil
}

func SeedEvents() error {
	count, err := models.Events().Count(context.Background(), DB)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		EventsToInsert := []models.Event{
			{
				EventcreatorID:  null.NewInt(1, true),
				PreregisterDate: null.NewTime(time.Now(), true),
				EndregisterDate: null.NewTime(time.Now().AddDate(0, 0, 7), true),
				StartEvent:      null.NewTime(time.Now().AddDate(0, 0, 10), true),
				EndEvent:        null.NewTime(time.Now().AddDate(0, 0, 12), true),
				CreatedAt:       null.NewTime(time.Now(), true),
				UpdatedAt:       null.NewTime(time.Now(), true),
				Latitude:        null.NewString("40.7128", true),
				Longitude:       null.NewString("-74.0060", true),
				Title:           null.NewString("Sample Event", true),
				Description:     null.NewString("This is a sample event description.", true),
				AverageRating:   null.NewFloat64(4.5, true),
				IsFinished:      null.NewBool(false, true),
				IsOnline:        null.NewBool(true, true),
			},
			{
				EventcreatorID:  null.NewInt(1, true),
				PreregisterDate: null.NewTime(time.Now(), true),
				EndregisterDate: null.NewTime(time.Now().AddDate(0, 0, 7), true),
				StartEvent:      null.NewTime(time.Now().AddDate(0, 0, 10), true),
				EndEvent:        null.NewTime(time.Now().AddDate(0, 0, 12), true),
				CreatedAt:       null.NewTime(time.Now(), true),
				UpdatedAt:       null.NewTime(time.Now(), true),
				Latitude:        null.NewString("40.7128", true),
				Longitude:       null.NewString("-74.0060", true),
				Title:           null.NewString("Sample Event", true),
				Description:     null.NewString("This is a sample event description.", true),
				AverageRating:   null.NewFloat64(4.5, true),
				IsFinished:      null.NewBool(false, true),
				IsOnline:        null.NewBool(true, true),
			},
			{
				EventcreatorID:  null.NewInt(2, true),
				PreregisterDate: null.NewTime(time.Now(), true),
				EndregisterDate: null.NewTime(time.Now().AddDate(0, 0, 7), true),
				StartEvent:      null.NewTime(time.Now().AddDate(0, 0, 10), true),
				EndEvent:        null.NewTime(time.Now().AddDate(0, 0, 12), true),
				CreatedAt:       null.NewTime(time.Now(), true),
				UpdatedAt:       null.NewTime(time.Now(), true),
				Latitude:        null.NewString("40.7128", true),
				Longitude:       null.NewString("-74.0060", true),
				Title:           null.NewString("Sample Event", true),
				Description:     null.NewString("This is a sample event description.", true),
				AverageRating:   null.NewFloat64(4.5, true),
				IsFinished:      null.NewBool(false, true),
				IsOnline:        null.NewBool(true, true),
			},
			{
				EventcreatorID:  null.NewInt(2, true),
				PreregisterDate: null.NewTime(time.Now(), true),
				EndregisterDate: null.NewTime(time.Now().AddDate(0, 0, 7), true),
				StartEvent:      null.NewTime(time.Now().AddDate(0, 0, 10), true),
				EndEvent:        null.NewTime(time.Now().AddDate(0, 0, 12), true),
				CreatedAt:       null.NewTime(time.Now(), true),
				UpdatedAt:       null.NewTime(time.Now(), true),
				Latitude:        null.NewString("40.7128", true),
				Longitude:       null.NewString("-74.0060", true),
				Title:           null.NewString("Sample Event", true),
				Description:     null.NewString("This is a sample event description.", true),
				AverageRating:   null.NewFloat64(4.5, true),
				IsFinished:      null.NewBool(false, true),
				IsOnline:        null.NewBool(true, true),
			},
			// Add more users as needed
		}
		for _, event := range EventsToInsert {
			err := event.Insert(context.Background(), DB, boil.Infer())
			if err != nil {
				fmt.Println("Error creating Event:", err)
				log.Fatal(err)
			}
		}

	} else {
		fmt.Println("Initial Events already exist, seeding process will not be executed")
		return nil
	}
	return nil
}

func SeedEventCreators() error {
	count, err := models.EventCreators().Count(context.Background(), DB)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		EventsCreatorToInsert := []models.EventCreator{
			{
				Name:           null.StringFrom("Kaiju BOP"),
				Longitude:      null.StringFrom("-73.985428"),
				Latitude:       null.StringFrom("40.748817"),
				Email:          null.StringFrom("john.doe@example.com"),
				Password:       null.StringFrom("securepassword1"),
				ProfilePicture: null.StringFrom("john_doe.jpg"),
				PhoneNumber:    null.StringFrom("+1234567890"),
			},
			{
				Name:           null.StringFrom("Dismi BOP"),
				Longitude:      null.StringFrom("-73.985428"),
				Latitude:       null.StringFrom("40.748817"),
				Email:          null.StringFrom("john.doe@example.com"),
				Password:       null.StringFrom("securepassword1"),
				ProfilePicture: null.StringFrom("john_doe.jpg"),
				PhoneNumber:    null.StringFrom("+1234567890"),
			},
		}
		for _, eventcreator := range EventsCreatorToInsert {
			err := eventcreator.Insert(context.Background(), DB, boil.Infer())
			if err != nil {
				fmt.Println("Error creating Event Creator :", err)
				log.Fatal(err)
			}
		}

	} else {
		fmt.Println("Initial Event Creator already exist, seeding process will not be executed")
		return nil
	}
	return nil
}
