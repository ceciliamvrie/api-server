package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/techmexdev/lineuplist/pkg/storage"
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
	store storage.Storage
}

var signature string

// New creates a router with all handlers
func New(store storage.Storage, options Options) *mux.Router {
	router := mux.NewRouter()
	h := handler{store}
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
