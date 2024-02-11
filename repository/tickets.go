package repository

import (
	"context"
	"event-hunters/config"
	"event-hunters/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetTickets(userID int) ([]*models.Event, error) {
	events, err := models.Events(
		qm.InnerJoin("tickets ON events.id = tickets.event_id"),
		qm.Where("tickets.user_id = ?", userID),
		qm.OrderBy("events.created_at DESC"), // Order by created_at in descending order
	).All(context.Background(), config.DB)

	if err != nil {
		return nil, err
	}

	return events, nil
}
