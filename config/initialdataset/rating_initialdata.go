package dataset

import (
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
)

func InitializeRatings() []models.Rating {
	RatingsToInsert := []models.Rating{
		{
			UserID:  null.IntFrom(2),
			EventID: null.IntFrom(68),
			Rating:  null.IntFrom(5),
		},
	}

	return RatingsToInsert
}
