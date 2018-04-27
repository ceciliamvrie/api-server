package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *handler) Artist(w http.ResponseWriter, r *http.Request) {
	aName := mux.Vars(r)["artistName"]

	a, err := h.aStore.Load(aName)
	if err != nil {
		serverError(w, err)
		return
	}

	h.render.JSON(w, 200, a)
}
