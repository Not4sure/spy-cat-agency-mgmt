package query

import (
	"time"

	"github.com/google/uuid"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/domain/cat"
)

type Cat struct {
	ID                uuid.UUID
	CreatedAt         time.Time
	Name              string
	YearsOfExperience int
	Breed             string
	Salary            int
}

func domainCatToApplication(c *cat.Cat) Cat {
	return Cat{
		ID:                c.ID(),
		CreatedAt:         c.CreatedAt(),
		Name:              c.Name(),
		YearsOfExperience: c.YearsOfExperience(),
		Breed:             c.Breed(),
		Salary:            c.Salary(),
	}
}

func domainCatsToApplication(cc []*cat.Cat) []Cat {
	cats := []Cat{}
	for _, c := range cc {
		cats = append(cats, domainCatToApplication(c))
	}

	return cats
}
