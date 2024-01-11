package dataset

import (
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
)

func InitializeRatings() []models.Rating {
	RatingsToInsert := []models.Rating{
		{
			UserID:  null.IntFrom(2),
			EventID: null.IntFrom(4),
			Rating:  null.IntFrom(3),
		},
		{
			UserID:  null.IntFrom(2),
			EventID: null.IntFrom(6),
			Rating:  null.IntFrom(5),
		},
		{
			UserID:  null.IntFrom(2),
			EventID: null.IntFrom(12),
			Rating:  null.IntFrom(2),
		},
		{
			UserID:  null.IntFrom(2),
			EventID: null.IntFrom(17),
			Rating:  null.IntFrom(5),
		},
		{
			UserID:  null.IntFrom(2),
			EventID: null.IntFrom(1),
			Rating:  null.IntFrom(1),
		},
	}

	return RatingsToInsert
}
