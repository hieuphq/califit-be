package city

import (
	"context"
	"errors"

	"github.com/hieuphq/califit-be/ext/utils/paginate"
	"github.com/hieuphq/califit-be/models"
)

// Repository for city
type Repository interface {
	List(ctx context.Context, paginate paginate.Input, filter SearchFilter) ([]*models.City, *paginate.Paginate, error)
	Show(ctx context.Context, cityID string) (*models.City, error)
}

// Var list
var (
	ErrCityNameExist = errors.New("city name already exist in company")
)
