package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *handler) Festival(w http.ResponseWriter, r *http.Request) {
	fName := mux.Vars(r)["festName"]

	f, err := h.fStore.Load(fName)
	if err != nil {
		serverError(w, err)
		return
	}

	h.render.JSON(w, 200, f)
}
