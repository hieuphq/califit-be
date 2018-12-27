package repository

import (
	"github.com/hieuphq/califit-be/ext/repository/center"
	"github.com/hieuphq/califit-be/ext/repository/city"
	"github.com/hieuphq/califit-be/ext/repository/schedule"
)

//Repository --
type Repository struct {
	City     city.Repository
	Center   center.Repository
	Schedule schedule.Repository
}
