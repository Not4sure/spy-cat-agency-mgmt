package http_api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/app/command"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/server/httperr"
)

type createCatParams struct {
	Name              string `json:"name"`
	YearsOfExperience uint   `json:"yearsOfExperience"`
	Breed             string `json:"breed"`
	Salary            uint   `json:"salary"`
}

func (p createCatParams) ToCmd() command.CreateCat {
	return command.CreateCat{
		ID:                uuid.New(),
		Name:              p.Name,
		YearsOfExperience: p.YearsOfExperience,
		Breed:             p.Breed,
		Salary:            p.Salary,
	}
}

type createCatReponse struct {
	ID string `json:"id"`
}

func (s APIServer) CreateCat(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := createCatParams{}
	if err := json.Unmarshal(body, &params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd := params.ToCmd()

	err = s.app.Commands.CreateCat.Handle(r.Context(), cmd)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	rsp := createCatReponse{ID: cmd.ID.String()}
	d, err := json.Marshal(rsp)
	if err != nil {
		http.Error(w, "cannot marshal response", http.StatusInternalServerError)
		return
	}

	w.Write(d)
}
