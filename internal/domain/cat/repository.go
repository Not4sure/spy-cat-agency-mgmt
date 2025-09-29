package cat

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	AddCat(ctx context.Context, c *Cat) error

	UpdateCat(
		ctx context.Context,
		catID uuid.UUID,
		updateFn func(ctx context.Context, c *Cat),
	) error

	DeleteCatByID(ctx context.Context, catID uuid.UUID) error

	GetCat(ctx context.Context, catID uuid.UUID) (*Cat, error)

	ListCats(ctx context.Context) ([]*Cat, error)
}
