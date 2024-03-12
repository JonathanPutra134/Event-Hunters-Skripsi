package repository

import (
	"context"
	"event-hunters/config"
	"event-hunters/dto"
	"event-hunters/helpers"
	"event-hunters/models"
	"fmt"
	"log"
	"strings"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func RegisterUser(request dto.UserRegistrationRequest, latitudeFloat float64, longitudeFloat float64) error {
	// Perform any additional validation on the form data as needed

	// Create a new user with the extracted data
	newUser := models.User{
		Name:        null.StringFrom(request.FullName),
		Email:       null.StringFrom(strings.ToLower(request.Email)),
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

func GetUserById(id int) (*models.User, error) {
	// You may need to replace "YourPrimaryKeyColumn" with the actual name of your primary key column.
	user, err := models.Users(qm.Where("id = ?", id)).One(context.Background(), config.DB)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func CheckEmailExists(email string) bool {
	exists, _ := models.Users(qm.Where("email = ?", email)).Exists(context.Background(), config.DB)
	return exists
}
func GetUserByEmail(email string) (*models.User, error) {
	// You may need to replace "YourPrimaryKeyColumn" with the actual name of your primary key column.
	user, err := models.Users(qm.Where("email = ?", email)).One(context.Background(), config.DB)
	if err != nil {
		return nil, err
	}

	return user, nil
}
