package repository

import (
	"github.com/hieuphq/califit-be/ext/repository/center"
	"github.com/hieuphq/califit-be/ext/repository/city"
)

//Repository --
type Repository struct {
	City   city.Repository
	Center center.Repository
}
