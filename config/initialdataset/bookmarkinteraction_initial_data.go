package dataset

import (
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
)

func InitializeUserBookmarkInteraction() []models.EventsBookmark {
	ViewsInteractionToInsert := []models.EventsBookmark{
		{
			UserID:  null.IntFrom(2),
			EventID: null.IntFrom(68),
		},
	}

	return ViewsInteractionToInsert
}
