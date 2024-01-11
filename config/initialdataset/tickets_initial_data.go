package dataset

import (
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
)

func InitializeTickets() []models.Ticket {
	TicketsToInsert := []models.Ticket{
		{
			UserID:  null.IntFrom(2),
			EventID: null.IntFrom(47),
		},
		{
			UserID:  null.IntFrom(2),
			EventID: null.IntFrom(48),
		},
		{
			UserID:  null.IntFrom(2),
			EventID: null.IntFrom(39),
		},
	}

	return TicketsToInsert
}
