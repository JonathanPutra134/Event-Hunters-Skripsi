package dataset

import (
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
)

func InitializeTickets() []models.Ticket {
	TicketsToInsert := []models.Ticket{
		{
			UserID:  null.IntFrom(3),
			EventID: null.IntFrom(1),
		},
		{
			UserID:  null.IntFrom(3),
			EventID: null.IntFrom(2),
		},
		{
			UserID:  null.IntFrom(3),
			EventID: null.IntFrom(3),
		},
	}

	return TicketsToInsert
}
