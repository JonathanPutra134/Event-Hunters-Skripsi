package dataset

import (
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
)

func InitializeUserViewInteraction() []models.EventsView {
	ViewsInteractionToInsert := []models.EventsView{
		{
			UserID:  null.IntFrom(3),
			EventID: null.IntFrom(3),
		},
		{
			UserID:  null.IntFrom(2),
			EventID: null.IntFrom(4),
		},
		{
			UserID:  null.IntFrom(8),
			EventID: null.IntFrom(6),
		},
		{
			UserID:  null.IntFrom(8),
			EventID: null.IntFrom(18),
		},
		{
			UserID:  null.IntFrom(8),
			EventID: null.IntFrom(57),
		},

		// Add more users as needed
	}

	return ViewsInteractionToInsert
}
