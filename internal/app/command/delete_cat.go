package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/common/decorator"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/domain/cat"
	"github.com/sirupsen/logrus"
)

type DeleteCat struct {
	ID uuid.UUID
}

type DeleteCatHandler decorator.CommandHandler[DeleteCat]

type deleteCatHandler struct {
	repo cat.Repository
}

func NewDeleteCatHandler(
	repo cat.Repository,
	logger *logrus.Entry,
) decorator.CommandHandler[DeleteCat] {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators(
		deleteCatHandler{repo: repo},
		logger,
	)
}

func (h deleteCatHandler) Handle(ctx context.Context, cmd DeleteCat) (err error) {
	return h.repo.DeleteCatByID(ctx, cmd.ID)
}
