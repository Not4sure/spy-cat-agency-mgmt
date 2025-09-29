package cat

import (
	"time"

	"github.com/google/uuid"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/common/errors"
)

var (
	ErrEmptyName  = errors.NewIncorrectInputError("empty name", "Name should not be empty")
	ErrEmptyBreed = errors.NewIncorrectInputError("empty breed", "Breed should not be empty")
	ErrZeroSalary = errors.NewIncorrectInputError("zero salary", "Salary should be greater than zero")

	ErrNotFound = errors.NewNotFoundError("not found", "Cat not found")
)

type Cat struct {
	id        uuid.UUID
	createdAt time.Time

	name              string
	yearsOfExperience uint
	breed             string
	salary            uint
}

func New(
	id uuid.UUID,
	name string,
	yearsOfExperience uint,
	breed string,
	salary uint,
) (*Cat, error) {
	if name == "" {
		return nil, ErrEmptyName
	}
	if breed == "" {
		return nil, ErrEmptyBreed
	}
	if salary == 0 {
		return nil, ErrZeroSalary
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

func (c *Cat) SetSalary(s uint) error {
	if s == 0 {
		return ErrZeroSalary
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

func (c Cat) YearsOfExperience() uint {
	return c.yearsOfExperience
}

func (c Cat) Breed() string {
	return c.breed
}

func (c Cat) Salary() uint {
	return c.salary
}
