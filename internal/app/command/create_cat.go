package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/common/decorator"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/domain/cat"
	"github.com/sirupsen/logrus"
)

type CreateCat struct {
	ID                uuid.UUID
	Name              string
	YearsOfExperience uint
	Breed             string
	Salary            uint
}

type CreateCatHandler decorator.CommandHandler[CreateCat]

type createCatHandler struct {
	repo cat.Repository
}

func NewCreateCatHandler(
	repo cat.Repository,
	logger *logrus.Entry,
) decorator.CommandHandler[CreateCat] {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators(
		createCatHandler{repo: repo},
		logger,
	)
}

func (h createCatHandler) Handle(ctx context.Context, cmd CreateCat) (err error) {
	cat, err := cat.New(cmd.ID, cmd.Name, cmd.YearsOfExperience, cmd.Breed, cmd.Salary)
	if err != nil {
		return err
	}

	return h.repo.AddCat(ctx, cat)
}
