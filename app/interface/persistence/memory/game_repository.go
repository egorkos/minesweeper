package memory

import (
	"net/http"
	"sync"

	"github.com/egorkos/minesweeper/app/domain/model"
	"github.com/egorkos/minesweeper/app/interface/apierr"
)

type gameRepository struct {
	mux   *sync.Mutex
	games map[int]*model.Game
}

func NewGameRepository() *gameRepository {
	return &gameRepository{
		mux:   &sync.Mutex{},
		games: map[int]*model.Game{},
	}
}

func (g *gameRepository) FindAll() ([]*model.Game, *apierr.ApiError) {
	g.mux.Lock()
	defer g.mux.Unlock()

	games := make([]*model.Game, len(g.games))
	for i, game := range g.games {
		games[i] = game
	}

	return games, nil
}

func (g *gameRepository) FindByID(id int) (*model.Game, *apierr.ApiError) {
	g.mux.Lock()
	defer g.mux.Unlock()

	game, exists := g.games[id]
	if exists {
		return game, nil
	}

	return nil, apierr.NewAPIError("Game Not Found", http.StatusNotFound)
}

func (g *gameRepository) Upsert(game *model.Game) *apierr.ApiError {
	g.mux.Lock()
	defer g.mux.Unlock()

	if game.ID == 0 {
		game.ID = len(g.games) + 1
	}
	g.games[game.ID] = game

	return nil
}
