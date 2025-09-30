package service

import (
	"context"

	"github.com/not4sure/spy-cat-agency-mgmt/internal/adapters/breed/thecatapi"
	cat_memory "github.com/not4sure/spy-cat-agency-mgmt/internal/adapters/cat/memory"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/app"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/app/command"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/app/query"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/domain/cat"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	catRepo := cat_memory.NewCatMemoryRepository()
	validator := thecatapi.NewCatAPIBreedValidator()
	breedFactory := cat.NewBreedFactory(validator)

	logger := logrus.NewEntry(logrus.StandardLogger())

	return app.Application{
		Commands: app.Commands{
			CreateCat:     command.NewCreateCatHandler(catRepo, breedFactory, logger),
			SetCatsSalary: command.NewSetCatsSalaryHandler(catRepo, logger),
		},
		Queries: app.Queries{
			CatByID:  query.NewCatByIDHandler(catRepo, logger),
			ListCats: query.NewListCatsHandler(catRepo, logger),
		},
	}
}
