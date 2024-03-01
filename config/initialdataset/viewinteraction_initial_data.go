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
		// Add more users as needed
	}

	return ViewsInteractionToInsert
}
