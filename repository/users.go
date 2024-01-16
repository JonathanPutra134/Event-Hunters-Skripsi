package repository

import (
	"context"
	"event-hunters/config"
	"event-hunters/dto"
	"event-hunters/helpers"
	"event-hunters/models"
	"fmt"
	"log"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func RegisterUser(request dto.UserRegistrationRequest, latitudeFloat float64, longitudeFloat float64) error {
	// Perform any additional validation on the form data as needed

	// Create a new user with the extracted data
	newUser := models.User{
		Name:        null.StringFrom(request.FullName),
		Email:       null.StringFrom(request.Email),
		Password:    helpers.HashPassword(request.Password),
		PhoneNumber: null.StringFrom(request.PhoneNumber),
		Address:     null.StringFrom(request.Address),
		Latitude:    null.StringFrom(request.Latitude),
		Longitude:   null.StringFrom(request.Longitude),
	}

	err := newUser.Insert(context.Background(), config.DB, boil.Infer())
	if err != nil {
		fmt.Println("Error registering user:", err)
		log.Fatal(err)
	}
	fmt.Printf("User %s inserted successfully.\n", newUser.Name.String)
	return nil
}
