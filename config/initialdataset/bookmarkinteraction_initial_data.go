package dataset

import (
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
)

func InitializeUserBookmarkInteraction() []models.EventsBookmark {
	ViewsInteractionToInsert := []models.EventsBookmark{
		{
			UserID:  null.IntFrom(3),
			EventID: null.IntFrom(1),
		},

		// Add more users as needed
	}

	return ViewsInteractionToInsert
}
