package dto

import (
	"event-hunters/models"
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

type Recommendation struct {
	EventID                int     `json:"event_id"`
	FinalScore             float64 `json:"final_score"`
	DaysBeforeRegistration int     `json:"days_before_registration"`
	DistanceKM             float64 `json:"distance(km)"`
	InteractionScore       int     `json:"interaction_score"`
	SimilarityScore        float64 `json:"similarity_score"`
}

type RecommendedEventDetails struct {
	Events                 []*models.Event
	Distance               []float64
	DaysBeforeRegistration []int
	InteractionScore       []int
	// Other fields you need
}

// Define a struct to represent the overall JSON response
type RecommendationsResponse []Recommendation
