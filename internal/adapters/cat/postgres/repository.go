package postgres

import (
	"context"

	"github.com/emicklei/pgtalk/convert"
	"github.com/google/uuid"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/common/db"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/domain/cat"
)

type PostgresCatRepository struct {
	queries *db.Queries
}

func NewPostgresCatRepository(q *db.Queries) *PostgresCatRepository {
	return &PostgresCatRepository{
		queries: q,
	}
}

// AddCat implements cat.Repository.
func (p *PostgresCatRepository) AddCat(ctx context.Context, c *cat.Cat) error {
	params := db.CreateCatParams{
		ID:                convert.UUID(c.ID()),
		CreatedAt:         convert.TimeToTimestamp(c.CreatedAt()),
		Name:              c.Name(),
		YearsOfExperience: int16(c.YearsOfExperience()),
		Breed:             c.Breed(),
		Salary:            int16(c.Salary()),
	}

	return p.queries.CreateCat(ctx, params)
}

// DeleteCatByID implements cat.Repository.
func (p *PostgresCatRepository) DeleteCatByID(ctx context.Context, catID uuid.UUID) error {
	panic("unimplemented")
}

// GetCat implements cat.Repository.
func (p *PostgresCatRepository) GetCat(ctx context.Context, catID uuid.UUID) (*cat.Cat, error) {
	panic("unimplemented")
}

// ListCats implements cat.Repository.
func (p *PostgresCatRepository) ListCats(ctx context.Context) ([]*cat.Cat, error) {
	rsp, err := p.queries.ListCats(ctx)
	if err != nil {
		return nil, err
	}

	var cats []*cat.Cat
	for _, v := range rsp {
		cats = append(cats, domainCatFromModel(v))
	}
	return cats, nil
}

// UpdateCat implements cat.Repository.
func (p *PostgresCatRepository) UpdateCat(ctx context.Context, catID uuid.UUID, updateFn func(ctx context.Context, c *cat.Cat) (*cat.Cat, error)) error {
	panic("unimplemented")
}

func domainCatFromModel(m db.Cat) *cat.Cat {
	return cat.UnmarshallFromDatabase(
		m.ID.Bytes,
		m.CreatedAt.Time,
		m.Name,
		int(m.YearsOfExperience),
		m.Breed,
		int(m.Salary),
	)
}
