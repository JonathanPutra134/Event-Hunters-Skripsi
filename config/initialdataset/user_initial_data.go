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
			Latitude:       null.StringFrom("-6.165"),
			Longitude:      null.StringFrom("106.6856"),
			Email:          null.StringFrom("john.doe@example.com"),
			Password:       helpers.HashPassword("securepassword1"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
		{
			Name:           null.StringFrom("Jonggun"),
			Latitude:       null.StringFrom("-6.165"),
			Longitude:      null.StringFrom("106.6856"),
			Email:          null.StringFrom("john.doe@example.com"),
			Password:       helpers.HashPassword("securepassword1"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
		{
			Name:           null.StringFrom("Jonathan"),
			Latitude:       null.StringFrom("-6.165"),
			Longitude:      null.StringFrom("106.6856"),
			Email:          null.StringFrom("jonathanputra134@gmail.com"),
			Password:       helpers.HashPassword("Berserker134"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
		{
			Name:           null.StringFrom("Mirsal"),
			Latitude:       null.StringFrom("-6.165"),
			Longitude:      null.StringFrom("106.6856"),
			Email:          null.StringFrom("john.doe@example.com"),
			Password:       helpers.HashPassword("securepassword1"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
		{
			Name:           null.StringFrom("Oliv"),
			Latitude:       null.StringFrom("-6.165"),
			Longitude:      null.StringFrom("106.6856"),
			Email:          null.StringFrom("john.doe@example.com"),
			Password:       helpers.HashPassword("securepassword1"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
		{
			Name:           null.StringFrom("Sanchez"),
			Latitude:       null.StringFrom("-6.165"),
			Longitude:      null.StringFrom("106.6856"),
			Email:          null.StringFrom("john.doe@example.com"),
			Password:       helpers.HashPassword("securepassword1"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
		{
			Name:           null.StringFrom("Yosua"),
			Latitude:       null.StringFrom("-6.165"),
			Longitude:      null.StringFrom("106.6856"),
			Email:          null.StringFrom("john.doe@example.com"),
			Password:       helpers.HashPassword("securepassword1"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
		{
			Name:           null.StringFrom("kimi"),
			Latitude:       null.StringFrom("-6.165"),
			Longitude:      null.StringFrom("106.6856"),
			Email:          null.StringFrom("kimi@gmail.com"),
			Password:       helpers.HashPassword("hehe"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
		// Add more users as needed
	}

	return UsersToInsert
}
