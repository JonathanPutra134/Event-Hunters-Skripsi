package repository

import (
	"context"
	"errors"
	"strconv"
	"time"

	"event-hunters/config"
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func InsertEventRating(ratingParams string, userID int, eventIDParams string) error {
	// Convert string parameters to integer

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
		return errors.New("user already rated this event")
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
	return nil
}

func CheckUserRatingExist(userID int, eventID int) bool {
	exist, _ := models.Ratings(
		qm.Where("user_id = ? AND event_id = ?", userID, eventID),
	).Exists(context.Background(), config.DB)

	return exist

}
