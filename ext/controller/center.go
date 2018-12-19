package controller

import (
	"database/sql"

	"github.com/goadesign/goa"
	"github.com/hieuphq/califit-be/ext/repository"
	"github.com/hieuphq/califit-be/ext/repository/center"
	"github.com/hieuphq/califit-be/ext/utils/paginate"
	"github.com/hieuphq/califit-be/goa/app"
)

// CenterController implements the center resource.
type CenterController struct {
	*goa.Controller
	*repository.Repository
}

// NewCenterController creates a center controller.
func NewCenterController(service *goa.Service, repo *repository.Repository) *CenterController {
	return &CenterController{
		Controller: service.NewController("CenterController"),
		Repository: repo,
	}
}

// List runs the list action.
func (c *CenterController) List(ctx *app.ListCenterContext) error {
	// CenterController_List: start_implement

	rs, p, err := c.Repository.Center.List(ctx, paginate.Input{
		Limit:  ctx.Limit,
		Offset: ctx.Offset,
	}, center.SearchFilter{
		Name:   ctx.Name,
		CityID: ctx.CityID,
	})
	if err != nil {
		return ctx.InternalServerError(goa.ErrInternal(err))
	}

	data := []*app.Center{}
	for _, v := range rs {
		tmp, err := center.MapModelToViewGeneral(v)
		if err != nil {
			continue
		}
		data = append(data, tmp)
	}

	res := &app.Centers{
		Data:     data,
		Paginate: paginate.MapToView(p),
	}

	return ctx.OK(res)
	// CenterController_List: end_implement
}

// Show runs the show action.
func (c *CenterController) Show(ctx *app.ShowCenterContext) error {
	// CenterController_Show: start_implement

	v, err := c.Repository.Center.Show(ctx, ctx.CenterID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.NotFound(goa.ErrNotFound(err))
		}
		return ctx.InternalServerError(goa.ErrInternal(err))
	}
	res, err := center.MapModelToViewGeneral(v)
	if err != nil {
		return ctx.InternalServerError(goa.ErrInternal(err))
	}

	return ctx.OK(res)
	// CenterController_Show: end_implement
}
