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
			UserID:  null.IntFrom(6),
			EventID: null.IntFrom(16),
		},
		{
			UserID:  null.IntFrom(7),
			EventID: null.IntFrom(64),
		},
		{
			UserID:  null.IntFrom(7),
			EventID: null.IntFrom(68),
		},
		{
			UserID:  null.IntFrom(7),
			EventID: null.IntFrom(56),
		},

		// Add more users as needed
	}

	return ViewsInteractionToInsert
}
