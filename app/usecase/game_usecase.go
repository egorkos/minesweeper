package usecase

import (
	"net/http"
	"time"

	"github.com/egorkos/minesweeper/app/domain/model"
	"github.com/egorkos/minesweeper/app/domain/repository"
	"github.com/egorkos/minesweeper/app/domain/service"
	"github.com/egorkos/minesweeper/app/interface/apierr"
)

type GameUsecase interface {
	StartGame(game model.Game) (model.Game, *apierr.ApiError)
	FindAll() ([]*model.Game, *apierr.ApiError)
	FindByID(id int) (*model.Game, *apierr.ApiError)
	Reveal(ID, row, col int) (*model.Game, *apierr.ApiError)
	Flag(ID, row, col int) (*model.Game, *apierr.ApiError)
}

type gameUsecase struct {
	repo    repository.GameRepository
	service *service.GameService
}

func NewGameUsecase(repo repository.GameRepository, service *service.GameService) *gameUsecase {
	return &gameUsecase{
		repo:    repo,
		service: service,
	}
}

func (g *gameUsecase) StartGame(game model.Game) (model.Game, *apierr.ApiError) {
	newGame := g.service.StartGame(game)
	g.repo.Upsert(&newGame)
	return newGame, nil
}

func (g *gameUsecase) FindAll() ([]*model.Game, *apierr.ApiError) {
	return g.repo.FindAll()
}

func (g *gameUsecase) FindByID(id int) (*model.Game, *apierr.ApiError) {
	return g.repo.FindByID(id)
}

func (g *gameUsecase) Reveal(ID, row, col int) (*model.Game, *apierr.ApiError) {
	game, err := g.FindByID(ID)
	if err != nil {
		return nil, err
	}

	apiError := validateCellUpdate(game, row, col)
	if apiError != nil {
		return nil, apiError
	}

	if game.Grid[row][col].Flagged {
		return nil, apierr.NewAPIError("Can't reveal a flagged cell", http.StatusBadRequest)
	}

	game.Grid[row][col].Revealed = true
	game.CellsRevealed++

	if loose(game, row, col) {
		game.Status = model.Loose
		game.FinishTime = time.Now()
		return game, nil
	}

	if game.Grid[row][col].MinesAround == 0 {
		//TODO:
		//revealAdjacentSquares(game, row, col)
	}

	if win(game) {
		game.Status = model.Win
		game.FinishTime = time.Now()
	}

	err = g.repo.Upsert(game)

	if err != nil {
		return nil, err
	}

	return game, nil
}

func (g *gameUsecase) Flag(ID, row, col int) (*model.Game, *apierr.ApiError) {
	game, err := g.FindByID(ID)
	if err != nil {
		return nil, err
	}

	apiError := validateCellUpdate(game, row, col)
	if apiError != nil {
		return nil, apiError
	}

	if game.Grid[row][col].Flagged {
		game.Grid[row][col].Flagged = false
	} else {
		game.Grid[row][col].Flagged = true
	}

	g.repo.Upsert(game)

	return game, nil
}

func win(game *model.Game) bool {
	if game.CellsRevealed == game.Rows*game.Cols-game.Mines {
		return true
	}
	return false
}

func loose(game *model.Game, row, col int) bool {
	if game.Grid[row][col].Mine {
		return true
	}
	return false
}

func validateCellUpdate(game *model.Game, row, col int) *apierr.ApiError {
	if game.Status != model.Undefined {
		return apierr.NewAPIError("Can't update cells on a finished game", http.StatusBadRequest)
	}

	if row < 0 || row >= game.Rows {
		return apierr.NewAPIError("Row value exceeded grid limits", http.StatusBadRequest)
	}

	if col < 0 || col >= game.Cols {
		return apierr.NewAPIError("Col value exceeded grid limits", http.StatusBadRequest)
	}

	if game.Grid[row][col].Revealed {
		return apierr.NewAPIError("Can't update an already revealed cell", http.StatusBadRequest)
	}

	return nil
}
