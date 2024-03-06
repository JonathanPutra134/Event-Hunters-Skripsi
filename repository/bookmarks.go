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

func InsertEventBookmark(userID int, eventIDParams string) error {
	// Convert string parameters to integers

	eventID, err := strconv.Atoi(eventIDParams)
	if err != nil {
		return err // Handle the error, e.g., invalid input
	}
	exist := CheckBookmarkExist(userID, eventID)

	if exist {
		fmt.Println("USER ALREADY BOOKMARKED THIS EVENT")
		return errors.New("user already bookmarked this event")
	}

	bookmarkToInsert := models.EventsBookmark{
		UserID:       null.IntFrom(userID),
		EventID:      null.IntFrom(eventID),
		BookmarkDate: null.TimeFrom(time.Now()),
	}

	err = bookmarkToInsert.Insert(context.Background(), config.DB, boil.Infer())

	if err != nil {
		return err
	}
	fmt.Println("SUCCESS INSERT BOOKMARK")
	return nil
}

func CheckBookmarkExist(userID int, eventID int) bool {
	exist, _ := models.EventsBookmarks(
		qm.Where("user_id = ? AND event_id = ?", userID, eventID),
	).Exists(context.Background(), config.DB)

	return exist

}

// func CheckUserExistInBookmark(userID int) bool {
// 	exist, _ := models.EventsBookmarks(
// 		qm.Where("user_id = ?", userID),
// 	).Exists(context.Background(), config.DB)

// 	return exist

// }
