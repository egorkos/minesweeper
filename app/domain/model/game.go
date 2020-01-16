package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Game struct {
	ID            int        `json:"id"`
	StartTime     time.Time  `json:"start_time"`
	FinishTime    time.Time  `json:"finish_time"`
	Rows          int        `json:"rows"`
	Cols          int        `json:"cols"`
	Mines         int        `json:"mines"`
	CellsRevealed int        `json:"cells_revealed"`
	Status        GameStatus `json:"game_status"`
	Grid          [][]Cell   `json:"grid,omitempty"`
}

func (g Game) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.Rows, validation.Required, validation.Min(1)),
		validation.Field(&g.Rows, validation.Required, validation.Max(50)),
		validation.Field(&g.Cols, validation.Required, validation.Min(1)),
		validation.Field(&g.Cols, validation.Required, validation.Max(50)),
		validation.Field(&g.Mines, validation.Required, validation.Min(1)),
		//At least 1 empty cell
		validation.Field(&g.Mines, validation.Required, validation.Max(g.Rows*g.Cols-1)),
	)
}