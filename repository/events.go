package repository

import (
	"context"
	"event-hunters/config"
	"event-hunters/dto"
	"event-hunters/models"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetLatestEvent(limit int) ([]*models.Event, error) {
	currentTime := time.Now()

	// Assuming registration_date is the name of your registration date column
	events, err := models.Events(
		qm.Where("preregister_date >= ?", currentTime), // Include only events with registration dates in the future
		qm.OrderBy("preregister_date ASC"),             // Order by registration_date in ascending order
		qm.Limit(limit),                                // Limit the result to the specified number
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

func GetSavedEvents(userID int) ([]*models.Event, error) {

	events, err := models.Events(
		qm.Select("events.*"),
		qm.InnerJoin("events_bookmarks ON events.id = events_bookmarks.event_id AND events_bookmarks.user_id = ?", userID),
	).All(context.Background(), config.DB)

	if err != nil {
		return nil, err
	}
	return events, nil
}

func GetRecommendedEvents(recommendations dto.RecommendationsResponse) (dto.RecommendedEventDetails, error) {
	// Extract event IDs from RecommendationsResponse
	var eventIDs []int
	var eventDistances []float64
	var eventDaysCounts []int
	var eventInteractionCounts []int
	var recommendedEventsWithDetails dto.RecommendedEventDetails

	for _, recommendation := range recommendations {
		eventIDs = append(eventIDs, recommendation.EventID)
		eventDistances = append(eventDistances, recommendation.DistanceKM)
		eventDaysCounts = append(eventDaysCounts, recommendation.DaysBeforeRegistration)
		eventInteractionCounts = append(eventInteractionCounts, recommendation.InteractionScore)
	}

	placeholders := make([]string, len(eventIDs))
	values := make([]interface{}, len(eventIDs))
	for i, id := range eventIDs {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		values[i] = id
	}

	// Use the placeholders in the IN clause
	query := fmt.Sprintf("SELECT * FROM events WHERE id IN (%s)", strings.Join(placeholders, ", "))

	// Fetch events using the query and values
	events, err := models.Events(qm.SQL(query, values...)).All(context.Background(), config.DB)

	if err != nil {
		return recommendedEventsWithDetails, err
	}

	// Reorder events based on the original order of eventIDs
	eventIDMap := make(map[int]*models.Event)
	for _, event := range events {
		eventIDMap[event.ID] = event
	}

	orderedEvents := make([]*models.Event, len(eventIDs))
	for i, id := range eventIDs {
		orderedEvents[i] = eventIDMap[id]
	}
	recommendedEventsWithDetails = dto.RecommendedEventDetails{
		Events:                 orderedEvents,
		Distance:               eventDistances,
		DaysBeforeRegistration: eventDaysCounts,
		InteractionScore:       eventInteractionCounts,
	}
	return recommendedEventsWithDetails, nil
}
