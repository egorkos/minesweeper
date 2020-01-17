package registry

import (
	"github.com/egorkos/minesweeper/app/domain/service"
	"github.com/egorkos/minesweeper/app/interface/persistence/memory"
	"github.com/egorkos/minesweeper/app/usecase"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}
	if err := builder.Add([]di.Def{
		{
			Name:  "game-usecase",
			Build: buildGameUsecase,
		},
	}...); err != nil {
		return nil, err
	}
	return &Container{
		ctn: builder.Build(),
	}, nil
}
func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}
func (c *Container) Clean() error {
	return c.ctn.Clean()
}
func buildGameUsecase(ctn di.Container) (interface{}, error) {
	repo := memory.NewGameRepository()
	service := service.NewGameService(repo)
	return usecase.NewGameUsecase(repo, service), nil
}
