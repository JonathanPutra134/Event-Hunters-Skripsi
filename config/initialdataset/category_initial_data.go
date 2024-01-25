package dataset

import (
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
)

func InitializeCategory() []models.Category {
	CategoryToInsert := []models.Category{
		{
			Name: null.StringFrom("Education and Career"),
		},
		{
			Name: null.StringFrom("Entertainment & Performance"),
		},
		{
			Name: null.StringFrom("Charity"),
		},
		{
			Name: null.StringFrom("Competition"),
		},
		{
			Name: null.StringFrom("Sports"),
		},
		{
			Name: null.StringFrom("Art & Culture"),
		},
		{
			Name: null.StringFrom("Expo"),
		},
	}

	return CategoryToInsert
}
