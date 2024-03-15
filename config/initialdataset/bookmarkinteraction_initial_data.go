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
		{
			UserID:  null.IntFrom(2),
			EventID: null.IntFrom(68),
		},
		{
			UserID:  null.IntFrom(1),
			EventID: null.IntFrom(4),
		},
		{
			UserID:  null.IntFrom(2),
			EventID: null.IntFrom(4),
		},
		{
			UserID:  null.IntFrom(4),
			EventID: null.IntFrom(4),
		},
		{
			UserID:  null.IntFrom(1),
			EventID: null.IntFrom(16),
		},
		{
			UserID:  null.IntFrom(8),
			EventID: null.IntFrom(6),
		},
		{
			UserID:  null.IntFrom(8),
			EventID: null.IntFrom(18),
		},

		// Add more users as needed
	}

	return ViewsInteractionToInsert
}
