package http_api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/app/command"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/server/httperr"
)

func (s APIServer) DeleteCat(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "cannot to parse cat id", http.StatusBadRequest)
		return
	}

	cmd := command.DeleteCat{
		ID: id,
	}

	err = s.app.Commands.DeleteCat.Handle(r.Context(), cmd)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
}
