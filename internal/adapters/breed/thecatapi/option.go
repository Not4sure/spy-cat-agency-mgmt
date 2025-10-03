package thecatapi

import "time"

type CatAPIOption func(*CatAPIBreedValidator)

// WithOrigin sets CatAPIBreedValidator origin.
//
//	default is api.thecatapi.com
func WithOrigin(o string) CatAPIOption {
	return func(cav *CatAPIBreedValidator) {
		cav.origin = o
	}
}

// WithCachingInterval sets CatAPIBreedValidator caching interval.
//
// default is 30 seconds.
func WithCachingInterval(d time.Duration) CatAPIOption {
	return func(cav *CatAPIBreedValidator) {
		cav.cachingInterval = d
	}
}
