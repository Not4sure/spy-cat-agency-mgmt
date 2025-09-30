package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/common/decorator"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/domain/cat"
	"github.com/sirupsen/logrus"
)

type SetCatsSalary struct {
	ID     uuid.UUID
	Salary int
}

type SetCatsSalaryHandler decorator.CommandHandler[SetCatsSalary]

type setCatsSalaryHandler struct {
	repo cat.Repository
}

func NewSetCatsSalaryHandler(
	repo cat.Repository,
	logger *logrus.Entry,
) decorator.CommandHandler[SetCatsSalary] {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators(
		setCatsSalaryHandler{repo: repo},
		logger,
	)
}

func (h setCatsSalaryHandler) Handle(ctx context.Context, cmd SetCatsSalary) (err error) {
	return h.repo.UpdateCat(ctx, cmd.ID, func(ctx context.Context, c *cat.Cat) (*cat.Cat, error) {
		err := c.SetSalary(cmd.Salary)
		return c, err
	})
}
