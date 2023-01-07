package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type

func main() {

	//Init Router
	r := mux.NewRouter()

	//Route Handlers / Endpoints
	r.HandleFunc("/api/games", getGames).Methods("GET")
	r.HandleFunc("/api/games/{id}", getGame).Methods("GET")
	r.HandleFunc("api/games", createGame).Methods("POST")
	r.HandleFunc("api/games/{id}", updateGame).Methods("PUT")
	r.HandleFunc("api/games/{id}", deleteGame).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
