package dataset

import (
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
)

func InitializeEventCreator() []models.EventCreator {
	EventCreatorToInsert := []models.EventCreator{
		{
			Name:           null.StringFrom("Kaiju BOP"),
			Email:          null.StringFrom("john.doe@example.com"),
			Password:       null.StringFrom("securepassword1"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
		{
			Name:           null.StringFrom("Dismi BOP"),
			Email:          null.StringFrom("john.doe@example.com"),
			Password:       null.StringFrom("securepassword1"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
		{
			Name:           null.StringFrom("Event Kampus"),
			Email:          null.StringFrom("john.doe@example.com"),
			Password:       null.StringFrom("securepassword1"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
	}

	return EventCreatorToInsert
}
