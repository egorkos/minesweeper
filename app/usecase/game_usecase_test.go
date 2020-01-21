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
		expStatus  model.GameStatus
		expGame    model.Game
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
			errText: CantUpdateCellsOnAFinishedGame,
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
						Status: model.Running,
					}
					return &game, nil
				},
			},
			row:     10,
			col:     1,
			errText: RowValueExceededGridLimits,
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
						Status: model.Running,
					}
					return &game, nil
				},
			},
			row:     0,
			col:     10,
			errText: ColValueExceededGridLimits,
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
						Status: model.Running,
					}
					return &game, nil
				},
			},
			row:     0,
			col:     0,
			errText: CantUpdateAnAlreadyRevealedCell,
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
						Status: model.Running,
					}
					return &game, nil
				},
			},
			row:     0,
			col:     0,
			errText: CantRevealAFlaggedCell,
		},
		{
			name: "OK/REVEAL/MINED_CELL",
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
						Status: model.Running,
					}
					return &game, nil
				},
			},
			row:       0,
			col:       2,
			errText:   "",
			expStatus: model.Loose,
		},
		{
			name: "OK/REVEAL/WIN_GAME",
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
							{model.Cell{
								Mine:        true,
								Revealed:    false,
								Flagged:     false,
								MinesAround: 0,
							},
								model.Cell{
									Mine:        false,
									Revealed:    false,
									Flagged:     false,
									MinesAround: 1,
								},
								model.Cell{
									Mine:        false,
									Revealed:    false,
									Flagged:     false,
									MinesAround: 0,
								}},
						},
						Status: model.Running,
					}
					return &game, nil
				},
			},
			row:       0,
			col:       2,
			errText:   "",
			expStatus: model.Win,
		},
		{
			name: "OK/REVEAL_ADJACENT_SQUARES",
			ID:   1,
			repository: &mockGameRepository{
				mockUpsert: func(game *model.Game) *apierr.ApiError {
					return nil
				},
				mockFindByID: func(ID int) (*model.Game, *apierr.ApiError) {
					game := model.Game{
						Rows:  3,
						Cols:  3,
						Mines: 1,
						Grid: [][]model.Cell{
							//1st row
							{model.Cell{
								Mine:        false,
								Revealed:    false,
								Flagged:     false,
								MinesAround: 0,
							},
								model.Cell{
									Mine:        false,
									Revealed:    false,
									Flagged:     false,
									MinesAround: 1,
								},
								model.Cell{
									Mine:        false,
									Revealed:    false,
									Flagged:     false,
									MinesAround: 1,
								}},
							//2nd row
							{model.Cell{
								Mine:        false,
								Revealed:    false,
								Flagged:     false,
								MinesAround: 0,
							},
								model.Cell{
									Mine:        false,
									Revealed:    false,
									Flagged:     false,
									MinesAround: 1,
								},
								model.Cell{
									Mine:        true,
									Revealed:    false,
									Flagged:     false,
									MinesAround: 0,
								}},
							//3rd row
							{model.Cell{
								Mine:        false,
								Revealed:    false,
								Flagged:     false,
								MinesAround: 0,
							},
								model.Cell{
									Mine:        false,
									Revealed:    false,
									Flagged:     false,
									MinesAround: 1,
								},
								model.Cell{
									Mine:        false,
									Revealed:    false,
									Flagged:     false,
									MinesAround: 1,
								}},
						},
						Status: model.Running,
					}
					return &game, nil
				},
			},
			row:       1,
			col:       0,
			errText:   "",
			expStatus: model.Running,
			expGame: model.Game{
				Rows:  3,
				Cols:  3,
				Mines: 1,
				Grid: [][]model.Cell{
					//1st row
					{model.Cell{
						Mine:        false,
						Revealed:    true,
						Flagged:     false,
						MinesAround: 0,
					},
						model.Cell{
							Mine:        false,
							Revealed:    true,
							Flagged:     false,
							MinesAround: 1,
						},
						model.Cell{
							Mine:        false,
							Revealed:    false,
							Flagged:     false,
							MinesAround: 1,
						}},
					//2nd row
					{model.Cell{
						Mine:        false,
						Revealed:    true,
						Flagged:     false,
						MinesAround: 0,
					},
						model.Cell{
							Mine:        false,
							Revealed:    true,
							Flagged:     false,
							MinesAround: 1,
						},
						model.Cell{
							Mine:        true,
							Revealed:    false,
							Flagged:     false,
							MinesAround: 0,
						}},
					//3rd row
					{model.Cell{
						Mine:        false,
						Revealed:    true,
						Flagged:     false,
						MinesAround: 0,
					},
						model.Cell{
							Mine:        false,
							Revealed:    true,
							Flagged:     false,
							MinesAround: 1,
						},
						model.Cell{
							Mine:        false,
							Revealed:    false,
							Flagged:     false,
							MinesAround: 1,
						}},
				},
				Status: model.Running,
			},
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
				assert.Equal(t, c.expStatus, upsertedGame.Status)
				if c.expStatus == model.Running {
					for x := 0; x < upsertedGame.Rows; x++ {
						for y := 0; y < upsertedGame.Cols; y++ {
							assert.Equal(t, c.expGame.Grid[x][y].Revealed, upsertedGame.Grid[x][y].Revealed)
						}
					}
				}
			}
		})
	}
}
