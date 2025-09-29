package http_api

import (
	"time"

	"github.com/not4sure/spy-cat-agency-mgmt/internal/app/query"
)

type Cat struct {
	ID                string    `json:"id"`
	CreatedAt         time.Time `json:"createdAt"`
	Name              string    `json:"name"`
	YearsOfExperience uint      `json:"yearsOfExperience"`
	Breed             string    `json:"breed"`
	Salary            uint      `json:"salary"`
}

func appCatToResponse(cat query.Cat) Cat {
	return Cat{
		ID:                cat.ID.String(),
		CreatedAt:         cat.CreatedAt,
		Name:              cat.Name,
		YearsOfExperience: cat.YearsOfExperience,
		Breed:             cat.Breed,
		Salary:            cat.Salary,
	}
}
