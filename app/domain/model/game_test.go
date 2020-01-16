package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame_Validate(t *testing.T) {
	cases := []struct {
		name    string
		game    Game
		errText string
	}{
		{
			name: "OK",
			game: Game{
				Rows:  10,
				Cols:  10,
				Mines: 1,
			},
			errText: "",
		},
		{
			name: "FAIL/ROWS_MISSING",
			game: Game{
				Cols:  10,
				Mines: 1,
			},
			errText: "mines: must be no greater than -1; rows: cannot be blank.",
		},
		{
			name: "FAIL/ROWS_GREATER_LIMIT",
			game: Game{
				Rows:  100,
				Cols:  10,
				Mines: 1,
			},
			errText: "rows: must be no greater than 50.",
		},
		{
			name: "FAIL/COLS_MISSING",
			game: Game{
				Rows:  10,
				Mines: 1,
			},
			errText: "cols: cannot be blank; mines: must be no greater than -1.",
		},
		{
			name: "FAIL/COLS_GREATER_LIMIT",
			game: Game{
				Rows:  10,
				Cols:  100,
				Mines: 1,
			},
			errText: "cols: must be no greater than 50.",
		},
		{
			name: "FAIL/MINES_MISSING",
			game: Game{
				Cols: 10,
				Rows: 1,
			},
			errText: "mines: cannot be blank.",
		},
		{
			name: "FAIL/MINES_GREATER_LIMIT",
			game: Game{
				Rows:  10,
				Cols:  10,
				Mines: 100,
			},
			errText: "mines: must be no greater than 99.",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expectedError := c.errText
			err := c.game.Validate()
			if err != nil {
				assert.Equal(t, expectedError, err.Error())
			}
		})
	}
}
