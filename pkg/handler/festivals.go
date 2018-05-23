package handler

import (
	"net/http"
)

func (h *handler) Festivals(w http.ResponseWriter, r *http.Request) {
	fests, err := h.fpStore.LoadAll()
	if err != nil {
		serverError(w, err)
		return
	}

	h.render.JSON(w, 200, fests)
}
