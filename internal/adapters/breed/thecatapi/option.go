package thecatapi

import "time"

type CatAPIOption func(*CatAPIBreedValidator)

func WithOrigin(o string) CatAPIOption {
	return func(cav *CatAPIBreedValidator) {
		cav.origin = o
	}
}

func WithCachingInterval(d time.Duration) CatAPIOption {
	return func(cav *CatAPIBreedValidator) {
		cav.cachingInterval = d
	}
}
