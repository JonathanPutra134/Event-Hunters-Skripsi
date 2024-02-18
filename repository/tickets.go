package repository

import (
	"context"
	"fmt"
	"strconv"

	"event-hunters/config"
	"event-hunters/dto"
	"event-hunters/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetTickets(userID int) ([]*dto.EventWithTicketID, error) {
	var eventsWithTicketID []*dto.EventWithTicketID

	err := models.Events(
		qm.Select("events.title, events.image, events.location, events.startevent_date, events.endevent_date, tickets.id AS ticket_id"),
		qm.InnerJoin("tickets ON events.id = tickets.event_id"),
		qm.Where("tickets.user_id = ?", userID),
		qm.OrderBy("events.created_at DESC"),
	).Bind(context.Background(), config.DB, &eventsWithTicketID)

	if err != nil {
		return nil, err
	}
	fmt.Println("EVENTS WITH TICKET ID")
	fmt.Println(eventsWithTicketID[0].Title)
	fmt.Println(eventsWithTicketID[0].Startevent_date)
	return eventsWithTicketID, nil
}

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

func ShowTicketInformation(idParams string) (*models.Event, error) {
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return nil, err
	}
	ticket, err := GetTicketById(id)
	if err != nil {
		return nil, err
	}

	fmt.Println("THIS IS THE TICKET EVENT ID")
	fmt.Println(ticket.EventID)
	event, err := models.Events(
		qm.Where("id = ?", ticket.EventID), // Filter by the specified ID
	).One(context.Background(), config.DB)

	if err != nil {
		return nil, err
	}

	return event, nil
}
