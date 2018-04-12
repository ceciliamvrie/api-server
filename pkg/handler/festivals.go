package handler

import "net/http"

func (h *handler) Festivals(w http.ResponseWriter, r *http.Request) {
	fests, err := h.store.GetFests()
	if err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, fests, 200)
}
