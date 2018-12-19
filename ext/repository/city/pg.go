package city

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

func (s *pgStore) List(ctx context.Context, paging paginate.Input, filter SearchFilter) ([]*models.City, *paginate.Paginate, error) {
	var q []qm.QueryMod

	if filter.Name != "" {
		q = append(q, qm.Where("city.name LIKE ?", "%"+filter.Name+"%"))
	}

	qWithPaginate := append(q, []qm.QueryMod{
		qm.Limit(paging.Limit),
		qm.Offset(paging.Offset),
	}...)

	cities, err := models.Cities(qWithPaginate...).All(ctx, s.db)
	if err != nil {
		return nil, nil, err
	}

	count, err := models.Cities(q...).Count(ctx, s.db)

	return cities, paginate.MakePaginate(int(count), paging.Offset, paging.Limit), nil
}

func (s *pgStore) Show(ctx context.Context, cityID string) (*models.City, error) {
	route, err := models.Cities(
		qm.Where("id = ?", cityID),
	).One(ctx, s.db)
	if err != nil {
		return nil, err
	}
	return route, nil
}
