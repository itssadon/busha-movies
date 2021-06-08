package services

import (
	"log"

	"github.com/itssadon/busha-movies/collections"
	"github.com/itssadon/busha-movies/database"
)

func AddMovieComment(comment collections.MovieComment) (bool, error) {
	newComment := &collections.MovieComment{
		MovieId: comment.MovieId,
		Email:   comment.Email,
		Comment: comment.Comment,
	}

	db := database.NewDBConn()
	_, err := db.Model(newComment).Insert()
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return true, nil
}
