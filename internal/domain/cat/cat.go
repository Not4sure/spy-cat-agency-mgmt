package cat

import (
	"time"

	"github.com/google/uuid"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/common/errors"
)

var (
	ErrEmptyName          = errors.NewIncorrectInputError("empty name", "Name should not be empty")
	ErrInvalidSalary      = errors.NewIncorrectInputError("invalid salary", "Salary should be greater than zero")
	ErrNegativeExperience = errors.NewIncorrectInputError("negative experience", "Experience should not be negative")

	ErrNotFound = errors.NewNotFoundError("not found", "Cat not found")
)

type Cat struct {
	id        uuid.UUID
	createdAt time.Time

	name              string
	yearsOfExperience int
	breed             Breed
	salary            int
}

func New(
	id uuid.UUID,
	name string,
	yearsOfExperience int,
	breed Breed,
	salary int,
) (*Cat, error) {
	if name == "" {
		return nil, ErrEmptyName
	}
	if yearsOfExperience < 0 {
		return nil, ErrNegativeExperience
	}
	if salary <= 0 {
		return nil, ErrInvalidSalary
	}

	return &Cat{
		id:        id,
		createdAt: time.Now(),

		name:              name,
		yearsOfExperience: yearsOfExperience,
		breed:             breed,
		salary:            salary,
	}, nil
}

// Use this function to unmarshall database object into domain cat.
// Do not use for cat creation as it can put domain into invalid state.
func UnmarshallFromDatabase(
	id uuid.UUID,
	createdAt time.Time,
	name string,
	yearsOfExperience int,
	breed string,
	salary int,
) *Cat {
	return &Cat{
		id:                id,
		createdAt:         createdAt,
		name:              name,
		yearsOfExperience: yearsOfExperience,
		breed:             MustNewBreed(breed),
		salary:            salary,
	}
}

func (c *Cat) SetSalary(s int) error {
	if s <= 0 {
		return ErrInvalidSalary
	}

	c.salary = s
	return nil
}

func (c Cat) ID() uuid.UUID {
	return c.id
}

func (c Cat) CreatedAt() time.Time {
	return c.createdAt
}

func (c Cat) Name() string {
	return c.name
}

func (c Cat) YearsOfExperience() int {
	return c.yearsOfExperience
}

func (c Cat) Breed() string {
	return c.breed.String()
}

func (c Cat) Salary() int {
	return c.salary
}
