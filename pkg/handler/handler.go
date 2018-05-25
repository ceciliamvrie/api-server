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
	fStore  lineuplist.FestivalStorage
	fpStore lineuplist.FestivalPreviewStorage
	aStore  lineuplist.ArtistStorage
	apStore lineuplist.ArtistPreviewStorage
	render  *render.Render
}

var signature string

// New creates a router with all handlers
func New(dsn string, options Options) *mux.Router {
	router := mux.NewRouter()
	h := handler{
		fStore:  postgres.NewFestivalStorage(dsn),
		fpStore: postgres.NewFestivalPreviewStorage(dsn),
		aStore:  postgres.NewArtistStorage(dsn),
		apStore: postgres.NewArtistPreviewStorage(dsn),
		render:  render.New(),
	}

	routes := []route{
		{method: "GET", path: "/api/festivals", handler: h.Festivals},
		{method: "GET", path: "/api/festivals/{name}", handler: h.Festival},
		{method: "GET", path: "/api/artists", handler: h.Artists},
		{method: "GET", path: "/api/artists/{name}", handler: h.Artist},
	}

	for _, r := range routes {
		router.HandleFunc(r.path, r.handler).Methods(r.method)
	}

	if options.Log {
		router.Use(LogMiddleware)
	}

	return router
}
