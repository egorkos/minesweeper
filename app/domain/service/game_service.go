package service

import (
	"math/rand"
	"time"

	"github.com/egorkos/minesweeper/app/domain/model"
	"github.com/egorkos/minesweeper/app/domain/repository"
)

type GameService struct {
	repo repository.GameRepository
}

func NewGameService(repo repository.GameRepository) *GameService {
	return &GameService{
		repo: repo,
	}
}

func (g *GameService) StartGame(game model.Game) model.Game {
	game.StartTime = time.Now()
	game.Status = model.Undefined
	createGrid(&game)
	return game
}

func createGrid(game *model.Game) {
	game.Grid = make([][]model.Cell, game.Rows)

	for i := range game.Grid {
		game.Grid[i] = make([]model.Cell, game.Cols)
	}

	setMines(game)
	setMineIndicatorsAroundCell(game)
}

func setMines(game *model.Game) {
	rand.Seed(time.Now().UnixNano())
	i := 0
	for i < game.Mines {
		x := rand.Intn(game.Rows)
		y := rand.Intn(game.Cols)
		if !game.Grid[x][y].Mine {
			game.Grid[x][y].Mine = true
			i++
		}
	}
}

func setMineIndicatorsAroundCell(game *model.Game) {
	for x := 0; x < game.Rows; x++ {
		for y := 0; y < game.Cols; y++ {
			if game.Grid[x][y].Mine {
				for z := x - 1; z < x+2; z++ {
					if z < 0 || z > game.Rows-1 {
						continue
					}
					for w := y - 1; w < y+2; w++ {
						if w < 0 || w > game.Cols-1 {
							continue
						}
						if z == x && w == y {
							continue
						}
						game.Grid[z][w].MinesAround++
					}
				}
			}
		}
	}
}
