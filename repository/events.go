package repository

import (
	"context"
	"event-hunters/config"
	"event-hunters/models"
	"strconv"

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

func GetFilteredEventsByCategory(filterbycategory string) ([]*models.Event, error) {
	events, err := models.Events(
		qm.OrderBy("created_at DESC"), // Order by created_at in descending order
	).All(context.Background(), config.DB)

	if err != nil {
		return nil, err
	}

	// Filter events that have "Entertainment & Performance" category
	filteredEvents := make([]*models.Event, 0)
	for _, event := range events {
		for _, category := range event.Category {
			if category == filterbycategory {
				filteredEvents = append(filteredEvents, event)
				break // Break out of the inner loop once the category is found
			}
		}
	}
	return filteredEvents, nil
}

func GetEventById(idParams string) (*models.Event, error) {
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return nil, err
	}
	event, err := models.Events(
		qm.Where("id = ?", id), // Filter by the specified ID
	).One(context.Background(), config.DB)

	if err != nil {
		return nil, err
	}

	return event, nil
}
