package repository

import (
	"context"
	"event-hunters/config"
	"event-hunters/dto"
	"event-hunters/helpers"
	"event-hunters/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/volatiletech/sqlboiler/v4/boil"
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

func GetSearchedEvents(keyword string, categories []int, parsedSearchDate dto.ParsedEventSearchDate, eventType string) ([]*models.Event, error) {
	boil.DebugMode = true
	queryMods := []qm.QueryMod{}

	// Add conditions based on the provided parameters
	if keyword != "" {
		searchTerm := fmt.Sprintf("%%%s%%", keyword)
		queryMods = append(queryMods, qm.Expr(
			qm.Where(models.EventColumns.Title+" ILIKE ?", searchTerm),
			qm.Or2(qm.Where(models.EventColumns.Description+" ILIKE ?", searchTerm)),
			qm.Or2(qm.Where(models.EventColumns.Location+" ILIKE ?", searchTerm)),
		))
	}

	if len(categories) > 0 {
		categoryInterfaces := make([]interface{}, len(categories))
		for i, v := range categories {
			categoryInterfaces[i] = v
		}
		//THIS IS FOR MAKING THE PLACE HOLDER $1, $2, $3 IN CLAUSE
		placeholders := make([]string, len(categories))
		for i := range categories {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
		}
		fmt.Println(helpers.JoinInts(categories))

		inCondition := fmt.Sprintf("%s IN (%s)", models.EventsCategoryColumns.CategoryID, strings.Join(placeholders, ", "))
		queryMods = append(queryMods,
			qm.Load(models.EventRels.EventsCategories),
			qm.InnerJoin(models.TableNames.EventsCategories+" ON events_categories.event_id = events.id AND "+inCondition, categoryInterfaces...),
			qm.GroupBy("events.id"),
		)
	}

	// Add conditions for date range
	if !parsedSearchDate.MinRegDate.IsZero() {
		queryMods = append(queryMods, qm.Where("preregister_date >= ?", parsedSearchDate.MinRegDate))
	}

	if !parsedSearchDate.MaxRegDate.IsZero() {
		queryMods = append(queryMods, qm.Where("preregister_date <= ?", parsedSearchDate.MaxRegDate))
	}

	if !parsedSearchDate.MinEventStartDate.IsZero() {
		queryMods = append(queryMods, qm.Where("startevent_date >= ?", parsedSearchDate.MinEventStartDate))
	}

	if !parsedSearchDate.MaxEventStartDate.IsZero() {
		queryMods = append(queryMods, qm.Where("startevent_date <= ?", parsedSearchDate.MaxEventStartDate))
	}

	// Add conditions for online status
	switch eventType {
	case "online":
		queryMods = append(queryMods, qm.Where("is_online = ?", true))
	case "offline":
		queryMods = append(queryMods, qm.Where("is_online = ?", false))
	// No filter for both online and offline
	default:
		// Do nothing
	}

	events, err := models.Events(queryMods...).All(context.Background(), config.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//FOR DEBUH
	// for i, event := range events {
	// 	fmt.Println("Event", i)
	// 	fmt.Println(event.Title)
	// 	fmt.Println(event.PreregisterDate)
	// 	fmt.Println(event.StarteventDate)
	// }

	return events, nil
}
