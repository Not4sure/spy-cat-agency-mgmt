package thecatapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type breedRsp []struct {
	Name string `json:"name"`
}

func (bv *CatAPIBreedValidator) FetchBreedNames(ctx context.Context) ([]string, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, bv.makeURL("/v1/breeds"), nil)
	if err != nil {
		return nil, err
	}

	rsp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("thecatapi returned non-200 status: %v", rsp.StatusCode)
	}

	bb, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	var brsp breedRsp
	err = json.Unmarshal(bb, &brsp)
	if err != nil {
		return nil, err
	}

	var breeds []string
	for _, b := range brsp {
		breeds = append(breeds, b.Name)
	}

	return breeds, nil
}
