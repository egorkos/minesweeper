package repository

import "github.com/egorkos/minesweeper/app/domain/model"

type GameRepository interface {
	FindAll() ([]*model.Game, error)
	FindById(id int) (*model.Game, error)
	Upsert(*model.Game) error
}
