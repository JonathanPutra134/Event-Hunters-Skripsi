package repository

import (
	"context"
	"fmt"
	"strconv"

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

//func GetTickets(userID int) ([]*dto.TicketWithEvent, error) {
//	query := `
//        SELECT
//            tickets.*,
//            events.title AS event_title,
//            events.id AS event_id,
//            events.location AS event_location,
//            events.start_date AS event_start_date
//        FROM tickets
//        INNER JOIN events ON events.id = tickets.event_id
//        WHERE tickets.user_id = $1
//        ORDER BY events.created_at DESC
//    `
//
//	rows, err := config.DB.Query(query, userID)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var ticketsWithEvents []*dto.TicketWithEvent
//
//	for rows.Next() {
//		var ticketWithEvent dto.TicketWithEvent
//		// Scan the results into the custom structure
//		if err := rows.Scan(
//			&ticketWithEvent.TicketId,
//			&ticketWithEvent.UserId,
//			&ticketWithEvent.EventId,
//			&ticketWithEvent.EventTitle,
//			&ticketWithEvent.EventId,
//			&ticketWithEvent.EventLocation,
//			&ticketWithEvent.EventStartDate,
//		); err != nil {
//			return nil, err
//		}
//
//		ticketsWithEvents = append(ticketsWithEvents, &ticketWithEvent)
//	}
//	fmt.Println("TICKETS WITH EVENTS")
//	fmt.Println(ticketsWithEvents)
//	if err := rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return ticketsWithEvents, nil
//}

func GetAllTicketsByUserId(userID int) ([]*models.Ticket, error) {
	tickets, err := models.Tickets(
		qm.Where("user_id = ?", userID), // Filter by the specified ID
	).All(context.Background(), config.DB)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func GetTicketById(id int) (*models.Ticket, error) {
	ticket, err := models.Tickets(
		qm.Where("id = ?", id), // Filter by the specified ID
	).One(context.Background(), config.DB)

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func ShowTicketInformation(idParams string) (*models.Event, *models.User, error) {
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return nil, nil, err
	}
	ticket, err := GetTicketById(id)
	fmt.Println("THIS IS THE TICKET EVENT ID")
	fmt.Println(ticket.EventID)
	event, err := models.Events(
		qm.Where("id = ?", ticket.EventID), // Filter by the specified ID
	).One(context.Background(), config.DB)

	if err != nil {
		return nil, nil, err
	}
	user, err := models.Users(
		qm.Where("id = ?", ticket.UserID), // Filter by the specified ID
	).One(context.Background(), config.DB)

	if err != nil {
		return nil, nil, err
	}

	return event, user, nil
}
