package repository

import (
	"context"
	"event-hunters/config"
	"event-hunters/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetLatestEvent(limit int) ([]*models.Event, error) {
	events, err := models.Events(
		qm.OrderBy("created_at DESC"), // Order by created_at in descending order
		qm.Limit(limit),               // Limit the result to the specified number (5 in this case)
	).All(context.Background(), config.DB)

	if err != nil {
		return nil, err
	}

	return events, nil
}
