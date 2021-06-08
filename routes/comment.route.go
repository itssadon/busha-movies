package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/itssadon/busha-movies/collections"
	"github.com/itssadon/busha-movies/services"
)

func AddComment(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	var newComment collections.MovieComment
	err := json.NewDecoder(req.Body).Decode(&newComment)
	if err != nil {
		log.Print(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = services.AddMovieComment(newComment)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusCreated)
}
