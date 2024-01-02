package config

import (
	"context"

	"fmt"
	"log"

	dataset "event-hunters/config/initialdataset"
	"event-hunters/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func SeedDataUsers() error {
	count, err := models.Users().Count(context.Background(), DB)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		usersToInsert := dataset.InitializeUsers()
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
		categoriesToInsert := dataset.InitializeCategory()
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
		EventsToInsert := dataset.InitializeEvents()
		for _, event := range EventsToInsert {
			err := event.Insert(context.Background(), DB, boil.Infer())
			if err != nil {
				fmt.Println("Error creating Event:", err)
				log.Fatal(err)
			}
		}
		fmt.Println("Successful inserting events")
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
		EventsCreatorToInsert := dataset.InitializeEventCreator()
		for _, eventcreator := range EventsCreatorToInsert {
			err := eventcreator.Insert(context.Background(), DB, boil.Infer())
			if err != nil {
				fmt.Println("Error creating Event Creator :", err)
				log.Fatal(err)
			}
		}
		fmt.Println("Successful inserting event creators")
	} else {
		fmt.Println("Initial Event Creator already exist, seeding process will not be executed")
		return nil
	}
	return nil
}

func SeedEventViewsInteraction() error {
	count, err := models.EventsViews().Count(context.Background(), DB)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		EventsViewToInsert := dataset.InitializeUserViewInteraction()
		for _, eventviews := range EventsViewToInsert {
			err := eventviews.Insert(context.Background(), DB, boil.Infer())
			if err != nil {
				fmt.Println("Error creating Event Views :", err)
				log.Fatal(err)
			}
		}
		fmt.Println("Successful inserting event views")

	} else {
		fmt.Println("Initial Event Views already exist, seeding process will not be executed")
		return nil
	}
	return nil
}

func SeedEventBookmarksInteraction() error {
	count, err := models.EventsBookmarks().Count(context.Background(), DB)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		EventsBookmarkToInsert := dataset.InitializeUserBookmarkInteraction()
		for _, eventbookmarks := range EventsBookmarkToInsert {
			err := eventbookmarks.Insert(context.Background(), DB, boil.Infer())
			if err != nil {
				fmt.Println("Error creating Event Bookmarks :", err)
				log.Fatal(err)
			}
		}
		fmt.Println("Successful inserting event Bookmarks")

	} else {
		fmt.Println("Initial Event Bookmarks already exist, seeding process will not be executed")
		return nil
	}
	return nil
}
