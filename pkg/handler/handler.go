package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/techmexdev/lineuplist"
	"github.com/techmexdev/lineuplist/pkg/postgres"
	"github.com/unrolled/render"
)

// Options is configuration for the router
type Options struct {
	Log bool
}

type route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

type handler struct {
	festivalStore lineuplist.FestivalStorage
	render        *render.Render
}

var signature string

// New creates a router with all handlers
func New(dsn string, options Options) *mux.Router {
	router := mux.NewRouter()
	h := handler{
		festivalStore: postgres.NewFestivalStorage(dsn),
		render:        render.New(),
	}

	routes := []route{
		{method: "GET", path: "/festivals", handler: h.Festivals},
	}

	for _, r := range routes {
		router.HandleFunc(r.path, r.handler).Methods(r.method)
	}

	if options.Log {
		router.Use(LogMiddleware)
	}

	return router
}
