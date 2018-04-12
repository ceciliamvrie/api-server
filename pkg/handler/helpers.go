package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	b, err := json.Marshal(data)
	if err != nil {
		serverError(w, err)
		return
	}

	w.WriteHeader(status)
	w.Write(b)
}

func serverError(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	log.Println(err)
}
