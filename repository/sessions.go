package repository

import (
	"context"
	"event-hunters/config"
	"event-hunters/models"
	"fmt"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func StoreSession(user *models.User) (string, error) {
	newSession := models.Session{
		UserID: null.IntFrom(user.ID),
	}

	err := newSession.Insert(context.Background(), config.DB, boil.Infer())
	if err != nil {
		return "", err
	}

	fmt.Printf("Session for user stored successfully. Session ID: %s\n", newSession.ID)
	return newSession.ID, nil
}

// CheckSessionExists checks if a session exists for the given user ID in the database.
func CheckSessionExists(userID int) (bool, error) {
	exists, err := models.Sessions(models.SessionWhere.UserID.EQ(null.IntFrom(userID))).Exists(context.Background(), config.DB)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func GetUserBySessionID(sessionID string) (*models.User, error) {
	session, err := models.Sessions(models.SessionWhere.ID.EQ(sessionID)).One(context.Background(), config.DB)
	if err != nil {
		return nil, err
	}
	user, err := GetUserById(session.UserID.Int)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUserSession(sessionID string) error {
	_, err := models.Sessions(models.SessionWhere.ID.EQ(sessionID)).DeleteAll(context.Background(), config.DB)
	if err != nil {
		return err
	}
	return nil
}
