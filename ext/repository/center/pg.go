package center

import (
	"context"
	"database/sql"

	"github.com/hieuphq/califit-be/ext/utils/paginate"
	"github.com/hieuphq/califit-be/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type pgStore struct {
	db *sql.DB
}

// NewPGStore ...
func NewPGStore(db *sql.DB) Repository {
	return &pgStore{
		db: db,
	}
}

func (s *pgStore) List(ctx context.Context, paging paginate.Input, filter SearchFilter) ([]*models.Center, *paginate.Paginate, error) {
	var q []qm.QueryMod

	if filter.Name != "" {
		q = append(q, qm.Where("center.name LIKE ?", "%"+filter.Name+"%"))
	}

	if filter.CityID != "" {
		q = append(q, qm.Where("center.city_id = ?", filter.CityID))
	}

	qWithPaginate := append(q, []qm.QueryMod{
		qm.Limit(paging.Limit),
		qm.Offset(paging.Offset),
	}...)

	cities, err := models.Centers(qWithPaginate...).All(ctx, s.db)
	if err != nil {
		return nil, nil, err
	}

	count, err := models.Centers(q...).Count(ctx, s.db)

	return cities, paginate.MakePaginate(int(count), paging.Offset, paging.Limit), nil
}

func (s *pgStore) Show(ctx context.Context, centerID string) (*models.Center, error) {
	route, err := models.Centers(
		qm.Where("id = ?", centerID),
	).One(ctx, s.db)
	if err != nil {
		return nil, err
	}

	r := route.R.NewStruct()

	// Add preload City
	if route.CityID.Valid {
		v, err := route.City().One(ctx, s.db)
		if err != nil {
			return nil, err
		}

		r.City = v
		route.R = r
	}

	// Add preload Address
	if route.AddressID.Valid {
		v, err := route.Address().One(ctx, s.db)
		if err != nil {
			return nil, err
		}

		r.Address = v
		route.R = r
	}

	return route, nil
}
