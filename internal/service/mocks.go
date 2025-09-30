package service

type BreedValidator struct {
	Result bool
}

func (bv BreedValidator) IsValid(_ string) bool {
	return bv.Result
}
