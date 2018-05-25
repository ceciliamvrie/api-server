package handler

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/techmexdev/lineuplist"
)

func (h *handler) Artist(w http.ResponseWriter, r *http.Request) {
	var a lineuplist.Artist
	var err error

	name := mux.Vars(r)["name"]
	log.Println("name: ", name)

	a, err = h.aStore.Load(name)
	if err != nil {
		a, err = spotifyArtist(name)
		if err != nil {
			serverError(w, err)
			return
		}
	}

	h.render.JSON(w, 200, a)
}

func spotifyArtist(name string) (lineuplist.Artist, error) {
	var a lineuplist.Artist

	res, err := http.Get("http://localhost:8081/artists/" + name)
	if err != nil {
		return lineuplist.Artist{}, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&a)
	if err != nil {
		return a, err
	}

	return a, nil
}
