package handler

import "net/http"

func (h *handler) Artists(w http.ResponseWriter, r *http.Request) {
	aa, err := h.apStore.LoadAll()
	if err != nil {
		serverError(w, err)
		return
	}

	h.render.JSON(w, 200, aa)
}
