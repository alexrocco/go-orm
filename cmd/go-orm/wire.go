// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	app "github.com/alexrocco/go-orm/internal/app/go-orm"
	"github.com/alexrocco/go-orm/internal/app/go-orm/handlers"
	"github.com/alexrocco/go-orm/internal/app/go-orm/repositories"
	"github.com/alexrocco/go-orm/internal/app/go-orm/services"
	"github.com/google/wire"
)

func InitializeApp() app.App {
	wire.Build(
		repositories.NewDB,
		repositories.NewPeopleRepo,
		services.NewPeopleService,
		handlers.NewPeopleHandler,
		app.NewApp)
	return app.App{}
}