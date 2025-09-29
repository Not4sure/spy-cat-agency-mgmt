package query

import (
	"context"

	"github.com/not4sure/spy-cat-agency-mgmt/internal/common/decorator"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/domain/cat"
	"github.com/sirupsen/logrus"
)

type ListCats struct{}

type ListCatsHandler decorator.QueryHandler[ListCats, []Cat]

type listCatsHandler struct {
	repo cat.Repository
}

func NewListCatsHandler(
	repo cat.Repository,
	logger *logrus.Entry,
) ListCatsHandler {
	if repo == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators(
		listCatsHandler{repo: repo},
		logger,
	)
}

func (h listCatsHandler) Handle(ctx context.Context, _ ListCats) ([]Cat, error) {
	cc, err := h.repo.ListCats(ctx)
	if err != nil {
		return []Cat{}, err
	}

	return domainCatsToApplication(cc), nil
}
