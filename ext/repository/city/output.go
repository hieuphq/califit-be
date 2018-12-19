package city

import (
	"github.com/hieuphq/califit-be/goa/app"
	"github.com/hieuphq/califit-be/models"
)

// MapModelToViewGeneral ...
func MapModelToViewGeneral(r *models.City) (*app.City, error) {
	return &app.City{
		CityID:    &r.ID,
		CityName:  &r.Name.String,
		CreatedAt: &r.CreatedAt,
	}, nil
}
