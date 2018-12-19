package paginate

import (
	"math"

	"github.com/hieuphq/califit-be/goa/app"
)

// Input -- paginate form
type Input struct {
	Limit  int
	Offset int
}

// Paginate --
type Paginate struct {
	CurrentPage int
	TotalPage   int
	TotalItem   int
}

// MapToView --
func MapToView(p *Paginate) *app.Paginate {
	return &app.Paginate{
		CurrentPage: &p.CurrentPage,
		TotalPage:   &p.TotalPage,
		TotalItem:   &p.TotalItem,
	}
}

// MakePaginate --
func MakePaginate(total, offset, limit int) *Paginate {
	if limit == 0 {
		return &Paginate{
			CurrentPage: 0,
			TotalPage:   0,
			TotalItem:   0,
		}
	}

	totalPage := int(math.Ceil(float64(total) / float64(limit)))
	currentPage := int(math.Floor(float64(offset)/float64(limit))) + 1

	return &Paginate{
		CurrentPage: currentPage,
		TotalPage:   totalPage,
		TotalItem:   total,
	}
}
