package thecatapi

import (
	"context"
	"fmt"
	"slices"
	"sync"
	"time"
)

// CatAPIBreedValidator is BreedValidator implementation
// which uses thecatapi.com for breed validation.
type CatAPIBreedValidator struct {
	sync.Mutex
	origin          string
	cachingInterval time.Duration

	breeds      []string
	lastFetched time.Time
}

// NewCatAPIBreedValidator creates CatAPIBreedValidator.
func NewCatAPIBreedValidator(opts ...CatAPIOption) *CatAPIBreedValidator {
	cav := &CatAPIBreedValidator{
		origin:          "api.thecatapi.com",
		cachingInterval: 30 * time.Second,
	}

	for _, opt := range opts {
		opt(cav)
	}

	return cav
}

// Breeds returns slice of valid breed names.
func (bv *CatAPIBreedValidator) Breeds(ctx context.Context) []string {
	bv.Lock()
	defer bv.Unlock()

	if bv.needToFetch() {
		breeds, err := bv.FetchBreedNames(ctx)
		if err != nil {
			return bv.breeds
		}

		bv.breeds = breeds
	}

	return bv.breeds
}

// IsValid implements cat.BreedValidator
func (bv *CatAPIBreedValidator) IsValid(breed string) bool {
	ctx := context.TODO()

	return slices.Contains(bv.Breeds(ctx), breed)
}

func (bv *CatAPIBreedValidator) makeURL(path string) string {
	return fmt.Sprintf("https://%s%s", bv.origin, path)
}

// needToFetch is true if there's no cached breeds or
// cachingInterval has pased.
func (bv *CatAPIBreedValidator) needToFetch() bool {
	if len(bv.breeds) == 0 {
		return true
	}

	return time.Now().After(bv.lastFetched.Add(bv.cachingInterval))
}
