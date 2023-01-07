package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Studio struct {
	Studioname   string `json:"studioname"`
	Yearproduced int    `json:"yearproduced"`
}

type Game struct {
	ID       string  `json:"id"`
	Isbn     string  `json:"isbn"`
	Gamename string  `json:"gamename"`
	Studio   *Studio `json:"studio"`
}

// Init games as a slice Game struct
var games []Game

// Get all Games
func getGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func getGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range games {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Game{})
}

func createGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var game Game
	_ = json.NewDecoder(r.Body).Decode(&game)
	game.ID = strconv.Itoa(rand.Intn(1000000))
	games = append(games, game)
	json.NewEncoder(w).Encode(game)
}

func updateGame(w http.ResponseWriter, r *http.Request) {

}

func deleteGame(w http.ResponseWriter, r *http.Request) {

}

func main() {

	//Init Router
	r := mux.NewRouter()

	//Mock Data
	games = append(games, Game{ID: "1", Isbn: "1000", Gamename: "Call of Duty", Studio: &Studio{Studioname: "Infinity Ward", Yearproduced: 2016}})
	games = append(games, Game{ID: "2", Isbn: "1001", Gamename: "Forza Horizon", Studio: &Studio{Studioname: "Playground Games", Yearproduced: 2021}})

	//Route Handlers / Endpoints
	r.HandleFunc("/api/games", getGames).Methods("GET")
	r.HandleFunc("/api/games/{id}", getGame).Methods("GET")
	r.HandleFunc("api/games", createGame).Methods("POST")
	r.HandleFunc("api/games/{id}", updateGame).Methods("PUT")
	r.HandleFunc("api/games/{id}", deleteGame).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
