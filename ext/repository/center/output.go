package center

import (
	"github.com/hieuphq/califit-be/goa/app"
	"github.com/hieuphq/califit-be/models"
)

// MapModelToViewGeneral ...
func MapModelToViewGeneral(r *models.Center) (*app.Center, error) {
	addressInfo := &app.Address{}

	// Not show vehicle info when route have vehicle ID (this trick for mobile not need to store vehicleID)
	if r.AddressID.Valid && r.R.Address != nil {
		addressInfo = &app.Address{
			ID:      &r.R.Address.ID,
			Name:    &r.R.Address.Name.String,
			Street1: &r.R.Address.Street1.String,
			Street2: &r.R.Address.Street2.String,
			State:   &r.R.Address.State.String,
			City:    &r.R.Address.City.String,
			Country: &r.R.Address.Country.String,
			ZipCode: &r.R.Address.ZipCode.String,
			Lat:     &r.R.Address.Lat.Float64,
			Lng:     &r.R.Address.LNG.Float64,
		}
	}

	cityInfo := &app.City{}

	// Not show vehicle info when route have vehicle ID (this trick for mobile not need to store vehicleID)
	if r.CityID.Valid && r.R.City != nil {
		cityInfo = &app.City{
			ID:        &r.R.City.ID,
			Name:      &r.R.City.Name.String,
			CreatedAt: &r.R.City.CreatedAt,
		}
	}
	return &app.Center{
		ID:        &r.ID,
		Name:      &r.Name.String,
		Address:   addressInfo,
		AddressID: &r.AddressID.String,
		City:      cityInfo,
		CityID:    &r.CityID.String,
		CreatedAt: &r.CreatedAt,
	}, nil
}
