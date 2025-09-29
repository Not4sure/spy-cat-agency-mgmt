package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/domain/cat"
)

// CatMemoryRepository is an in-memory implementation
// of a cat.Repository.
type CatMemoryRepository struct {
	sync.Mutex
	cats map[uuid.UUID]cat.Cat
}

func NewCatMemoryRepository() cat.Repository {
	return &CatMemoryRepository{
		cats: map[uuid.UUID]cat.Cat{},
	}
}

// AddCat implements cat.Repository.
func (cr *CatMemoryRepository) AddCat(ctx context.Context, c *cat.Cat) error {
	cr.Lock()
	defer cr.Unlock()

	if _, ok := cr.cats[c.ID()]; ok {
		return errors.New("cat with this id already exists")
	}

	cr.cats[c.ID()] = *c
	return nil
}

// DeleteCatByID implements cat.Repository.
func (cr *CatMemoryRepository) DeleteCatByID(ctx context.Context, catID uuid.UUID) error {
	cr.Lock()
	defer cr.Unlock()

	if _, ok := cr.cats[catID]; !ok {
		return errors.New("cat not found")
	}

	delete(cr.cats, catID)
	return nil
}

// GetCat implements cat.Repository.
func (cr *CatMemoryRepository) GetCat(ctx context.Context, catID uuid.UUID) (*cat.Cat, error) {
	cr.Lock()
	defer cr.Unlock()

	if v, ok := cr.cats[catID]; ok {
		return &v, nil
	}

	return nil, errors.New("cat not found")
}

// ListCats implements cat.Repository.
func (cr *CatMemoryRepository) ListCats(ctx context.Context) ([]*cat.Cat, error) {
	cr.Lock()
	defer cr.Unlock()

	var cc []*cat.Cat
	for _, c := range cr.cats {
		cc = append(cc, &c)
	}

	return cc, nil
}

// UpdateCat implements cat.Repository.
func (cr *CatMemoryRepository) UpdateCat(
	ctx context.Context,
	catID uuid.UUID,
	updateFn func(ctx context.Context, c *cat.Cat) (*cat.Cat, error),
) error {
	cr.Lock()
	defer cr.Unlock()

	cat, ok := cr.cats[catID]
	if !ok {
		return errors.New("cat not found")
	}

	c, err := updateFn(ctx, &cat)
	cr.cats[catID] = *c
	return err
}
