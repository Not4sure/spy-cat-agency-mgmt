package app

import (
	"github.com/not4sure/spy-cat-agency-mgmt/internal/app/command"
	"github.com/not4sure/spy-cat-agency-mgmt/internal/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateCat     command.CreateCatHandler
	SetCatsSalary command.SetCatsSalaryHandler
	DeleteCat     command.DeleteCatHandler
}

type Queries struct {
	CatByID  query.CatByIDHandler
	ListCats query.ListCatsHandler
}
