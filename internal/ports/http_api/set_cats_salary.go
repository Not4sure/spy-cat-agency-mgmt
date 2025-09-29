package http_api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/app/command"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/server/httperr"
)

type setCatsSalaryParams struct {
	Salary uint `json:"salary"`
}

func (p setCatsSalaryParams) ToCmd() command.SetCatsSalary {
	return command.SetCatsSalary{
		Salary: p.Salary,
	}
}

func (s APIServer) SetCatsSalary(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "cannot to parse cat id", http.StatusBadRequest)
		return
	}

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

	cmd := command.SetCatsSalary{
		ID:     id,
		Salary: params.Salary,
	}

	err = s.app.Commands.SetCatsSalary.Handle(r.Context(), cmd)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
}
