package controller

import (
	"github.com/goadesign/goa"

	"github.com/hieuphq/califit-be/ext/repository"
	"github.com/hieuphq/califit-be/goa/app"
	"github.com/hieuphq/califit-be/ext/repository/schedule"
	"github.com/hieuphq/califit-be/ext/utils/paginate"

)

// ScheduleController implements the schedule resource.
type ScheduleController struct {
	*goa.Controller
	*repository.Repository
}

// NewScheduleController creates a schedule controller.
func NewScheduleController(service *goa.Service,repo *repository.Repository) *ScheduleController {
	return &ScheduleController{
		Controller: service.NewController("ScheduleController"),
		Repository: repo,
	}
}

// List runs the show action.
func (c *ScheduleController) List(ctx *app.ListScheduleContext) error {
	// ScheduleController_Show: start_implement

	rs, p,err := c.Repository.Schedule.List(
		ctx,
		ctx.CenterID,
		paginate.Input{
			Limit:  ctx.Limit,
			Offset: ctx.Offset,
		},
		schedule.SearchFilter{
			DateFrom: "",
			DateTo: "",
		})
		if err != nil {
			return ctx.InternalServerError(goa.ErrInternal(err))
		}

		data := []*app.Schedule{}
		for _, v := range rs {
			tmp, err := schedule.MapModelToViewGeneral(v)
			if err != nil {
				continue
			}
			data = append(data, tmp)
		}

		res := &app.Schedules{
			Data:     data,
			Paginate: paginate.MapToView(p),
		}

		return ctx.OK(res)
	// ScheduleController_Show: end_implement
}
