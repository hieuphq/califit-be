package center

import (
	"context"
	"errors"

	"github.com/hieuphq/califit-be/ext/utils/paginate"
	"github.com/hieuphq/califit-be/models"
)

// Repository for Center
type Repository interface {
	List(ctx context.Context, paginate paginate.Input, filter SearchFilter) ([]*models.Center, *paginate.Paginate, error)
	Show(ctx context.Context, centerID string) (*models.Center, error)
}

// Var list
var (
	ErrCenterNameExist = errors.New("Center name already exist in company")
)
