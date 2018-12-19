package controller

import (
	"database/sql"

	"github.com/goadesign/goa"
	"github.com/hieuphq/califit-be/ext/repository"
	"github.com/hieuphq/califit-be/ext/repository/city"
	"github.com/hieuphq/califit-be/ext/utils/paginate"
	"github.com/hieuphq/califit-be/goa/app"
)

// CityController implements the city resource.
type CityController struct {
	*goa.Controller
	*repository.Repository
}

// NewCityController creates a city controller.
func NewCityController(service *goa.Service, repo *repository.Repository) *CityController {
	return &CityController{
		Controller: service.NewController("CityController"),
		Repository: repo,
	}
}

// List runs the list action.
func (c *CityController) List(ctx *app.ListCityContext) error {
	// CityController_List: start_implement

	rs, p, err := c.Repository.City.List(ctx, paginate.Input{
		Limit:  ctx.Limit,
		Offset: ctx.Offset,
	}, city.SearchFilter{
		Name: ctx.Name,
	})
	if err != nil {
		return ctx.InternalServerError(goa.ErrInternal(err))
	}

	data := []*app.City{}
	for _, v := range rs {
		tmp, err := city.MapModelToViewGeneral(v)
		if err != nil {
			continue
		}
		data = append(data, tmp)
	}

	res := &app.Cities{
		Data:     data,
		Paginate: paginate.MapToView(p),
	}

	return ctx.OK(res)
	// CityController_List: end_implement
}

// Show runs the show action.
func (c *CityController) Show(ctx *app.ShowCityContext) error {
	// CityController_Show: start_implement

	v, err := c.Repository.City.Show(ctx, ctx.CityID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.NotFound(goa.ErrNotFound(err))
		}
		return ctx.InternalServerError(goa.ErrInternal(err))
	}
	res, err := city.MapModelToViewGeneral(v)
	if err != nil {
		return ctx.InternalServerError(goa.ErrInternal(err))
	}

	return ctx.OK(res)
	// CityController_Show: end_implement
}
