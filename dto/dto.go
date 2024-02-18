package dto

import (
	"time"
)

type UserRegistrationRequest struct {
	FullName    string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
	Latitude    string
	Longitude   string
}

type Session struct {
	ID         string
	Data       string
	Expiration time.Time
}

type ParsedEventSearchDate struct {
	MinRegDate        time.Time
	MaxRegDate        time.Time
	MinEventStartDate time.Time
	MaxEventStartDate time.Time
}
type EventWithTicketID struct {
	Title           string
	Image           string
	Location        string
	Startevent_date time.Time
	Endevent_date   time.Time
	TicketID        int
}
