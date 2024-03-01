package dataset

import (
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
)

func InitializeTickets() []models.Ticket {
	TicketsToInsert := []models.Ticket{
		{
			UserID:  null.IntFrom(3),
			EventID: null.IntFrom(69),
		},
		{
			UserID:  null.IntFrom(3),
			EventID: null.IntFrom(70),
		},
		{
			UserID:  null.IntFrom(3),
			EventID: null.IntFrom(62),
		},
		{
			UserID:  null.IntFrom(3),
			EventID: null.IntFrom(51),
		},
		{
			UserID:  null.IntFrom(3),
			EventID: null.IntFrom(49),
		},
	}

	return TicketsToInsert
}
