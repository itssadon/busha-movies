package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/itssadon/busha-movies/routes"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-type", "application/json")
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte(`{"status": true, "message": "Welcome to Busha MLAPI (Movie Listing API"}`))
	})

	router.HandleFunc("/movies", routes.GetMovies).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(":"+port, router))
}
