package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/k2realty/deals/pkg/config"
	"github.com/k2realty/deals/pkg/handlers"
)

// routes defines all of our application routes, middleware, and fileserver
func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	// define our middleware
	mux.Use(middleware.Recoverer)
	//mux.Use(NoSurf)
	mux.Use(SessionLoad)

	//define our page routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/deals", handlers.Repo.Deals)
	mux.Post("/deals", handlers.Repo.PostDeals)
	mux.Get("/login", handlers.Repo.Login)

	// this code provides our templates with the content inside the static folder.
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
