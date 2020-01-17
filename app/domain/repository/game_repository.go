package repository

import (
	"github.com/egorkos/minesweeper/app/domain/model"
	"github.com/egorkos/minesweeper/app/interface/apierr"
)

type GameRepository interface {
	FindAll() ([]*model.Game, *apierr.ApiError)
	FindByID(ID int) (*model.Game, *apierr.ApiError)
	Upsert(*model.Game) *apierr.ApiError
}
