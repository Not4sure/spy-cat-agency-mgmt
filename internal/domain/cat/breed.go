package cat

import "github.com/not4sure/spy-cat-agency-mgmt/internal/common/errors"

var (
	ErrInvalidBreed = errors.NewIncorrectInputError("invalid breed", "Breed name is invalid")
	ErrEmptyBreed   = errors.NewIncorrectInputError("empty breed", "Breed should not be empty")
)

type BreedValidator interface {
	IsValid(breed string) bool
}

type BreedFactory struct {
	vaidator BreedValidator
}

func NewBreedFactory(v BreedValidator) BreedFactory {
	return BreedFactory{
		vaidator: v,
	}
}

func (f BreedFactory) NewBreed(name string) (Breed, error) {
	if name == "" {
		return Breed{}, ErrEmptyBreed
	}
	if !f.vaidator.IsValid(name) {
		return Breed{}, ErrInvalidBreed
	}

	return Breed{name}, nil
}

func MustNewBreed(b string) Breed {
	return Breed{b}
}

type Breed struct {
	b string
}

func (b Breed) IsZero() bool {
	return b == Breed{}
}

func (b Breed) String() string {
	return b.b
}
