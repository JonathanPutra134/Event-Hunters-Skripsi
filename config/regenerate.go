package config

import (
	"fmt"
)

func RegenerateDataEvents() error {
	err := DeleteEvents()
	if err != nil {
		fmt.Println("error deleting Events")
	}
	err = SeedEvents()
	if err != nil {
		fmt.Println("error seeding events")
	}
	return nil
}
func RegenerateDataCategory() error {
	err := DeleteCategories()
	if err != nil {
		fmt.Println("error deleting categories")
	}
	err = SeedDataCategory()
	if err != nil {
		fmt.Println("error seeding categories")
	}
	return nil
}

func RegenerateDataUsers() error {
	err := DeleteUsers()
	if err != nil {
		fmt.Println("error deleting users")
	}
	err = SeedDataUsers()
	if err != nil {
		fmt.Println("error seeding users")
	}
	return nil
}
func RegenerateDataEventCreators() error {
	err := DeleteEventCreators()
	if err != nil {
		fmt.Println("error deleting Event Creators")
	}
	err = SeedDataUsers()
	if err != nil {
		fmt.Println("error seeding Event Creators")
	}
	return nil
}

func RegenerateEventViewsInteraction() error {
	err := DeleteEventViewsInteraction()
	if err != nil {
		fmt.Println("error deleting Event Views interaction")
	}
	err = SeedEventViewsInteraction()
	if err != nil {
		fmt.Println("error seeding Event Views interaction")
	}
	return nil
}

func RegenerateAllScenarioForRecommendation() error {
	err := DeleteEventViewsInteraction()
	if err != nil {
		fmt.Println("error deleting Event Views interaction")
	}
	err = DeleteEventBookmarksInteraction()
	if err != nil {
		fmt.Println("error deleting Event Bookmarks interaction")
	}
	err = DeleteEventsCategoriesRelation()
	if err != nil {
		fmt.Println("error deleting EventCategories Relation")
	}

	err = DeleteCategories()
	if err != nil {
		fmt.Println("error deleting categories")
	}

	err = DeleteTickets()
	if err != nil {
		fmt.Println("error deleting Tickets")
	}
	err = DeleteRatings()
	if err != nil {
		fmt.Println("error deleting Ratings")
	}
	err = DeleteUsers()
	if err != nil {
		fmt.Println("error deleting Users")
	}

	err = DeleteEvents()
	if err != nil {
		fmt.Println("error deleting Events")
	}

	err = DeleteEventCreators()
	if err != nil {
		fmt.Println("error deleting the event creators")
	}

	err = ResetSerial("events")
	if err != nil {
		fmt.Println("error resetting serial id in events")
	}
	err = ResetSerial("event_creators")
	if err != nil {
		fmt.Println("error resetting serial id in events creators")
	}
	err = ResetSerial("users")
	if err != nil {
		fmt.Println("error resetting serial id in users")
	}
	err = ResetSerial("events_categories")
	if err != nil {
		fmt.Println("error resetting serial id in events_categories")
	}
	err = ResetSerial("category")
	if err != nil {
		fmt.Println("error resetting serial id in events_categories")
	}
	err = ResetSerial("tickets")
	if err != nil {
		fmt.Println("error resetting serial id in tickets")
	}
	err = SeedEventCreators()
	if err != nil {
		fmt.Println("error seeding Event Creators")
	}
	err = SeedDataUsers()
	if err != nil {
		fmt.Println("error seeding Users")
	}
	err = SeedDataCategory()
	if err != nil {
		fmt.Println("error seeding Event Creators")
	}

	err = SeedEvents()
	if err != nil {
		fmt.Println("error seeding Event Views interaction")
	}

	err = SeedEventViewsInteraction()
	if err != nil {
		fmt.Println("error seeding Event Views interaction")
	}
	err = SeedEventBookmarksInteraction()
	if err != nil {
		fmt.Println("error seeding Event Bookmarks interaction")
	}
	err = SeedTickets()
	if err != nil {
		fmt.Println("error seeding Tickets")
	}
	err = SeedRatings()
	if err != nil {
		fmt.Println("error seeding Ratings")
	}
	return nil
}
