package dataset

import (
	"event-hunters/helpers"
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
)

func InitializeUsers() []models.User {
	UsersToInsert := []models.User{
		{
			Name:           null.StringFrom("John Doe"),
			Longitude:      null.StringFrom("-6.165"),
			Latitude:       null.StringFrom("106.6856"),
			Email:          null.StringFrom("john.doe@example.com"),
			Password:       helpers.HashPassword("securepassword1"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
		{
			Name:           null.StringFrom("Jonggun"),
			Longitude:      null.StringFrom("-6.165"),
			Latitude:       null.StringFrom("106.6856"),
			Email:          null.StringFrom("john.doe@example.com"),
			Password:       helpers.HashPassword("securepassword1"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
		{
			Name:           null.StringFrom("Jonathan"),
			Longitude:      null.StringFrom("-6.165"),
			Latitude:       null.StringFrom("106.6856"),
			Email:          null.StringFrom("jonathanputra134@gmail.com"),
			Password:       helpers.HashPassword("Berserker134"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
		// Add more users as needed
	}

	return UsersToInsert
}
