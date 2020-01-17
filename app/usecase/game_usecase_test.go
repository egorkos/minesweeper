package usecase

import (
	"testing"

	"github.com/egorkos/minesweeper/app/domain/model"
	"github.com/egorkos/minesweeper/app/domain/repository"
	"github.com/egorkos/minesweeper/app/domain/service"
	"github.com/egorkos/minesweeper/app/interface/apierr"
	"github.com/stretchr/testify/assert"
)

type mockGameRepository struct {
	mockFindAll  func() ([]*model.Game, *apierr.ApiError)
	mockFindByID func(id int) (*model.Game, *apierr.ApiError)
	mockUpsert   func(*model.Game) *apierr.ApiError
}

func (m mockGameRepository) FindAll() ([]*model.Game, *apierr.ApiError) {
	return m.mockFindAll()
}

func (m mockGameRepository) FindByID(id int) (*model.Game, *apierr.ApiError) {
	return m.mockFindByID(id)
}

func (m mockGameRepository) Upsert(game *model.Game) *apierr.ApiError {
	return m.mockUpsert(game)
}

func TestGameUsecaseReveal(t *testing.T) {
	minedCell := model.Cell{
		Mine:        true,
		Revealed:    false,
		Flagged:     false,
		MinesAround: 0,
	}
	emptyCell := model.Cell{
		Mine:        false,
		Revealed:    false,
		Flagged:     false,
		MinesAround: 0,
	}
	revealedCell := model.Cell{
		Mine:        false,
		Revealed:    true,
		Flagged:     false,
		MinesAround: 0,
	}
	flaggedCell := model.Cell{
		Mine:        false,
		Revealed:    false,
		Flagged:     true,
		MinesAround: 0,
	}

	cases := []struct {
		name       string
		ID         int
		repository repository.GameRepository
		row        int
		col        int
		errText    string
	}{
		{
			name: "FAIL/FINISHED_GAME",
			ID:   1,
			repository: &mockGameRepository{
				mockFindByID: func(ID int) (*model.Game, *apierr.ApiError) {
					game := model.Game{
						Rows:  1,
						Cols:  3,
						Mines: 1,
						Grid: [][]model.Cell{
							{emptyCell, emptyCell, minedCell},
						},
						Status: model.Win,
					}
					return &game, nil
				},
			},
			row:     1,
			col:     1,
			errText: "Can't update cells on a finished game",
		},
		{
			name: "FAIL/EXCEED_ROW",
			ID:   1,
			repository: &mockGameRepository{
				mockFindByID: func(ID int) (*model.Game, *apierr.ApiError) {
					game := model.Game{
						Rows:  1,
						Cols:  3,
						Mines: 1,
						Grid: [][]model.Cell{
							{emptyCell, emptyCell, minedCell},
						},
						Status: model.Undefined,
					}
					return &game, nil
				},
			},
			row:     10,
			col:     1,
			errText: "Row value exceeded grid limits",
		},
		{
			name: "FAIL/EXCEED_COL",
			ID:   1,
			repository: &mockGameRepository{
				mockFindByID: func(ID int) (*model.Game, *apierr.ApiError) {
					game := model.Game{
						Rows:  1,
						Cols:  3,
						Mines: 1,
						Grid: [][]model.Cell{
							{emptyCell, emptyCell, minedCell},
						},
						Status: model.Undefined,
					}
					return &game, nil
				},
			},
			row:     0,
			col:     10,
			errText: "Col value exceeded grid limits",
		},
		{
			name: "FAIL/ALREADY_REVEALED_CELL",
			ID:   1,
			repository: &mockGameRepository{
				mockFindByID: func(ID int) (*model.Game, *apierr.ApiError) {
					game := model.Game{
						Rows:  1,
						Cols:  3,
						Mines: 1,
						Grid: [][]model.Cell{
							{revealedCell, emptyCell, minedCell},
						},
						Status: model.Undefined,
					}
					return &game, nil
				},
			},
			row:     0,
			col:     0,
			errText: "Can't update an already revealed cell",
		},
		{
			name: "FAIL/FLAGGED_CELL",
			ID:   1,
			repository: &mockGameRepository{
				mockFindByID: func(ID int) (*model.Game, *apierr.ApiError) {
					game := model.Game{
						Rows:  1,
						Cols:  3,
						Mines: 1,
						Grid: [][]model.Cell{
							{flaggedCell, emptyCell, minedCell},
						},
						Status: model.Undefined,
					}
					return &game, nil
				},
			},
			row:     0,
			col:     0,
			errText: "Can't reveal a flagged cell",
		},
		{
			name: "OK/MINED_CELL",
			ID:   1,
			repository: &mockGameRepository{
				mockUpsert: func(game *model.Game) *apierr.ApiError {
					return nil
				},
				mockFindByID: func(ID int) (*model.Game, *apierr.ApiError) {
					game := model.Game{
						Rows:  1,
						Cols:  3,
						Mines: 1,
						Grid: [][]model.Cell{
							{flaggedCell, emptyCell, minedCell},
						},
						Status: model.Undefined,
					}
					return &game, nil
				},
			},
			row:     0,
			col:     2,
			errText: "",
		},
		{
			name: "OK/REVEAL",
			ID:   1,
			repository: &mockGameRepository{
				mockUpsert: func(game *model.Game) *apierr.ApiError {
					return nil
				},
				mockFindByID: func(ID int) (*model.Game, *apierr.ApiError) {
					game := model.Game{
						Rows:  1,
						Cols:  3,
						Mines: 1,
						Grid: [][]model.Cell{
							{emptyCell, emptyCell, minedCell},
						},
						Status: model.Undefined,
					}
					return &game, nil
				},
			},
			row:     0,
			col:     1,
			errText: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := c.repository
			service := service.NewGameService(repo)

			gameUsecase := gameUsecase{
				service: service,
				repo:    repo,
			}

			upsertedGame, err := gameUsecase.Reveal(c.ID, c.row, c.col)
			expectedError := c.errText
			if expectedError != "" {
				assert.Equal(t, expectedError, err.Error())
			} else {
				assert.Equal(t, 1, upsertedGame.CellsRevealed)
			}
		})
	}
}
