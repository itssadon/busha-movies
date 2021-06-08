package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/itssadon/busha-movies/routes"
)

/*
Main function
*/
func main() {
	// Create router with mux
	router := mux.NewRouter()

	// Base route
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-type", "application/json")
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte(`{"status": true, "message": "Welcome to Busha MLAPI (Movie Listing API"}`))
	})

	// Movies listing endpoint
	router.HandleFunc("/movies", routes.GetMovies).Methods("GET")

	// Endpoint to add comment to a movie
	router.HandleFunc("/movies/comments", routes.AddComment).Methods("POST")

	// Port definition
	port := os.Getenv("PORT")
	if port == "" {
		port = "8800"
	}

	// Listen on defined port and serve
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(":"+port, router))
}
