package repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"event-hunters/config"
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func InsertEventRating(ratingParams string, userIDParams string, eventIDParams string) error {
	// Convert string parameters to integers
	userID, err := strconv.Atoi(userIDParams)
	if err != nil {

		return err // Handle the error, e.g., invalid input
	}

	eventID, err := strconv.Atoi(eventIDParams)
	if err != nil {
		return err // Handle the error, e.g., invalid input
	}

	rating, err := strconv.Atoi(ratingParams)
	if err != nil {
		return err // Handle the error, e.g., invalid input
	}
	exist := CheckUserRatingExist(userID, eventID)

	if exist {
		fmt.Println("USER ALREADY RATED THIS EVENT")
		return errors.New("User Already Rated This Event")
	}

	ratingToInsert := models.Rating{
		UserID:     null.IntFrom(userID),
		EventID:    null.IntFrom(eventID),
		Rating:     null.IntFrom(rating),
		RatingDate: null.TimeFrom(time.Now()),
	}

	err = ratingToInsert.Insert(context.Background(), config.DB, boil.Infer())

	if err != nil {
		return err
	}
	fmt.Println("SUCCESS INSERT RATING")
	return nil
}

func CheckUserRatingExist(userID int, eventID int) bool {
	exist, _ := models.Ratings(
		qm.Where("user_id = ? AND event_id = ?", userID, eventID),
	).Exists(context.Background(), config.DB)

	return exist

}
