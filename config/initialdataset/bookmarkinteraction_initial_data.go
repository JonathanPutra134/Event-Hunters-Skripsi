package dataset

import (
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
)

func InitializeUserBookmarkInteraction() []models.EventsBookmark {
	ViewsInteractionToInsert := []models.EventsBookmark{
		{
			UserID:  null.IntFrom(16),
			EventID: null.IntFrom(27),
		},
		{
			UserID:  null.IntFrom(16),
			EventID: null.IntFrom(28),
		},
		{
			UserID:  null.IntFrom(16),
			EventID: null.IntFrom(31),
		},
		// Add more users as needed
	}

	return ViewsInteractionToInsert
}
