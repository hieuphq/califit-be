package schedule

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/hieuphq/califit-be/ext/utils/paginate"
	"github.com/hieuphq/califit-be/models"
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

func (s *pgStore) List(ctx context.Context, centerID string, paging paginate.Input, filter SearchFilter) ([]*models.Schedule, *paginate.Paginate, error) {
	var rs []*models.Schedule
	var q []qm.QueryMod
	q = append(q, qm.Where("schedule.center_id = ?", centerID))
	qWithPaginate := append(q, []qm.QueryMod{
		qm.Limit(paging.Limit),
		qm.Offset(paging.Offset),
	}...)

	schedules, err := models.Schedules(qWithPaginate...).All(ctx, s.db)
	if err != nil {
		return nil, nil, err
	}
	for _, schedule := range schedules {
		if schedule.CenterID.Valid {
			r := schedule.R.NewStruct()
			v, err := schedule.Center().One(ctx, s.db)
			if err != nil {
				return nil, nil, err
			}
			r.Center = v
			schedule.R = r
			rs = append(rs, schedule)
		}
	}
	count, err := models.Schedules(q...).Count(ctx, s.db)
	return rs, paginate.MakePaginate(int(count), paging.Offset, paging.Limit), nil
}
