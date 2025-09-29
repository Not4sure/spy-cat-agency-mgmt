package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/common/decorator"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/domain/cat"
	"github.com/sirupsen/logrus"
)

type CatByID struct {
	ID uuid.UUID
}

type CatByIDHandler decorator.QueryHandler[CatByID, Cat]

type catByIDHandler struct {
	repo cat.Repository
}

func NewCatByIDHandler(
	repo cat.Repository,
	logger *logrus.Entry,
) CatByIDHandler {
	if repo == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators(
		catByIDHandler{repo: repo},
		logger,
	)
}

type CatByIDReadModel interface {
	GetCatByID(ctx context.Context, id uuid.UUID) (Cat, error)
}

func (h catByIDHandler) Handle(ctx context.Context, q CatByID) (Cat, error) {
	c, err := h.repo.GetCat(ctx, q.ID)
	if err != nil {
		return Cat{}, err
	}

	return domainCatToApplication(c), nil
}
