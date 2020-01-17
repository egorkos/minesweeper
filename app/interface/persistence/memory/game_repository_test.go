package memory

import (
	"testing"

	"github.com/egorkos/minesweeper/app/domain/model"
)

func TestGameRepositoryUpsert(t *testing.T) {
	cases := []struct {
		name string
		game model.Game
	}{
		{
			name: "OK/SAVE",
			game: model.Game{ID: 0},
		},
		{
			name: "OK/UPDATE",
			game: model.Game{ID: 1, Rows: 10},
		},
	}

	repo := NewGameRepository()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_ = repo.Upsert(&c.game)

		})
	}
}
