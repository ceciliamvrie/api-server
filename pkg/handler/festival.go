package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func (h *handler) Festival(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	name = strings.Replace(name, "-", " ", -1)

	log.Println("handler festival - name: ", name)
	f, err := h.fStore.Load(name)
	if err != nil {
		serverError(w, err)
		return
	}

	h.render.JSON(w, 200, f)
}
