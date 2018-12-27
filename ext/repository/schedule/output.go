package schedule

import (
	"github.com/hieuphq/califit-be/goa/app"
	"github.com/hieuphq/califit-be/models"
)

// MapModelToViewGeneral ...
func MapModelToViewGeneral(r *models.Schedule) (*app.Schedule, error) {
	centerInfo := &app.Center{}
	if r.CenterID.Valid && r.R.Center != nil {
		centerInfo = &app.Center{
			ID:        &r.R.Center.ID,
			Name:      &r.R.Center.Name.String,
			CreatedAt: &r.R.Center.CreatedAt,
		}
	}
	return &app.Schedule{
		ID:       &r.ID,
		CenterID: &r.CenterID.String,
		Center:   centerInfo,
		StartAt:  &r.StartAt.Time,
		EndAt:    &r.EndAt.Time,
	}, nil
}
