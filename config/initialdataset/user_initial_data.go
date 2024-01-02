package dataset

import (
	"event-hunters/models"
	"time"

	"github.com/volatiletech/null/v8"
)

func InitializeUsers() []models.User {
	UsersToInsert := []models.User{
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

	return UsersToInsert
}
