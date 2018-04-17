package handler

import (
	"log"
	"net/http"
)

func serverError(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	log.Println(err)
}
