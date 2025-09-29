package main

import (
	"context"
	"net/http"

	"github.com/not4sure/spy-cat-agency-mgmt/internal/ports/http_api"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/server"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/service"
)

func main() {
	ctx := context.Background()

	app := service.NewApplication(ctx)

	server.RunServer(func(router *http.ServeMux) {
		apiServer := http_api.NewAPIServer(app)
		apiServer.RegisterRoutes(router)
	})
}
