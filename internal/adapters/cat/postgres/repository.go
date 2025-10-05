package postgres

import (
	"context"
	"fmt"

	"github.com/emicklei/pgtalk/convert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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
	return p.queries.DeleteCat(ctx, convert.UUID(catID))
}

// GetCat implements cat.Repository.
func (p *PostgresCatRepository) GetCat(ctx context.Context, catID uuid.UUID) (*cat.Cat, error) {
	rsp, err := p.queries.GetCat(ctx, convert.UUID(catID))
	if err != nil {
		fmt.Println(err)
		if err == pgx.ErrNoRows {
			return nil, cat.ErrNotFound
		}

		return nil, err
	}

	return domainCatFromModel(rsp), nil
}

// ListCats implements cat.Repository.
func (p *PostgresCatRepository) ListCats(ctx context.Context) ([]*cat.Cat, error) {
	rsp, err := p.queries.ListCats(ctx)
	if err != nil {
		return nil, err
	}

	cats := []*cat.Cat{}
	for _, v := range rsp {
		cats = append(cats, domainCatFromModel(v))
	}
	return cats, nil
}

// UpdateCat implements cat.Repository.
func (p *PostgresCatRepository) UpdateCat(ctx context.Context, catID uuid.UUID, updateFn func(ctx context.Context, c *cat.Cat) (*cat.Cat, error)) error {
	cat, err := p.GetCat(ctx, catID)
	if err != nil {
		return err
	}

	updated, err := updateFn(ctx, cat)
	if err != nil {
		return err
	}

	arg := db.UpdateCatParams{
		ID:                convert.UUID(updated.ID()),
		Name:              updated.Name(),
		YearsOfExperience: int16(updated.YearsOfExperience()),
		Breed:             updated.Breed(),
		Salary:            int16(updated.Salary()),
	}
	return p.queries.UpdateCat(ctx, arg)
}

// domainCatFromModel unmarshalls doman Cat from database model.
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
