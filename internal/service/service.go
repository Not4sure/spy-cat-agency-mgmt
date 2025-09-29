package service

import (
	"context"

	cat_memory "github.com/not4sure/spy-cat-agency-mgmt/internal/adapters/cat/memory"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/app"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/app/command"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	catRepo := cat_memory.NewCatMemoryRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())

	return app.Application{
		Commands: app.Commands{
			CreateCat: command.NewCreateCatHandler(catRepo, logger),
		},
		Queries: app.Queries{
			CatByID: query.NewCatByIDHandler(catRepo, logger),
		},
	}
}
