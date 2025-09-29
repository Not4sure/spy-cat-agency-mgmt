package http_api

import (
	"net/http"

	"github.com/not4sure/spy-cat-agency-mgmt/internal/app"
)

type APIServer struct{ app app.Application }

func NewAPIServer(app app.Application) APIServer {
	return APIServer{
		app: app,
	}
}

func (s APIServer) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /cat", s.ListCats)
	r.HandleFunc("GET /cat/{id}", s.GetCatByID)
	r.HandleFunc("POST /cat", s.CreateCat)
	r.HandleFunc("PUT /cat/{id}/salary", s.SetCatsSalary)

	// r.HandleFunc("DELETE /cat/{id}", s.DeleteCat)
	//

}
