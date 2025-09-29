package http_api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/app/query"
)

func (s APIServer) GetCatByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "cannot to parse cat id", http.StatusBadRequest)
		return
	}

	q := query.CatByID{
		ID: id,
	}

	appCat, err := s.app.Queries.CatByID.Handle(r.Context(), q)
	if err != nil {
		// TODO: handle error
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	rsp := appCatToResponse(appCat)
	d, err := json.Marshal(rsp)
	if err != nil {
		http.Error(w, "cannot marshal response", http.StatusInternalServerError)
		return
	}

	w.Write(d)
}
