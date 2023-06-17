package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (app *Application) routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	mux.Get("/", app.home)

	mux.Get("/query", app.query)

	return mux
}
