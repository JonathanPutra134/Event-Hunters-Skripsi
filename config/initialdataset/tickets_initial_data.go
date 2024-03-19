package dataset

import (
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
)

func InitializeTickets() []models.Ticket {
	TicketsToInsert := []models.Ticket{
		{
			UserID:  null.IntFrom(2),
			EventID: null.IntFrom(68),
		},
	}

	return TicketsToInsert
}
