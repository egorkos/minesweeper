package service

import (
	"testing"

	"github.com/egorkos/minesweeper/app/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestGameService(t *testing.T) {
	cases := []struct {
		name string
		game model.Game
	}{
		{
			name: "OK/ONE_MINE",
			game: model.Game{
				Rows:  10,
				Cols:  10,
				Mines: 1,
			},
		},
		{
			name: "OK/FIVE_MINES",
			game: model.Game{
				Rows:  10,
				Cols:  10,
				Mines: 5,
			},
		},
	}

	gameService := &GameService{
		repo: nil,
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mines := 0
			newGame := gameService.StartGame(c.game)
			for x := 0; x < c.game.Rows; x++ {
				for y := 0; y < c.game.Cols; y++ {
					if newGame.Grid[x][y].Mine {
						mines++
					}
				}
			}
			assert.Equal(t, mines, c.game.Mines)
		})
	}
}
