package http_api

import (
	"encoding/json"
	"net/http"

	"github.com/not4sure/spy-cat-agency-mgmt/internal/app/query"
)

type listCatsResponse struct {
	Cats []Cat `json:"cats"`
}

func (s APIServer) ListCats(w http.ResponseWriter, r *http.Request) {
	appCats, err := s.app.Queries.ListCats.Handle(r.Context(), query.ListCats{})
	if err != nil {
		// TODO: handle error
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	rsp := listCatsResponse{
		Cats: appCatsToResponse(appCats),
	}
	d, err := json.Marshal(rsp)
	if err != nil {
		http.Error(w, "cannot marshal response", http.StatusInternalServerError)
		return
	}

	w.Write(d)
}
