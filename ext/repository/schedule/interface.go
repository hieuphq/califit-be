package schedule

import (
	"context"

	"github.com/hieuphq/califit-be/ext/utils/paginate"
	"github.com/hieuphq/califit-be/models"
)

// Repository for Schedule
type Repository interface {
	List(ctx context.Context, centerID string, paginate paginate.Input, filter SearchFilter) ([]*models.Schedule, *paginate.Paginate, error)
}
