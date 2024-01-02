package dataset

import (
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
)

func InitializeUserViewInteraction() []models.EventsView {
	ViewsInteractionToInsert := []models.EventsView{
		{
			UserID:  null.IntFrom(16),
			EventID: null.IntFrom(1),
		},
		{
			UserID:  null.IntFrom(16),
			EventID: null.IntFrom(2),
		},
		{
			UserID:  null.IntFrom(16),
			EventID: null.IntFrom(3),
		},
		// Add more users as needed
	}

	return ViewsInteractionToInsert
}
